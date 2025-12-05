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
    rustc && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY requirements.txt ./
RUN pip install --upgrade pip setuptools wheel && \
    pip install --no-cache-dir 'numpy<2.0.0' && \
    pip install --no-cache-dir -r requirements.txt --retries 5 && \
    pip install --no-cache-dir 'transformers>=4.30.0' --force-reinstall

COPY app ./app

EXPOSE 5000

CMD ["uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "5000"]
