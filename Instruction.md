root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs -f
python-service  | /usr/local/lib/python3.10/site-packages/librosa/core/intervals.py:8: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   from pkg_resources import resource_filename
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  | INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  |  > Downloading model to /root/.local/share/tts/tts_models--multilingual--multi-dataset--your_tts
python-service  |  > Model's license - CC BY-NC-ND 4.0
python-service  |  > Check https://creativecommons.org/licenses/by-nc-nd/4.0/ for more info.
python-service  |  > Using model: vits
python-service  |  > Setting up Audio Processor...
python-service  |  | > sample_rate:16000
python-service  |  | > resample:False
python-service  |  | > num_mels:80
python-service  |  | > log_func:np.log10
python-service  |  | > min_level_db:0
python-service  |  | > frame_shift_ms:None
python-service  |  | > frame_length_ms:None
python-service  |  | > ref_level_db:None
python-service  |  | > fft_size:1024
python-service  |  | > power:None
python-service  |  | > preemphasis:0.0
python-service  |  | > griffin_lim_iters:None
python-service  |  | > signal_norm:None
python-service  |  | > symmetric_norm:None
python-service  |  | > mel_fmin:0
python-service  |  | > mel_fmax:None
python-service  |  | > pitch_fmin:None
python-service  |  | > pitch_fmax:None
python-service  |  | > spec_gain:20.0
python-service  |  | > stft_pad_mode:reflect
python-service  |  | > max_norm:1.0
python-service  |  | > clip_norm:True
python-service  |  | > do_trim_silence:False
python-service  |  | > trim_db:60
python-service  |  | > do_sound_norm:False
python-service  |  | > do_amp_to_db_linear:True
python-service  |  | > do_amp_to_db_mel:True
python-service  |  | > do_rms_norm:False
python-service  |  | > db_level:None
python-service  |  | > stats_path:None
python-service  |  | > base:10
python-service  |  | > hop_length:256
python-service  |  | > win_length:1024
python-service  |  > Model fully restored.
python-service  |  > Setting up Audio Processor...
python-service  |  | > sample_rate:16000
python-service  |  | > resample:False
python-service  |  | > num_mels:64
python-service  |  | > log_func:np.log10
python-service  |  | > min_level_db:-100
python-service  |  | > frame_shift_ms:None
python-service  |  | > frame_length_ms:None
python-service  |  | > ref_level_db:20
python-service  |  | > fft_size:512
python-service  |  | > power:1.5
python-service  |  | > preemphasis:0.97
python-service  |  | > griffin_lim_iters:60
python-service  |  | > signal_norm:False
python-service  |  | > symmetric_norm:False
python-service  |  | > mel_fmin:0
python-service  |  | > mel_fmax:8000.0
python-service  |  | > pitch_fmin:1.0
python-service  |  | > pitch_fmax:640.0
python-service  |  | > spec_gain:20.0
python-service  |  | > stft_pad_mode:reflect
python-service  |  | > max_norm:4.0
caddy-proxy     | {"level":"info","ts":1764942046.293748,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy     | {"level":"info","ts":1764942046.2939353,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1764942046.2940297,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1764942046.2953532,"msg":"adapted config to JSON","adapter":"caddyfile"}
tts-frontend    |
caddy-proxy     | {"level":"warn","ts":1764942046.2954898,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1764942046.2959268,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy     | {"level":"info","ts":1764942046.296195,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc00060c880"}
caddy-proxy     | {"level":"info","ts":1764942046.296123,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
python-service  |  | > clip_norm:False
python-service  |  | > do_trim_silence:False
python-service  |  | > trim_db:60
python-service  |  | > do_sound_norm:False
python-service  |  | > do_amp_to_db_linear:True
python-service  |  | > do_amp_to_db_mel:True
python-service  |  | > do_rms_norm:True
python-service  |  | > db_level:-27.0
python-service  |  | > stats_path:None
python-service  |  | > base:10
python-service  |  | > hop_length:160
python-service  |  | > win_length:400
python-service  |  > External Speaker Encoder Loaded !!
python-service  |  > initialization of language-embedding layers.
caddy-proxy     | {"level":"warn","ts":1764942046.29684,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"warn","ts":1764942046.2969632,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"info","ts":1764942046.2969694,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy     | {"level":"info","ts":1764942046.297227,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1764942046.2975073,"msg":"serving initial configuration"}
caddy-proxy     | {"level":"info","ts":1764942046.299053,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765028446.2990494,"try_again_in":86399.999999279}
caddy-proxy     | {"level":"info","ts":1764942046.2992718,"logger":"tls","msg":"finished cleaning storage units"}
go-service      | 2025/12/05 13:40:45 Go Service (SaaS Edition) listening on :8080
go-service      | 2025/12/05 13:42:30 Python service response: status=500, body=Internal Server Error
python-service  |  > Model fully restored.
python-service  |  > Setting up Audio Processor...
python-service  |  | > sample_rate:16000
python-service  |  | > resample:False
python-service  |  | > num_mels:64
python-service  |  | > log_func:np.log10
python-service  |  | > min_level_db:-100
python-service  |  | > frame_shift_ms:None
python-service  |  | > frame_length_ms:None
python-service  |  | > ref_level_db:20
python-service  |  | > fft_size:512
python-service  |  | > power:1.5
python-service  |  | > preemphasis:0.97
python-service  |  | > griffin_lim_iters:60
python-service  |  | > signal_norm:False
python-service  |  | > symmetric_norm:False
python-service  |  | > mel_fmin:0
python-service  |  | > mel_fmax:8000.0
python-service  |  | > pitch_fmin:1.0
python-service  |  | > pitch_fmax:640.0
python-service  |  | > spec_gain:20.0
python-service  |  | > stft_pad_mode:reflect
python-service  |  | > max_norm:4.0
python-service  |  | > clip_norm:False
python-service  |  | > do_trim_silence:False
python-service  |  | > trim_db:60
python-service  |  | > do_sound_norm:False
python-service  |  | > do_amp_to_db_linear:True
python-service  |  | > do_amp_to_db_mel:True
python-service  |  | > do_rms_norm:True
python-service  |  | > db_level:-27.0
python-service  |  | > stats_path:None
python-service  |  | > base:10
python-service  |  | > hop_length:160
python-service  |  | > win_length:400
python-service  |  > Text splitted to sentences.
python-service  | ['hi i am mehrshad yadkuri and this is test']
python-service  |  > Processing time: 2.0611979961395264
python-service  |  > Real-time factor: 0.6638318828146623
python-service  | INFO:     172.18.0.4:38092 - "POST /voice-clone HTTP/1.1" 200 OK
python-service  | INFO:     172.18.0.4:40824 - "GET /trained-models HTTP/1.1" 200 OK
python-service  |  > You must confirm the following:
python-service  |  | > "I have purchased a commercial license from Coqui: licensing@coqui.ai"
python-service  |  | > "Otherwise, I agree to the terms of the non-commercial CPML: https://coqui.ai/cpml" - [y/n]
python-service  |  | | > INFO:     172.18.0.4:40824 - "POST /tts HTTP/1.1" 500 Internal Server Error
100%|██████████| 425M/425M [00:19<00:00, 77.5MiB/s]ERROR:    Exception in ASGI application
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/protocols/http/httptools_impl.py", line 399, in run_asgi
python-service  |     result = await app(  # type: ignore[func-returns-value]
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/middleware/proxy_headers.py", line 70, in __call__
python-service  |     return await self.app(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/fastapi/applications.py", line 1054, in __call__
python-service  |     await super().__call__(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/applications.py", line 123, in __call__
python-service  |     await self.middleware_stack(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/middleware/errors.py", line 186, in __call__
python-service  |     raise exc
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/middleware/errors.py", line 164, in __call__
python-service  |     await self.app(scope, receive, _send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/middleware/exceptions.py", line 65, in __call__
python-service  |     await wrap_app_handling_exceptions(self.app, conn)(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/_exception_handler.py", line 64, in wrapped_app
python-service  |     raise exc
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/_exception_handler.py", line 53, in wrapped_app
python-service  |     await app(scope, receive, sender)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/routing.py", line 756, in __call__
python-service  |     await self.middleware_stack(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/routing.py", line 776, in app
python-service  |     await route.handle(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/routing.py", line 297, in handle
python-service  |     await self.app(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/routing.py", line 77, in app
python-service  |     await wrap_app_handling_exceptions(app, request)(scope, receive, send)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/_exception_handler.py", line 64, in wrapped_app
python-service  |     raise exc
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/_exception_handler.py", line 53, in wrapped_app
python-service  |     await app(scope, receive, sender)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/routing.py", line 72, in app
python-service  |     response = await func(request)
python-service  |   File "/usr/local/lib/python3.10/site-packages/fastapi/routing.py", line 278, in app
python-service  |     raw_response = await run_endpoint_function(
python-service  |   File "/usr/local/lib/python3.10/site-packages/fastapi/routing.py", line 191, in run_endpoint_function
python-service  |     return await dependant.call(**values)
python-service  |   File "/app/app/main.py", line 130, in tts_endpoint
python-service  |     return await generate_audio(payload, background_tasks)
python-service  |   File "/app/app/main.py", line 115, in generate_audio
python-service  |     engine = get_tts_engine(voice_config["model_name"])
python-service  |   File "/app/app/main.py", line 63, in get_tts_engine
python-service  |     engine = TTS(model_name=model_name)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 74, in __init__
python-service  |     self.load_tts_model_by_name(model_name, gpu)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 171, in load_tts_model_by_name
python-service  |     model_path, config_path, vocoder_path, vocoder_config_path, model_dir = self.download_model_by_name(
tts-frontend    | > frontend@0.1.0 start
tts-frontend    | > next start
tts-frontend    |
tts-frontend    |    ▲ Next.js 16.0.6
tts-frontend    |    - Local:         http://localhost:3000
tts-frontend    |    - Network:       http://172.18.0.5:3000
tts-frontend    |
tts-frontend    |  ✓ Starting...
tts-frontend    |  ✓ Ready in 314ms
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 129, in download_model_by_name
python-service  |     model_path, config_path, model_item = self.manager.download_model(model_name)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/manage.py", line 411, in download_model
python-service  |     self.create_dir_and_download_model(model_name, model_item, output_path)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/manage.py", line 337, in create_dir_and_download_model
tts-db          |
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db          |
tts-db          | 2025-12-05 13:40:35.256 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 14.2.0) 14.2.0, 64-bit
tts-db          | 2025-12-05 13:40:35.257 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-05 13:40:35.257 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-05 13:40:35.260 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-05 13:40:35.265 UTC [29] LOG:  database system was shut down at 2025-12-05 13:40:02 UTC
tts-db          | 2025-12-05 13:40:35.274 UTC [1] LOG:  database system is ready to accept connections
python-service  |     if not self.ask_tos(output_path):
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/manage.py", line 316, in ask_tos
python-service  |     answer = input(" | | > ")
python-service  | EOFError: EOF when reading a line