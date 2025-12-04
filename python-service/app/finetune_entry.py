import os
os.environ.setdefault("TORCHAUDIO_FORCE_OLD_AUDIO_IO", "1")

# Patch torchaudio.load to use soundfile before importing TTS
from app.audio_patch import patch_torchaudio, patch_phoneme_dataset, patch_transforms
patch_torchaudio()
patch_phoneme_dataset()
patch_transforms()

from TTS.tts import datasets as tts_datasets
from app.formatter import tts_service_two_column

# Register custom formatter before invoking Coqui's CLI
setattr(tts_datasets, "tts_service_two_column", tts_service_two_column)

from TTS.bin.train_tts import main  # noqa: E402  (import after registration)


if __name__ == "__main__":
    main()
