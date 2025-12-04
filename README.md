# TTS Inference Prototype

مرحله اول برای راه‌اندازی هسته TTS در این مخزن آماده شده است. این پروژه شامل دو سرویس داکری است:

1. **python-service**: سرویس FastAPI که از کتابخانه `coqui-tts` برای تولید صدا استفاده می‌کند و روی پورت 5000 گوش می‌دهد.
2. **go-service**: وب‌سرور Go که روی پورت 8080 قرار دارد و درخواست‌ها را به سرویس پایتون پروکسی می‌کند.

هر دو سرویس در شبکه مشترک `tts-network` اجرا می‌شوند تا بتوانند با هم ارتباط داشته باشند.

## اجرای پروژه

1. مطمئن شوید Docker و Docker Compose نصب شده‌اند.
2. در ریشه مخزن دستور زیر را اجرا کنید:

```bash
docker compose up --build
```

- سرویس پایتون روی `http://localhost:5000` در دسترس است.
- سرویس Go روی `http://localhost:8080` در دسترس است.

گوینده‌های فعلی:

- `en_p225`: انگلیسی (مدل `tts_models/en/vctk/vits` با اسپیکر `p225`)
- `de_thorsten`: آلمانی (مدل `tts_models/de/thorsten/vits`)
- `fa_cv`: فارسی (مدل `tts_models/fa/cv/vits`)

## تست سریع

```bash
curl -X POST \
  http://localhost:8080/api/v1/synthesize \
  -H 'Content-Type: application/json' \
  -o output.wav \
  -d '{"text":"Hello from Cascade!","speaker_id":"p225"}'
```

در صورت موفقیت، فایل `output.wav` حاوی صدای تولید شده خواهد بود.

## تنظیمات محیطی

- `TTS_MODEL_NAME_EN`: مدل انگلیسی (پیش‌فرض: `tts_models/en/vctk/vits`).
- `TTS_MODEL_NAME_DE`: مدل آلمانی (پیش‌فرض: `tts_models/de/thorsten/vits`).
- `TTS_MODEL_NAME_FA`: مدل فارسی (پیش‌فرض: `tts_models/fa/cv/vits`).
- `DATA_ROOT`: مسیر داده‌های خام و استاندارد شده (پیش‌فرض: `/app/data`).
- `TRAIN_OUTPUT_ROOT`: مسیر ذخیره خروجی آموزش (پیش‌فرض: `/app/training_runs`).
- `PYTHON_SERVICE_URL`: آدرس سرویس پایتون که توسط سرویس Go استفاده می‌شود (پیش‌فرض: `http://python-service:5000/generate`).
- `PYTHON_TRAIN_URL`: آدرس اندپوینت آموزش سرویس پایتون (پیش‌فرض: `http://python-service:5000/train`).
- `PORT`: پورتی که سرویس Go به آن گوش می‌دهد (پیش‌فرض: `8080`).

## آماده‌سازی دیتاست برای Fine-Tune

1. **استانداردسازی فایل‌ها** (از طریق اندپوینت Go):
   ```bash
   curl -X POST http://localhost:8080/api/v1/dataset/upload \
     -F dataset_name=my_dataset \
     -F transcript="Crisp crickets click by creek." \
     -F audio=@/path/to/en_002.mp3
   ```
   این درخواست فایل را با `ffmpeg` به WAV (22050Hz Mono) تبدیل کرده و خط متن را به `./data/<dataset>/metadata.csv` اضافه می‌کند.

2. **ساختار نهایی دیتاست**:
   ```
   /app/data/my_dataset/
   ├── metadata.csv        # فرمت: filename.wav|متن
   └── wavs/
       ├── en_001_xxx.wav
       └── ...
   ```

3. **شروع فاین‌تیون** (از طریق Go یا مستقیم پایتون):
   ```bash
   curl -X POST http://localhost:8080/api/v1/finetune/start \
     -H 'Content-Type: application/json' \
     -d '{"dataset_name":"my_dataset","epochs":4,"batch_size":2}'
   ```
   یا مستقیم:
   ```bash
   curl -X POST http://localhost:5001/train \
     -H 'Content-Type: application/json' \
     -d '{"dataset_name":"my_dataset","epochs":4,"batch_size":2}'
   ```

4. خروجی هر job در `./training_runs/<dataset>_<timestamp>/` ذخیره می‌شود و فایل `train.log` شامل وضعیت دقیق است.
