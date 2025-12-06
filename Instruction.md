۱. حل مشکل تغییر لحن (Latent Consistency)
در تابع voice_clone (و سایر توابع تولید)، شما برای هر چانک فایل صوتی مرجع را دوباره به مدل می‌دهید. این کار باعث می‌شود مدل برای هر جمله، استایل کمی متفاوتی بگیرد.

راه حل: باید قبل از حلقه (Loop)، یک‌بار Latents را استخراج کنید و همان را به همه چانک‌ها پاس دهید.

تغییرات در app.py (داخل تابع _clone_voice):

Copy
# ... (داخل تابع _clone_voice)

# ۱. استخراج ویژگی‌های صدا فقط یک‌بار قبل از شروع پردازش متن
logger.info("Computing speaker latents...")
gpt_cond_latent, speaker_embedding = engine.get_conditioning_latents(
    audio_path=[wav_sample_path], 
    gpt_cond_len=30, 
    max_ref_length=60
)

for idx, chunk in enumerate(text_chunks):
    chunk_output = output_path.replace(".wav", f"_chunk_{idx}.wav")
    
    # ۲. ارسال Latentهای ثابت به جای فایل صوتی
    engine.tts_to_file(
        text=chunk,
        language=tts_language,
        gpt_cond_latent=gpt_cond_latent,    # <--- پارامتر جدید
        speaker_embedding=speaker_embedding, # <--- پارامتر جدید
        file_path=chunk_output,
        temperature=temperature_float,
        top_p=top_p_float
    )
    chunk_audio_paths.append(chunk_output)
۲. بازنویسی تابع merge_audio_files (حذف Crossfade مخرب)
استفاده از Crossfade (محو شدن صداها در هم) برای موسیقی عالی است، اما برای کلام (Speech) خطرناک است چون انتهای کلمات را حذف می‌کند.

راه حل: استراتژی "Trim & Pad". یعنی سکوت اضافه را ببرید و یک سکوت استاندارد (مثلاً ۳۰۰ میلی‌ثانیه) تزریق کنید.

این نسخه اصلاح شده تابع merge_audio_files را جایگزین کنید:

Copy
def merge_audio_files(audio_paths: List[str], output_path: str, silence_duration: float = 0.25):
    """
    Merge audio files by trimming silence and adding fixed natural pauses.
    Replaces Crossfade with Smart Concatenation.
    """
    import soundfile as sf
    import numpy as np
    import librosa # نیاز است به نیازمندی‌ها اضافه شود

    if not audio_paths:
        raise ValueError("No audio files to merge")

    merged_segments = []
    target_sr = 24000  # نرخ نمونه‌برداری استاندارد XTTS

    # ایجاد سکوت استاندارد (مثلاً ۲۵۰ میلی‌ثانیه برای مکث طبیعی بین جملات)
    silence_samples = int(silence_duration * target_sr)
    silence_array = np.zeros(silence_samples, dtype=np.float32)

    for i, path in enumerate(audio_paths):
        # بارگذاری با librosa برای مدیریت راحت‌تر نرخ نمونه‌برداری
        y, sr = librosa.load(path, sr=target_sr)
        
        # ۱. حذف سکوت‌های طولانی از ابتدا و انتهای هر چانک
        # top_db=30 یعنی صداهای زیر ۳۰ دسی‌بل سکوت تلقی می‌شوند
        y_trimmed, _ = librosa.effects.trim(y, top_db=30)
        
        merged_segments.append(y_trimmed)
        
        # اضافه کردن سکوت بین چانک‌ها (به جز چانک آخر)
        if i < len(audio_paths) - 1:
            merged_segments.append(silence_array)

    # اتصال نهایی
    final_audio = np.concatenate(merged_segments)

    # نرمال‌سازی صدا (جلوگیری از دیستورشن)
    max_val = np.max(np.abs(final_audio))
    if max_val > 0:
        final_audio = final_audio / max_val * 0.95

    sf.write(output_path, final_audio, target_sr)
    logger.info(f"Merged {len(audio_paths)} files with smart trimming.")
۳. بهبود split_text_into_chunks (جلوگیری از برش بد)
کد فعلی شما خوب است، اما یک نکته ظریف دارد. در زبان فارسی و انگلیسی، بهتر است برش روی ، (ویرگول) با برش روی . (نقطه) رفتار متفاوتی داشته باشد، اما چون ما چانک‌ها را جدا می‌فرستیم، مدل همیشه در انتهای چانک "لحن پایان جمله" (Falling Intonation) می‌گیرد.

پیشنهاد: سعی کنید تا جای ممکن فقط روی نقطه، علامت سوال و تعجب برش بزنید و max_chars را کمی افزایش دهید (مثلاً به ۳۵۰ برسانید، XTTS تا ۴۰۰ را هم خوب مدیریت می‌کند اگر RAM GPU اجازه دهد).

در تابع split_text_into_chunks:

Copy
# تغییر مقدار پیش‌فرض به عدد بالاتر برای کاهش تعداد برش‌ها
def split_text_into_chunks(text: str, max_chars: int = 350, language: str = "en") -> List[str]:
    # ... بقیه کد ...
