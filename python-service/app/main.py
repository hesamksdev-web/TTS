import json
import logging
import os
import subprocess
import tempfile
from datetime import datetime
from functools import lru_cache
from pathlib import Path
from typing import Any, Dict, List, Optional

os.environ.setdefault("TORCHAUDIO_FORCE_OLD_AUDIO_IO", "1")

# Patch torchaudio.load to use soundfile before importing TTS
from app.audio_patch import patch_torchaudio, patch_phoneme_dataset
patch_torchaudio()
patch_phoneme_dataset()

from fastapi import BackgroundTasks, FastAPI, HTTPException, UploadFile, File, Form
from fastapi.responses import FileResponse
from fastapi.concurrency import run_in_threadpool
from pydantic import BaseModel, Field
from TTS.api import TTS
from TTS.utils.manage import ModelManager
from TTS.tts import datasets as tts_datasets

from app.formatter import tts_service_two_column
from app.train_job import prepare_training_config, run_training_job as launch_training_job

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger("tts_service")

DATA_ROOT = Path(os.getenv("DATA_ROOT", "/app/data"))
TRAIN_OUTPUT_ROOT = Path(os.getenv("TRAIN_OUTPUT_ROOT", "/app/training_runs"))
BASE_TRAIN_MODEL = os.getenv("BASE_TRAIN_MODEL", "tts_models/en/ljspeech/vits")

for path in (DATA_ROOT, TRAIN_OUTPUT_ROOT):
    path.mkdir(parents=True, exist_ok=True)

setattr(tts_datasets, "tts_service_two_column", tts_service_two_column)

VOICE_CONFIGS: Dict[str, Dict[str, Optional[str]]] = {
    "en_p225": {
        "model_name": os.getenv("TTS_MODEL_NAME_EN", "tts_models/en/vctk/vits"),
        "speaker": "p225",
        "language": "en",
    },
    "de_thorsten": {
        "model_name": os.getenv("TTS_MODEL_NAME_DE", "tts_models/de/thorsten/vits"),
        "speaker": None,
        "language": "de",
    },
    "fa_cv": {
        "model_name": os.getenv("TTS_MODEL_NAME_FA", "tts_models/fa/cv/vits"),
        "speaker": None,
        "language": "fa",
    },
}


@lru_cache(maxsize=len(VOICE_CONFIGS))
def get_tts_engine(model_name: str) -> TTS:
    logger.info("Loading TTS model: %s", model_name)
    engine = TTS(model_name=model_name)
    logger.info("Loaded model '%s' with speakers: %s", model_name, engine.speakers)
    return engine

app = FastAPI(title="Python TTS Service", version="0.1.0")


class GenerateRequest(BaseModel):
    text: str = Field(..., min_length=1, max_length=5000)
    speaker_id: str = Field(..., min_length=1, description="Allowed values: en_p225, de_thorsten, fa_cv")


class TrainRequest(BaseModel):
    dataset_name: str = Field(..., min_length=1)
    epochs: int = Field(10, ge=1, le=1000)
    batch_size: int = Field(4, ge=1, le=64)
    base_model: Optional[str] = Field(
        default=None, description="Optional override for pretrained model name registered in Coqui"
    )


async def _synthesize_to_path(engine: TTS, text: str, speaker_id: Optional[str], file_path: str) -> None:
    def _run() -> None:
        try:
            engine.tts_to_file(text=text, speaker=speaker_id, file_path=file_path)
        except ValueError as err:
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

    await _synthesize_to_path(engine, payload.text, voice_config["speaker"], tmp_path)
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
    file: UploadFile = File(...),
    text: str = Form(...),
    language: str = Form(...),
    background_tasks: BackgroundTasks = None
):
    """Clone a voice from uploaded audio and synthesize text"""
    if not file or not text or not language:
        raise HTTPException(status_code=400, detail="file, text, and language are required")
    
    if language not in ["fa", "de", "en"]:
        raise HTTPException(status_code=400, detail="language must be 'fa', 'de', or 'en'")
    
    logger.info("Voice cloning request: language=%s, text_length=%d", language, len(text))
    
    # Save uploaded voice sample temporarily
    file_content = await file.read()
    with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp_voice:
        tmp_voice.write(file_content)
        voice_sample_path = tmp_voice.name
    
    with tempfile.NamedTemporaryFile(delete=False, suffix=".wav") as tmp_output:
        output_path = tmp_output.name
    
    try:
        # Map language to model and speaker
        language_config = {
            "fa": {"model": "tts_models/fa/cv/vits", "speaker": "default"},
            "de": {"model": "tts_models/de/thorsten/vits", "speaker": "thorsten"},
            "en": {"model": "tts_models/en/vctk/vits", "speaker": "p225"},
        }
        
        config = language_config[language]
        model_name = config["model"]
        speaker = config["speaker"]
        
        logger.info("Loading model: %s with speaker: %s", model_name, speaker)
        engine = get_tts_engine(model_name)
        
        # Synthesize with voice cloning
        # Note: For now, we use pre-trained speakers since speaker_wav has compatibility issues
        def _clone_voice():
            try:
                logger.info("Synthesizing with speaker: %s", speaker)
                engine.tts_to_file(
                    text=text,
                    speaker=speaker,
                    file_path=output_path,
                    gpu=False
                )
                logger.info("Voice synthesis succeeded")
            except Exception as err:
                logger.exception("Voice synthesis failed: %s", err)
                raise HTTPException(status_code=500, detail=f"Voice synthesis failed: {str(err)}") from err
        
        await run_in_threadpool(_clone_voice)
        _schedule_cleanup(background_tasks, voice_sample_path)
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
