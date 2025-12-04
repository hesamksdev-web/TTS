import logging
from pathlib import Path
from typing import Any, Dict, List

logger = logging.getLogger("tts_formatter")


def tts_service_two_column(root_path: str, meta_file: str, **_: Any) -> List[Dict[str, Any]]:
    dataset_root = Path(root_path)
    manifest_path = dataset_root / meta_file

    if not manifest_path.exists():
        raise FileNotFoundError(f"metadata file not found: {manifest_path}")

    samples: List[Dict[str, Any]] = []
    with open(manifest_path, "r", encoding="utf-8") as handle:
        for line in handle:
            stripped = line.strip()
            if not stripped or stripped.startswith("#"):
                continue

            cols = [segment.strip() for segment in stripped.split("|")]
            if len(cols) < 2:
                continue

            wav_file, text = cols[0], cols[1]
            audio_path = dataset_root / "wavs" / wav_file
            if not audio_path.exists():
                logger.warning("Skipping missing audio file: %s", audio_path)
                continue

            samples.append(
                {
                    "text": text,
                    "audio_file": str(audio_path),
                    "speaker_name": "tts_service",
                    "root_path": str(dataset_root),
                }
            )

    if not samples:
        raise ValueError(f"No valid samples found in {manifest_path}")

    return samples
