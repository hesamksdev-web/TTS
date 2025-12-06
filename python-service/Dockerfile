FROM python:3.10-slim AS base

ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1 \
    PIP_NO_CACHE_DIR=1 \
    TORCHAUDIO_FORCE_OLD_AUDIO_IO=1

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
    build-essential \
    libsndfile1 \
    espeak-ng \
    cargo \
    rustc \
    git && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt ./
RUN pip install --upgrade pip setuptools wheel && \
    pip install --no-cache-dir 'numpy>=1.24.3,<2.0.0' 'scipy>=1.11.2' 'cython>=0.29.30' && \
    pip install --no-cache-dir fastapi==0.111.0 \
        uvicorn[standard]==0.30.3 \
        soundfile \
        'trainer>=0.0.32' \
        'transformers>=4.33.0' \
        'torch>=2.1.0' \
        'torchaudio>=2.1.0' \
        librosa>=0.10.0 \
        coqpit>=0.0.17 \
        pyyaml>=6.0 \
        fsspec>=2023.6.0 \
        aiohttp>=3.8.1 \
        packaging>=23.1 \
        tqdm>=4.64.1 \
        pysbd>=0.3.4 \
        matplotlib>=3.7.0 \
        anyascii>=0.3.0 \
        inflect>=7.0.0 \
        flask>=2.0.1 \
        umap-learn>=0.5.1 \
        'pandas>=1.4,<2.0' \
        'scikit-learn>=1.3.0' \
        'numba>=0.57.0' \
        bangla>=0.0.2 \
        bnnumerizer>=0.0.2 \
        bnunicodenormalizer>=0.1.4 \
        jieba>=0.42.1 \
        pypinyin>=0.47.1 \
        hangul_romanize>=0.1.0 \
        jamo>=0.4.1 \
        nltk>=3.8.1 \
        g2pkk>=0.1.1 \
        num2words>=0.5.12 \
        'gruut[de,es,fr]==2.2.3' \
        unidecode>=1.3.2 \
        'g2p-en>=2.1.0' \
        epitran>=1.15 \
        'gruut-ipa>=0.13.0' \
        einops>=0.6.0 \
        encodec>=0.1.1 \
        piper-phonemize>=1.1.0 \
        'spacy[ja]>=3.0.0' \
        --retries 5 && \
    pip install --no-cache-dir TTS==0.22.0 --no-deps

COPY app ./app

EXPOSE 5000

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "5000"]
