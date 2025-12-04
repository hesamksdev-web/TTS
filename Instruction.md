root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs
tts-db  |
tts-db  | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db  |
tts-db  | 2025-12-04 15:54:31.645 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 14.2.0) 14.2.0, 64-bit
tts-db  | 2025-12-04 15:54:31.645 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db  | 2025-12-04 15:54:31.645 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db  | 2025-12-04 15:54:31.648 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db  | 2025-12-04 15:54:31.654 UTC [28] LOG:  database system was shut down at 2025-12-04 15:54:24 UTC
tts-db  | 2025-12-04 15:54:31.662 UTC [1] LOG:  database system is ready to accept connections
tts-frontend  |
tts-frontend  | > frontend@0.1.0 start
python-service  | /usr/local/lib/python3.10/site-packages/librosa/core/intervals.py:8: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   from pkg_resources import resource_filename
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  | INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  |  > Downloading model to /root/.local/share/tts/tts_models--en--vctk--vits
python-service  |  > Model's license - apache 2.0
python-service  |  > Check https://choosealicense.com/licenses/apache-2.0/ for more info.
python-service  |  > Using model: vits
python-service  |  > Setting up Audio Processor...
python-service  |  | > sample_rate:22050
python-service  |  | > resample:False
python-service  |  | > num_mels:80
python-service  |  | > log_func:np.log10
python-service  |  | > min_level_db:0
python-service  |  | > frame_shift_ms:None
python-service  |  | > frame_length_ms:None
tts-frontend    | > next start
python-service  |  | > ref_level_db:None
python-service  |  | > fft_size:1024
tts-frontend    |
tts-frontend    |    ▲ Next.js 16.0.6
tts-frontend    |    - Local:         http://localhost:3000
tts-frontend    |    - Network:       http://172.18.0.5:3000
tts-frontend    |
tts-frontend    |  ✓ Starting...
tts-frontend    |  ✓ Ready in 343ms
python-service  |  | > power:None
python-service  |  | > preemphasis:0.0
go-service    | 2025/12/04 15:54:42 Go Service (SaaS Edition) listening on :8080
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
caddy-proxy     | {"level":"info","ts":1764863682.7844183,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy     | {"level":"info","ts":1764863682.7845788,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1764863682.7846112,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1764863682.7858505,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy     | {"level":"warn","ts":1764863682.7858849,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1764863682.7862172,"logger":"admin","msg":"admin endpoint disabled"}
python-service  |  | > base:10
python-service  |  | > hop_length:256
python-service  |  | > win_length:1024
python-service  |  > initialization of speaker-embedding layers.
python-service  |  > Text splitted to sentences.
python-service  | ['hi this is test']
 96%|█████████▋| 142M/148M [00:01<00:00, 83.9MiB/s]ERROR:tts_service:Voice cloning with speaker_wav failed: conv1d() received an invalid combination of arguments - got (NoneType, Parameter, Parameter, tuple, tuple, tuple, int), but expected one of:
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, tuple of ints padding = 0, tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
caddy-proxy     | {"level":"info","ts":1764863682.786398,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy     | {"level":"info","ts":1764863682.7864742,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc000151200"}
caddy-proxy     | {"level":"warn","ts":1764863682.7869244,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"warn","ts":1764863682.7869563,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"info","ts":1764863682.78697,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy     | {"level":"info","ts":1764863682.7872524,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1764863682.7872918,"msg":"serving initial configuration"}
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, str padding = "valid", tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
python-service  | Traceback (most recent call last):
python-service  |   File "/app/app/main.py", line 219, in _clone_voice
python-service  |     engine.tts_to_file(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 334, in tts_to_file
python-service  |     wav = self.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 276, in tts
python-service  |     wav = self.synthesizer.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/synthesizer.py", line 398, in tts
python-service  |     outputs = synthesis(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/utils/synthesis.py", line 221, in synthesis
python-service  |     outputs = run_model_torch(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/utils/synthesis.py", line 53, in run_model_torch
python-service  |     outputs = _func(
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/utils/_contextlib.py", line 120, in decorate_context
python-service  |     return func(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py", line 1161, in inference
python-service  |     o = self.waveform_decoder((z * y_mask)[:, :, : self.max_inference_len], g=g)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1775, in _wrapped_call_impl
python-service  |     return self._call_impl(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1786, in _call_impl
python-service  |     return forward_call(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/vocoder/models/hifigan_generator.py", line 251, in forward
python-service  |     o = o + self.cond_layer(g)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1775, in _wrapped_call_impl
python-service  |     return self._call_impl(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1786, in _call_impl
python-service  |     return forward_call(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/conv.py", line 371, in forward
python-service  |     return self._conv_forward(input, self.weight, self.bias)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/conv.py", line 366, in _conv_forward
python-service  |     return F.conv1d(
python-service  | TypeError: conv1d() received an invalid combination of arguments - got (NoneType, Parameter, Parameter, tuple, tuple, tuple, int), but expected one of:
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, tuple of ints padding = 0, tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, str padding = "valid", tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
python-service  |
python-service  | ERROR:tts_service:Voice cloning failed
python-service  | Traceback (most recent call last):
python-service  |   File "/app/app/main.py", line 219, in _clone_voice
python-service  |     engine.tts_to_file(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 334, in tts_to_file
python-service  |     wav = self.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 276, in tts
python-service  |     wav = self.synthesizer.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/synthesizer.py", line 398, in tts
python-service  |     outputs = synthesis(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/utils/synthesis.py", line 221, in synthesis
python-service  |     outputs = run_model_torch(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/utils/synthesis.py", line 53, in run_model_torch
python-service  |     outputs = _func(
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/utils/_contextlib.py", line 120, in decorate_context
python-service  |     return func(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py", line 1161, in inference
python-service  |     o = self.waveform_decoder((z * y_mask)[:, :, : self.max_inference_len], g=g)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1775, in _wrapped_call_impl
python-service  |     return self._call_impl(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1786, in _call_impl
python-service  |     return forward_call(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/vocoder/models/hifigan_generator.py", line 251, in forward
caddy-proxy     | {"level":"info","ts":1764863682.7895477,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1764950082.7895436,"try_again_in":86399.999999649}
caddy-proxy     | {"level":"info","ts":1764863682.7897065,"logger":"tls","msg":"finished cleaning storage units"}
python-service  |     o = o + self.cond_layer(g)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1775, in _wrapped_call_impl
python-service  |     return self._call_impl(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/module.py", line 1786, in _call_impl
python-service  |     return forward_call(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/conv.py", line 371, in forward
python-service  |     return self._conv_forward(input, self.weight, self.bias)
python-service  |   File "/usr/local/lib/python3.10/site-packages/torch/nn/modules/conv.py", line 366, in _conv_forward
python-service  |     return F.conv1d(
python-service  | TypeError: conv1d() received an invalid combination of arguments - got (NoneType, Parameter, Parameter, tuple, tuple, tuple, int), but expected one of:
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, tuple of ints padding = 0, tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, str padding = "valid", tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
python-service  |
python-service  |
python-service  | The above exception was the direct cause of the following exception:
python-service  |
python-service  | Traceback (most recent call last):
python-service  |   File "/app/app/main.py", line 230, in voice_clone
python-service  |     await run_in_threadpool(_clone_voice)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/concurrency.py", line 42, in run_in_threadpool
python-service  |     return await anyio.to_thread.run_sync(func, *args)
python-service  |   File "/usr/local/lib/python3.10/site-packages/anyio/to_thread.py", line 61, in run_sync
python-service  |     return await get_async_backend().run_sync_in_worker_thread(
python-service  |   File "/usr/local/lib/python3.10/site-packages/anyio/_backends/_asyncio.py", line 2525, in run_sync_in_worker_thread
python-service  |     return await future
python-service  |   File "/usr/local/lib/python3.10/site-packages/anyio/_backends/_asyncio.py", line 986, in run
python-service  |     result = context.run(func, *args)
python-service  |   File "/app/app/main.py", line 228, in _clone_voice
python-service  |     raise HTTPException(status_code=500, detail=f"Voice cloning failed: {str(err)}") from err
python-service  | fastapi.exceptions.HTTPException: 500: Voice cloning failed: conv1d() received an invalid combination of arguments - got (NoneType, Parameter, Parameter, tuple, tuple, tuple, int), but expected one of:
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, tuple of ints padding = 0, tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
python-service  |  * (Tensor input, Tensor weight, Tensor bias = None, tuple of ints stride = 1, str padding = "valid", tuple of ints dilation = 1, int groups = 1)
python-service  |       didn't match because some of the arguments have invalid types: (!NoneType!, !Parameter!, !Parameter!, !tuple of (int,)!, !tuple of (int,)!, !tuple of (int,)!, !int!)
python-service  |
python-service  | INFO:     172.18.0.4:58986 - "POST /voice-clone HTTP/1.1" 500 Internal Server Error
root@Automations:/opt/TTS#