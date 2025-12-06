import json
import logging
import os
import subprocess
import tempfile
from datetime import datetime
from functools import lru_cache
from pathlib import Path
from typing import Any, Dict, List, Optional

# Set environment variables before importing TTS
os.environ.setdefault("TORCHAUDIO_FORCE_OLD_AUDIO_IO", "1")
os.environ["TTS_HOME"] = "/root/.local/share/tts"
# Auto-accept TTS license agreements
os.environ["TTS_LICENSE_ACCEPTED"] = "1"

# Patch TTS to auto-accept license agreements
def patch_tts_license():
    """Patch TTS ModelManager to auto-accept license agreements"""
    try:
        from TTS.utils.manage import ModelManager
        original_ask_tos = ModelManager.ask_tos
        
        def patched_ask_tos(self, output_path):
            """Auto-accept TOS without prompting"""
            return True
        
        ModelManager.ask_tos = patched_ask_tos
    except Exception as e:
        print(f"Warning: Could not patch TTS license: {e}")

patch_tts_license()

# Patch torchaudio.load to use soundfile before importing TTS
from app.audio_patch import patch_torchaudio, patch_phoneme_dataset
patch_torchaudio()
patch_phoneme_dataset()

# Patch torch.load to handle weights_only parameter for PyTorch 2.6+
def patch_torch_load():
    """Patch torch.load to disable weights_only for TTS model loading"""
    import torch
    original_load = torch.load
    
    def patched_load(f, map_location=None, pickle_module=None, **kwargs):
        # If weights_only is not specified, set it to False for compatibility
        if 'weights_only' not in kwargs:
            kwargs['weights_only'] = False
        return original_load(f, map_location=map_location, pickle_module=pickle_module, **kwargs)
    
    torch.load = patched_load

patch_torch_load()

from fastapi import BackgroundTasks, FastAPI, HTTPException, UploadFile, File, Form
from fastapi.responses import FileResponse
from fastapi.concurrency import run_in_threadpool
from pydantic import BaseModel, Field
from TTS.api import TTS
from TTS.utils.manage import ModelManager
from TTS.tts import datasets as tts_datasets

from app.formatter import tts_service_two_column
from app.train_job import prepare_training_config, run_training_job as launch_training_job

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger("tts_service")

def split_text_into_chunks(text: str, max_chars: int = 240) -> List[str]:
    """Split long text into chunks that respect sentence boundaries"""
    sentences = text.replace('\r\n', '\n').split('.')
    chunks = []
    current_chunk = ""
    
    for sentence in sentences:
        sentence = sentence.strip()
        if not sentence:
            continue
        
        sentence_with_period = sentence + "."
        
        # If single sentence exceeds limit, split by comma
        if len(sentence_with_period) > max_chars:
            parts = sentence.split(',')
            for part in parts:
                part = part.strip()
                if not part:
                    continue
                part_with_comma = part + ","
                if len(current_chunk) + len(part_with_comma) > max_chars:
                    if current_chunk:
                        chunks.append(current_chunk.rstrip(',').strip() + ".")
                    current_chunk = part
                else:
                    current_chunk += (" " if current_chunk else "") + part
        else:
            if len(current_chunk) + len(sentence_with_period) > max_chars:
                if current_chunk:
                    chunks.append(current_chunk.rstrip('.').strip() + ".")
                current_chunk = sentence
            else:
                current_chunk += (" " if current_chunk else "") + sentence
    
    if current_chunk:
        chunks.append(current_chunk.rstrip('.').strip() + ".")
    
    return [chunk.strip() for chunk in chunks if chunk.strip()]

