import json
import os
import subprocess
from pathlib import Path
from typing import Any, Dict

from app.formatter import tts_service_two_column


def run_training_job(config_path: Path, output_dir: Path, env: Dict[str, str]) -> tuple[int, Path]:
    """Launch fine-tune process using custom formatter."""
    cmd = [
        "python3",
        "-m",
        "app.finetune_entry",
        "--config_path",
        str(config_path),
        "--output_path",
        str(output_dir),
        "--restore_path",
        env.get("RESTORE_PATH", ""),
        "--finetune",
    ]

    env = env.copy()
    env.setdefault("USE_CUDA", "0")
    env.setdefault("TORCHAUDIO_FORCE_OLD_AUDIO_IO", "1")
    existing_pythonpath = env.get("PYTHONPATH")
    env["PYTHONPATH"] = "/app" if not existing_pythonpath else f"/app:{existing_pythonpath}"

    log_file = output_dir / "train.log"
    with open(log_file, "w", encoding="utf-8") as fp:
        process = subprocess.run(cmd, cwd=str(output_dir), env=env, stdout=fp, stderr=fp)
    return process.returncode, log_file


def prepare_training_config(base_config: Dict[str, Any], dataset_name: str, metadata_file: Path, dataset_root: Path) -> Dict[str, Any]:
    config_data = json.loads(json.dumps(base_config))
    dataset_conf = (config_data.get("datasets") or [{}])[0]
    dataset_conf["path"] = str(dataset_root)
    dataset_conf["meta_file_train"] = metadata_file.name
    dataset_conf["meta_file_val"] = metadata_file.name
    dataset_conf["formatter"] = "tts_service_two_column"
    dataset_conf["language"] = dataset_conf.get("language", "en")
    dataset_conf["dataset_name"] = dataset_name
    dataset_conf.setdefault("ignored_speakers", None)
    dataset_conf.setdefault("cache_path", str(dataset_root / ".phoneme_cache"))
    config_data["datasets"] = [dataset_conf]

    config_data.setdefault("formatters", {})["tts_service_two_column"] = {
        "function": "tts_service_two_column",
        "module": "app.formatter",
    }

    config_data.setdefault("output_path", str(dataset_root / "training"))
    config_data.setdefault("run_name", f"{dataset_name}_finetune")
    config_data.setdefault("init_discriminator", True)
    # Set phoneme cache path at top level to ensure it's used by phoneme dataset
    config_data.setdefault("phoneme_cache_path", str(dataset_root / ".phoneme_cache"))
    # Disable evaluation to avoid empty tensor bug in TTS transforms during early training
    config_data.setdefault("eval_split_size", 0)

    discriminator_conf = config_data.setdefault("discriminator", {})
    discriminator_conf.setdefault("periods", [3, 5, 7, 11, 13])
    discriminator_conf.setdefault("use_spectral_norm", False)

    model_args = config_data.setdefault("model_args", {})
    model_args["init_discriminator"] = True
    model_args.setdefault("use_spectral_norm_disriminator", False)

    return config_data
