۱. اصلاح حیاتی: مشکل تغییر Pitch در تغییر سرعت
در حال حاضر شما برای تغییر سرعت از np.interp (در خط ۴۶۰) استفاده می‌کنید.

مشکل: این روش مثل "تند کردن نوار کاست" عمل می‌کند. یعنی اگر سرعت را زیاد کنید، صدا نازک (Chipmunk effect) و اگر کم کنید، صدا کلفت می‌شود.
راه حل: باید از الگوریتم Time Stretching استفاده کنید که سرعت را تغییر می‌دهد اما تُن صدا (Pitch) را ثابت نگه می‌دارد. خوشبختانه librosa این قابلیت را دارد.
جایگزین کردن کد در بخش _clone_voice (انتهای تابع):

Copy
# ... (کد قبلی شما)

# Apply speed adjustment if needed
if speed_float != 1.0:
    logger.info("Applying speed adjustment: %f (Time Stretch)", speed_float)
    wav_data, sr = sf.read(output_path)
    
    # روش غلط قبلی (باعث تغییر نازکی/کلفتی صدا می‌شد):
    # wav_data_adjusted = np.interp(...) 
    
    # روش صحیح (حفظ تن صدا):
    import librosa
    # نکته: librosa ورودی را float32 و (n,) می‌خواهد که sf.read معمولا می‌دهد
    # اگر سرعت > 1 باشد، صدا سریع‌تر و کوتاه تر می‌شود
    wav_data_adjusted = librosa.effects.time_stretch(wav_data, rate=speed_float)
    
    sf.write(output_path, wav_data_adjusted, sr)
    logger.info("Speed adjustment applied with pitch preservation")

# Final normalization
# ... (ادامه کد شما)
۲. استانداردسازی نرخ نمونه‌برداری (Sample Rate)
مدل XTTS v2 به صورت ذاتی با 24000Hz کار می‌کند.

در کد شما، ورودی کاربر (Voice Sample) به 22050Hz تبدیل می‌شود (خط ۳۵۴).
اما در تابع merge_audio_files خروجی روی 24000Hz تنظیم شده است.
بهترین کار: همه چیز را روی 24000 تنظیم کنید تا از تبدیل‌های مکرر (Resampling) که کیفیت را کم می‌کند جلوگیری شود.
تغییر در خط ۳۵۴ (داخل _validate_audio):

Copy
# تغییر از 22050 به 24000 برای هماهنگی با XTTS v2
if sr != 24000:
    import librosa
    audio_data = librosa.resample(audio_data, orig_sr=sr, target_sr=24000)

# Ensure float32 dtype
audio_data = audio_data.astype(np.float32)

# Save as WAV file
sf.write(wav_sample_path, audio_data, 24000)