def merge_audio_files(audio_paths: List[str], output_path: str, crossfade_duration: float = 0.2):
    """Merge multiple audio files with smooth crossfade transitions"""
    import soundfile as sf
    import numpy as np
    
    if not audio_paths:
        raise ValueError("No audio files to merge")
    
    if len(audio_paths) == 1:
        # If only one file, just copy it
        import shutil
        shutil.copy(audio_paths[0], output_path)
        return
    
    # Load all audio files
    audio_data_list = []
    sample_rate = None
    
    for path in audio_paths:
        data, sr = sf.read(path)
        if sample_rate is None:
            sample_rate = sr
        elif sr != sample_rate:
            # Resample if needed
            import librosa
            data = librosa.resample(data, orig_sr=sr, target_sr=sample_rate)
        audio_data_list.append(data)
    
    # Calculate crossfade samples with minimum of 0.2 seconds
    crossfade_samples = max(int(crossfade_duration * sample_rate), int(0.2 * sample_rate))
    
    # Merge with smooth crossfade
    merged = audio_data_list[0].copy()
    
    for i in range(1, len(audio_data_list)):
        next_audio = audio_data_list[i]
        
        if len(merged) > crossfade_samples and len(next_audio) > crossfade_samples:
            # Create smooth crossfade using cosine curve (sounds more natural)
            fade_out = np.cos(np.linspace(0, np.pi / 2, crossfade_samples)) ** 2
            fade_in = np.sin(np.linspace(0, np.pi / 2, crossfade_samples)) ** 2
            
            # Apply crossfade with overlap
            merged[-crossfade_samples:] *= fade_out
            next_audio[:crossfade_samples] *= fade_in
            
            # Combine overlapping sections
            merged[-crossfade_samples:] += next_audio[:crossfade_samples]
            merged = np.concatenate([merged, next_audio[crossfade_samples:]])
        else:
            # If files too short for crossfade, use simple concatenation with small fade
            if len(merged) > 100 and len(next_audio) > 100:
                # Apply small fade at boundaries
                fade_samples = min(100, len(merged) // 4, len(next_audio) // 4)
                fade_out = np.linspace(1, 0.9, fade_samples)
                fade_in = np.linspace(0.9, 1, fade_samples)
                merged[-fade_samples:] *= fade_out
                next_audio[:fade_samples] *= fade_in
            merged = np.concatenate([merged, next_audio])
    
    # Normalize final audio with gentle compression to prevent clipping
    max_val = np.max(np.abs(merged))
    if max_val > 0:
        # Use 0.9 instead of 0.95 to leave more headroom
        merged = merged / max_val * 0.9
    
    # Apply gentle high-pass filter to reduce clicks (optional)
    # This helps smooth out any remaining discontinuities
    try:
        from scipy import signal
        # Design a high-pass filter to remove very low frequencies that can cause clicks
        sos = signal.butter(4, 50, 'hp', fs=sample_rate, output='sos')
        merged = signal.sosfilt(sos, merged)
    except ImportError:
        # If scipy not available, skip filtering
        pass
    
    # Save merged audio
    sf.write(output_path, merged, sample_rate)
    logger.info("Audio files merged successfully: %d files merged with smooth crossfade (%.2f seconds)", 
                len(audio_paths), crossfade_duration)

DATA_ROOT = Path(os.getenv("DATA_ROOT", "/app/data"))
TRAIN_OUTPUT_ROOT = Path(os.getenv("TRAIN_OUTPUT_ROOT", "/app/training_runs"))
BASE_TRAIN_MODEL = os.getenv("BASE_TRAIN_MODEL", "tts_models/multilingual/multi-dataset/xtts_v2")

for path in (DATA_ROOT, TRAIN_OUTPUT_ROOT):
    path.mkdir(parents=True, exist_ok=True)

setattr(tts_datasets, "tts_service_two_column", tts_service_two_column)

VOICE_CONFIGS: Dict[str, Dict[str, Optional[str]]] = {
    "fa_standard": {
        "model_name": os.getenv("TTS_MODEL_NAME_FA", "tts_models/multilingual/multi-dataset/xtts_v2"),
        "speaker": os.getenv("TTS_SPEAKER_ID_FA") or None,
        "language": "fa",
    },
    "de_standard": {
        "model_name": os.getenv("TTS_MODEL_NAME_DE", "tts_models/multilingual/multi-dataset/xtts_v2"),
        "speaker": os.getenv("TTS_SPEAKER_ID_DE") or None,
        "language": "de",
    },
    "en_standard": {
        "model_name": os.getenv("TTS_MODEL_NAME_EN", "tts_models/multilingual/multi-dataset/xtts_v2"),
        "speaker": os.getenv("TTS_SPEAKER_ID_EN") or None,
        "language": "en",
    },
}


@lru_cache(maxsize=len(VOICE_CONFIGS))
def get_tts_engine(model_name: str) -> TTS:
    logger.info("Loading TTS model: %s", model_name)
    engine = TTS(model_name=model_name)
    speakers = getattr(engine, 'speakers', None)
    logger.info("Loaded model '%s' with speakers: %s", model_name, speakers)
    return engine

app = FastAPI(title="Python TTS Service", version="0.1.0")

# Track service readiness
_service_ready = False
_tts_engines_loaded = {}


@app.on_event("startup")
async def startup_event():
    """Initialize TTS engines on startup"""
    global _service_ready, _tts_engines_loaded
    logger.info("Service starting up, pre-loading TTS engines...")
    try:
        for speaker_id, config in VOICE_CONFIGS.items():
            logger.info("Pre-loading engine for %s with model %s", speaker_id, config["model_name"])
            engine = get_tts_engine(config["model_name"])
            speakers = getattr(engine, 'speakers', None)
            logger.info("Engine %s loaded with speakers: %s", speaker_id, speakers)
            _tts_engines_loaded[speaker_id] = engine
        _service_ready = True
        logger.info("Service startup complete - all engines ready")
    except Exception as e:
        logger.error("Failed to pre-load engines: %s", e)
        _service_ready = False


@app.get("/health")
async def health_check():
    """Health check endpoint for Docker healthcheck and service readiness"""
    if not _service_ready:
        return {"status": "starting", "ready": False}
    return {"status": "ready", "ready": True}


class GenerateRequest(BaseModel):
    text: str = Field(..., min_length=1, max_length=5000)
    speaker_id: str = Field(..., min_length=1, description="Allowed values: en_xtts, de_xtts, fa_xtts")


class TrainRequest(BaseModel):
    dataset_name: str = Field(..., min_length=1)
    epochs: int = Field(10, ge=1, le=1000)
    batch_size: int = Field(4, ge=1, le=64)
    base_model: Optional[str] = Field(
        default=None, description="Optional override for pretrained model name registered in Coqui"
    )


async def _synthesize_to_path(engine: TTS, text: str, speaker_id: Optional[str], file_path: str, language: Optional[str] = None) -> None:
    def _run() -> None:
        try:
            # XTTS v2 uses language parameter and requires a speaker for multi-speaker models
            if language:
                # Get available speakers from the engine
                speakers = getattr(engine, 'speakers', None)
                logger.info("Engine speakers attribute: %s (type: %s)", speakers, type(speakers).__name__)
                
                if speakers:
                    # Use first available speaker as default
                    if isinstance(speakers, list):
                        default_speaker = speakers[0]
                    elif isinstance(speakers, dict):
                        default_speaker = next(iter(speakers))
                    else:
                        default_speaker = None
                    
                    logger.info("Using default speaker '%s' for language '%s'", default_speaker, language)
                    engine.tts_to_file(text=text, speaker=default_speaker, language=language, file_path=file_path)
                else:
                    # Fallback if no speakers available - try without speaker
                    logger.warning("No speakers found, attempting synthesis without speaker parameter")
                    engine.tts_to_file(text=text, language=language, file_path=file_path)
            else:
                engine.tts_to_file(text=text, speaker=speaker_id, file_path=file_path)
        except ValueError as err:
            logger.error("ValueError during synthesis: %s", str(err))
            raise HTTPException(status_code=400, detail=str(err)) from err
        except Exception as err:  # pragma: no cover - unexpected failure
            logger.exception("Unexpected error during synthesis")
            raise HTTPException(status_code=500, detail="Synthesis failed") from err

    await run_in_threadpool(_run)


def _schedule_cleanup(background_tasks: BackgroundTasks, path: str) -> None:
    def _remove_file() -> None:
        try:
            os.remove(path)
        except FileNotFoundError:
            logger.warning("Temporary file already removed: %s", path)
        except Exception:  # pragma: no cover
            logger.exception("Failed to remove temporary file: %s", path)

    background_tasks.add_task(_remove_file)


@app.post("/generate", response_class=FileResponse, responses={200: {"content": {"audio/wav": {}}, "description": "Generated speech"}})
async def generate_audio(payload: GenerateRequest, background_tasks: BackgroundTasks):
    voice_config = VOICE_CONFIGS.get(payload.speaker_id)
    if not voice_config:
        raise HTTPException(status_code=400, detail=f"speaker_id '{payload.speaker_id}' is not supported. Choose from: {list(VOICE_CONFIGS.keys())}")

    engine = get_tts_engine(voice_config["model_name"])

    with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp_file:
        tmp_path = tmp_file.name

    await _synthesize_to_path(engine, payload.text, voice_config["speaker"], tmp_path, language=voice_config["language"])
    _schedule_cleanup(background_tasks, tmp_path)

    filename = f"tts_{payload.speaker_id}.wav"
    return FileResponse(path=tmp_path, media_type="audio/wav", filename=filename)


@app.post("/tts", response_class=FileResponse, responses={200: {"content": {"audio/wav": {}}, "description": "Generated speech"}})
async def tts_endpoint(payload: GenerateRequest, background_tasks: BackgroundTasks):
    """Alias for /generate endpoint to support Go service proxy"""
    return await generate_audio(payload, background_tasks)


class SynthesizeTrainedRequest(BaseModel):
    text: str = Field(..., min_length=1, max_length=5000)
    job_id: str = Field(..., min_length=1, description="Training job ID (e.g., my_dataset_20251204145511)")


class VoiceCloneRequest(BaseModel):
    text: str = Field(..., min_length=1, max_length=5000)
    language: str = Field(..., description="Language code: fa (Farsi), de (German), en (English)")
    speed: float = Field(1.0, ge=0.5, le=2.0, description="Speech speed: 0.5 (slow) to 2.0 (fast)")
    emotion: str = Field("neutral", description="Speech emotion/style: neutral, happy, sad, angry")
    model_type: str = Field("your_tts", description="Model type: 'your_tts' (better voice cloning) or 'glow_tts' (faster)")


@app.post("/tts-trained", response_class=FileResponse, responses={200: {"content": {"audio/wav": {}}, "description": "Generated speech from trained model"}})
async def synthesize_trained(payload: SynthesizeTrainedRequest, background_tasks: BackgroundTasks):
    """Synthesize audio using a trained model"""
    job_dir = TRAIN_OUTPUT_ROOT / payload.job_id
    best_model_path = None
    
    # Find best model checkpoint
    if job_dir.exists():
        for file in job_dir.glob("best_model_*.pth"):
            best_model_path = file
            break
    
    if not best_model_path:
        raise HTTPException(status_code=404, detail=f"No trained model found for job_id '{payload.job_id}'")
    
    logger.info("Loading trained model from: %s", best_model_path)
    engine = TTS(model_path=str(best_model_path), gpu=False)
    
    with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp_file:
        tmp_path = tmp_file.name
    
    try:
        await _synthesize_to_path(engine, payload.text, None, tmp_path)
        _schedule_cleanup(background_tasks, tmp_path)
        
        filename = f"tts_trained_{payload.job_id}.wav"
        return FileResponse(path=tmp_path, media_type="audio/wav", filename=filename)
    except Exception as err:
        logger.exception("Synthesis from trained model failed")
        raise HTTPException(status_code=500, detail="Synthesis failed") from err


@app.post("/voice-clone", response_class=FileResponse, responses={200: {"content": {"audio/wav": {}}, "description": "Generated speech with cloned voice"}})
async def voice_clone(
    audio: UploadFile = File(...),
    text: str = Form(...),
    language: str = Form(...),
    speed: str = Form(default="1.0"),
    emotion: str = Form(default="neutral"),
    model_type: str = Form(default="your_tts"),
    background_tasks: BackgroundTasks = None
):
    """Clone a voice from uploaded audio and synthesize text with optional speed and emotion control"""
    logger.info("Voice clone request received: audio=%s, text_len=%d, language=%s, speed=%s, emotion=%s, model_type=%s", 
                audio.filename if audio else None, len(text), language, speed, emotion, model_type)
    
    if not audio or not text or not language:
        logger.error("Missing required fields: audio=%s, text=%s, language=%s", audio is not None, bool(text), bool(language))
        raise HTTPException(status_code=400, detail="audio, text, and language are required")
    
    # Convert speed to float
    try:
        speed_float = float(speed)
        logger.info("Speed converted to float: %f", speed_float)
    except (ValueError, TypeError) as e:
        logger.error("Speed conversion failed: %s", e)
        raise HTTPException(status_code=400, detail="speed must be a valid number")
    
    # Validate speed parameter
    if speed_float < 0.5 or speed_float > 2.0:
        logger.error("Speed out of range: %f", speed_float)
        raise HTTPException(status_code=400, detail="speed must be between 0.5 and 2.0")
    
    # Validate emotion parameter
    valid_emotions = ["neutral", "happy", "sad", "angry"]
    if emotion not in valid_emotions:
        logger.error("Invalid emotion: %s", emotion)
        raise HTTPException(status_code=400, detail=f"emotion must be one of: {', '.join(valid_emotions)}")
    
    if language not in ["fa", "de", "en"]:
        logger.error("Invalid language: %s", language)
        raise HTTPException(status_code=400, detail="language must be 'fa', 'de', or 'en'")
    
    # Validate model_type parameter
    valid_models = ["your_tts", "glow_tts"]
    if model_type not in valid_models:
        logger.error("Invalid model_type: %s", model_type)
        raise HTTPException(status_code=400, detail=f"model_type must be one of: {', '.join(valid_models)}")
    
    logger.info("Voice cloning request: language=%s, text_length=%d", language, len(text))
    
    # Save uploaded voice sample temporarily
    file_content = await audio.read()
    
    # Determine file extension from original filename
    original_filename = audio.filename or "audio.wav"
    file_ext = Path(original_filename).suffix.lower() or ".wav"
    
    with tempfile.NamedTemporaryFile(delete=False, suffix=file_ext) as tmp_voice:
        tmp_voice.write(file_content)
        voice_sample_path = tmp_voice.name
    
    # Path for converted WAV file
    wav_sample_path = voice_sample_path.replace(file_ext, ".wav")
    
    with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp_output:
        output_path = tmp_output.name
    
    try:
        # Validate and convert audio file to proper format
        def _validate_audio():
            try:
                import soundfile as sf
                import numpy as np
                
                # Try to read the audio file (handles MP3, WAV, etc.)
                logger.info("Validating audio file: %s", voice_sample_path)
                audio_data, sr = sf.read(voice_sample_path)
                
                # Ensure audio is mono or convert to mono
                if len(audio_data.shape) > 1:
                    audio_data = np.mean(audio_data, axis=1)
                
                # Resample to 22050 Hz if needed (standard for TTS)
                if sr != 22050:
                    import librosa
                    audio_data = librosa.resample(audio_data, orig_sr=sr, target_sr=22050)
                
                # Ensure float32 dtype
                audio_data = audio_data.astype(np.float32)
                
                # Save as WAV file
                sf.write(wav_sample_path, audio_data, 22050)
                logger.info("Audio file validated and normalized to: %s", wav_sample_path)
            except Exception as err:
                logger.exception("Audio validation failed: %s", err)
                raise HTTPException(status_code=400, detail=f"Invalid audio file: {str(err)}") from err
        
        await run_in_threadpool(_validate_audio)
        
        # Map language to model that supports voice cloning
        # Use XTTS v2 for best voice cloning quality
        language_config = {
            "fa": {"model": "tts_models/multilingual/multi-dataset/xtts_v2"},
            "de": {"model": "tts_models/multilingual/multi-dataset/xtts_v2"},
            "en": {"model": "tts_models/multilingual/multi-dataset/xtts_v2"},
        }
        
        config = language_config[language]
        model_name = config["model"]
        
        logger.info("Loading model for voice cloning: %s (type: %s)", model_name, model_type)
        engine = get_tts_engine(model_name)
        
        # Synthesize with voice cloning using the uploaded voice sample
        def _clone_voice():
            try:
                import soundfile as sf
                import numpy as np
                
                # Split text into chunks if it's too long
                text_chunks = split_text_into_chunks(text, max_chars=240)
                logger.info("Text split into %d chunks for processing", len(text_chunks))
                
                chunk_audio_paths = []
                
                # Process each chunk
                for idx, chunk in enumerate(text_chunks):
                    chunk_output = output_path.replace(".wav", f"_chunk_{idx}.wav")
                    logger.info("Processing chunk %d/%d: %s", idx + 1, len(text_chunks), chunk[:50])
                    
                    # Synthesize chunk with voice cloning
                    engine.tts_to_file(
                        text=chunk,
                        speaker_wav=wav_sample_path,
                        language=language,
                        file_path=chunk_output
                    )
                    chunk_audio_paths.append(chunk_output)
                
                logger.info("All chunks synthesized, merging audio files")
                
                # Merge all chunks with smooth crossfade (0.2 seconds for natural transitions)
                merge_audio_files(chunk_audio_paths, output_path, crossfade_duration=0.2)
                
                # Clean up chunk files
                for chunk_path in chunk_audio_paths:
                    try:
                        Path(chunk_path).unlink()
                    except:
                        pass
                
                # Apply speed adjustment if needed
                if speed_float != 1.0:
                    logger.info("Applying speed adjustment: %f", speed_float)
                    wav_data, sr = sf.read(output_path)
                    # Adjust speed by resampling
                    adjusted_length = int(len(wav_data) / speed_float)
                    wav_data_adjusted = np.interp(
                        np.linspace(0, len(wav_data) - 1, adjusted_length),
                        np.arange(len(wav_data)),
                        wav_data
                    )
                    sf.write(output_path, wav_data_adjusted, sr)
                    logger.info("Speed adjustment applied")
                
                # Final normalization
                wav_data, sr = sf.read(output_path)
                max_val = np.max(np.abs(wav_data))
                if max_val > 0:
                    wav_data = wav_data / max_val * 0.95
                    sf.write(output_path, wav_data, sr)
                    logger.info("Audio normalized")
                
                # Verify output file exists and has content
                if os.path.exists(output_path):
                    file_size = os.path.getsize(output_path)
                    logger.info("Voice synthesis succeeded - output file created: %s (size: %d bytes)", output_path, file_size)
                else:
                    logger.error("Voice synthesis failed - output file not created at: %s", output_path)
                    raise Exception(f"Output file not created at {output_path}")
            except Exception as err:
                logger.exception("Voice synthesis failed with speaker_wav: %s", err)
                raise HTTPException(status_code=500, detail=f"Voice cloning failed: {str(err)}") from err
        
        await run_in_threadpool(_clone_voice)
        _schedule_cleanup(background_tasks, voice_sample_path)
        _schedule_cleanup(background_tasks, wav_sample_path)
        _schedule_cleanup(background_tasks, output_path)
        
        filename = f"voice_clone_{language}.wav"
        return FileResponse(path=output_path, media_type="audio/wav", filename=filename)
    except Exception as err:
        logger.exception("Voice cloning failed")
        raise HTTPException(status_code=500, detail="Voice cloning failed") from err


@app.get("/trained-models")
async def list_trained_models() -> dict[str, Any]:
    """List all available trained models"""
    trained_models = []
    
    if TRAIN_OUTPUT_ROOT.exists():
        for job_dir in sorted(TRAIN_OUTPUT_ROOT.iterdir(), reverse=True):
            if job_dir.is_dir():
                # Check if best model exists
                best_models = list(job_dir.glob("best_model_*.pth"))
                if best_models:
                    trained_models.append({
                        "job_id": job_dir.name,
                        "model_path": str(best_models[0]),
                        "created_at": job_dir.stat().st_mtime,
                    })
    
    return {
        "status": "ok",
        "models": trained_models,
        "count": len(trained_models),
    }


@app.get("/training-jobs")
async def list_training_jobs() -> dict[str, Any]:
    """List all training jobs with their status"""
    jobs = []
    
    if TRAIN_OUTPUT_ROOT.exists():
        for job_dir in sorted(TRAIN_OUTPUT_ROOT.iterdir(), reverse=True):
            if job_dir.is_dir():
                job_info = {
                    "job_id": job_dir.name,
                    "created_at": datetime.fromtimestamp(job_dir.stat().st_mtime).isoformat(),
                    "status": "completed",
                    "has_model": False,
                    "log_file": None,
                }
                
                # Check if best model exists
                best_models = list(job_dir.glob("best_model_*.pth"))
                if best_models:
                    job_info["has_model"] = True
                
                # Check for log file
                log_file = job_dir / "train.log"
                if log_file.exists():
                    job_info["log_file"] = str(log_file)
                    # Check if training is still running
                    if not best_models:
                        job_info["status"] = "training"
                
                jobs.append(job_info)
    
    return {
        "status": "ok",
        "jobs": jobs,
        "count": len(jobs),
    }


@app.get("/health")
async def health() -> dict[str, Any]:
    return {
        "status": "ok",
        "voices": list(VOICE_CONFIGS.keys()),
    }


@app.post("/train", status_code=202)
async def start_training(payload: TrainRequest, background_tasks: BackgroundTasks) -> dict[str, Any]:
    dataset_dir, metadata_path = _validate_dataset(payload.dataset_name)

    timestamp = datetime.utcnow().strftime("%Y%m%d%H%M%S")
    job_id = f"{payload.dataset_name}_{timestamp}"
    output_dir = TRAIN_OUTPUT_ROOT / job_id
    output_dir.mkdir(parents=True, exist_ok=True)

    background_tasks.add_task(
        run_training_job,
        payload.model_dump(),
        dataset_dir,
        metadata_path,
        output_dir,
    )

    return {
        "status": "accepted",
        "job_id": job_id,
        "output_dir": str(output_dir),
    }


def _validate_dataset(dataset_name: str) -> tuple[Path, Path]:
    dataset_dir = DATA_ROOT / dataset_name
    metadata_path = dataset_dir / "metadata.csv"
    wav_dir = dataset_dir / "wavs"

    if not metadata_path.exists():
        raise HTTPException(status_code=400, detail=f"metadata.csv not found under {dataset_dir}")
    if not wav_dir.exists():
        raise HTTPException(status_code=400, detail=f"wavs directory not found under {dataset_dir}")

    return dataset_dir, metadata_path


def run_training_job(payload: dict[str, Any], dataset_dir: Path, metadata_path: Path, output_dir: Path) -> None:
    spec = TrainRequest(**payload)
    base_model_name = spec.base_model or BASE_TRAIN_MODEL

    logger.info("Starting fine-tune job '%s' with base model '%s'", output_dir.name, base_model_name)

    try:
        manager = ModelManager()
        pretrained_path, config_path, _ = manager.download_model(base_model_name)

        with open(config_path, "r", encoding="utf-8") as fh:
            config_data = json.load(fh)

        dataset_conf = (config_data.get("datasets") or [{}])[0]
        dataset_conf["path"] = str(dataset_dir)
        config_data = prepare_training_config(config_data, spec.dataset_name, metadata_path, dataset_dir)
        config_data["batch_size"] = spec.batch_size
        config_data["eval_batch_size"] = max(1, spec.batch_size // 2) or 1
        config_data["batch_group_size"] = config_data.get("batch_group_size") or 0
        config_data["num_loader_workers"] = config_data.get("num_loader_workers") or 0
        config_data["num_eval_loader_workers"] = config_data.get("num_eval_loader_workers") or 0
        config_data["output_path"] = str(output_dir)
        config_data["run_name"] = output_dir.name
        prepared_config_path = output_dir / "finetune_config.json"
        with open(prepared_config_path, "w", encoding="utf-8") as fh:
            json.dump(config_data, fh, indent=2)

        env = os.environ.copy()
        env.setdefault("RESTORE_PATH", pretrained_path)

        result_code, log_path = launch_training_job(prepared_config_path, output_dir, env)

        if result_code == 0:
            logger.info("Fine-tune job '%s' finished successfully", output_dir.name)
        else:
            logger.error("Fine-tune job '%s' failed. Check %s", output_dir.name, log_path)
    except Exception:  # pragma: no cover - background job failure
        logger.exception("Fine-tune job '%s' crashed", output_dir.name)
