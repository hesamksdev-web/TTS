root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs
tts-db  |
tts-db  | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db  |
tts-db  | 2025-12-06 16:48:04.454 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 15.2.0) 15.2.0, 64-bit
tts-db  | 2025-12-06 16:48:04.454 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db  | 2025-12-06 16:48:04.454 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db  | 2025-12-06 16:48:04.459 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db  | 2025-12-06 16:48:04.465 UTC [29] LOG:  database system was shut down at 2025-12-06 16:31:51 UTC
tts-db  | 2025-12-06 16:48:04.477 UTC [1] LOG:  database system is ready to accept connections
tts-db  | 2025-12-06 16:53:04.514 UTC [27] LOG:  checkpoint starting: time
tts-db  | 2025-12-06 16:53:04.526 UTC [27] LOG:  checkpoint complete: wrote 3 buffers (0.0%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.003 s, sync=0.002 s, total=0.013 s; sync files=2, longest=0.001 s, average=0.001 s; distance=0 kB, estimate=0 kB
caddy-proxy  | {"level":"info","ts":1765039767.8417764,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy  | {"level":"info","ts":1765039767.8421283,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy  | {"level":"info","ts":1765039767.8425238,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy  | {"level":"info","ts":1765039767.847033,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy  | {"level":"warn","ts":1765039767.847059,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy  | {"level":"warn","ts":1765039767.847731,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy  | {"level":"info","ts":1765039767.849134,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy  | {"level":"info","ts":1765039767.849416,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc0005be180"}
caddy-proxy  | {"level":"warn","ts":1765039767.8534634,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy  | {"level":"warn","ts":1765039767.8535082,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy  | {"level":"info","ts":1765039767.8535156,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy  | {"level":"info","ts":1765039767.8544347,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy  | {"level":"info","ts":1765039767.8544602,"msg":"serving initial configuration"}
caddy-proxy  | {"level":"info","ts":1765039767.8581247,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765126167.8581223,"try_again_in":86399.99999933}
caddy-proxy  | {"level":"info","ts":1765039767.8582346,"logger":"tls","msg":"finished cleaning storage units"}
go-service   | 2025/12/06 16:49:21 Voice clone worker started
tts-frontend  |
tts-frontend  | > frontend@0.1.0 start
tts-frontend  | > next start
tts-frontend  |
tts-frontend  |    ▲ Next.js 16.0.6
tts-frontend  |    - Local:         http://localhost:3000
tts-frontend  |    - Network:       http://172.18.0.5:3000
tts-frontend  |
tts-frontend  |  ✓ Starting...
tts-frontend  |  ✓ Ready in 525ms
go-service    | 2025/12/06 16:49:21 Go Service (SaaS Edition) listening on :8080
go-service    |
go-service    | 2025/12/06 16:49:21 /app/main.go:656 record not found
go-service    | [1.734ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service    |
go-service    | 2025/12/06 16:49:23 /app/main.go:656 record not found
go-service    | [0.646ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service    |
go-service    | 2025/12/06 16:49:25 /app/main.go:656 record not found
go-service    | [0.812ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service    |
go-service    | 2025/12/06 16:49:27 /app/main.go:656 record not found
go-service    | [0.650ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service    |
python-service  | /usr/local/lib/python3.10/site-packages/jieba/_compat.py:18: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   import pkg_resources
go-service      | 2025/12/06 16:49:29 /app/main.go:656 record not found
python-service  | INFO:     Started server process [1]
go-service      | [0.883ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:31 /app/main.go:656 record not found
go-service      | [0.840ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:33 /app/main.go:656 record not found
go-service      | [0.596ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:35 /app/main.go:656 record not found
go-service      | [0.444ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:37 /app/main.go:656 record not found
go-service      | [0.520ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:39 /app/main.go:656 record not found
go-service      | [0.482ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:41 /app/main.go:656 record not found
go-service      | [0.585ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     Waiting for application startup.
python-service  |  > Downloading model to /root/.local/share/tts/tts/tts_models--multilingual--multi-dataset--xtts_v2
go-service      |
go-service      | 2025/12/06 16:49:43 /app/main.go:656 record not found
go-service      | [0.370ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:45 /app/main.go:656 record not found
go-service      | [0.707ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:47 /app/main.go:656 record not found
go-service      | [0.540ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
100%|█████████▉| 1.86G/1.87G [00:22<00:00, 85.4MiB/s]
go-service      |
100%|██████████| 1.87G/1.87G [00:22<00:00, 81.8MiB/s]6 record not found
4.37kiB [00:00, 19.5kiB/s]
361kiB [00:00, 1.78MiB/s]
100%|██████████| 32.0/32.0 [00:00<00:00, 138iB/s]
python-service  |  > Model's license - CPML
python-service  |  > Check https://coqui.ai/cpml.txt for more info.
python-service  |  > Using model: xtts
100%|██████████| 7.75M/7.75M [00:16<00:00, 60.7MiB/s]INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
go-service      | [0.467ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:51 /app/main.go:656 record not found
go-service      | [0.579ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:53 /app/main.go:656 record not found
go-service      | [0.680ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:55 /app/main.go:656 record not found
go-service      | [0.318ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:57 /app/main.go:656 record not found
go-service      | [0.667ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:49:59 /app/main.go:656 record not found
go-service      | [0.588ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:01 /app/main.go:656 record not found
go-service      | [0.561ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:03 /app/main.go:656 record not found
go-service      | [0.608ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:05 /app/main.go:656 record not found
go-service      | [0.533ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:07 /app/main.go:656 record not found
go-service      | [0.533ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:09 /app/main.go:656 record not found
python-service  | INFO:     127.0.0.1:37740 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40704 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35616 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:57016 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46378 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52806 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54002 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38998 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60462 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:41918 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:44712 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35434 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40922 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43304 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:51346 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37086 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40876 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:41316 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:39434 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56404 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58694 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:39682 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54864 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:44670 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45816 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58374 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60912 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40616 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54492 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46532 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46490 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40070 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43752 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Text splitted to sentences.
python-service  | ['And so the town continues to breathe through stories, one letter at a time, a little louder, a little brighter, the sea outside listening like a patient confidant.']
python-service  | ERROR:tts_service:Voice synthesis failed with speaker_wav:  ❗ Language fa is not supported. Supported languages are ['en', 'es', 'fr', 'de', 'it', 'pt', 'pl', 'tr', 'ru', 'nl', 'cs', 'ar', 'zh-cn', 'hu', 'ko', 'ja', 'hi']
python-service  | Traceback (most recent call last):
python-service  |   File "/app/app/main.py", line 611, in _clone_voice
python-service  |     engine.tts_to_file(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 334, in tts_to_file
python-service  |     wav = self.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 276, in tts
python-service  |     wav = self.synthesizer.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/synthesizer.py", line 386, in tts
python-service  |     outputs = self.tts_model.synthesize(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/models/xtts.py", line 399, in synthesize
python-service  |     "zh-cn" if language == "zh" else language in self.config.languages
python-service  | AssertionError:  ❗ Language fa is not supported. Supported languages are ['en', 'es', 'fr', 'de', 'it', 'pt', 'pl', 'tr', 'ru', 'nl', 'cs', 'ar', 'zh-cn', 'hu', 'ko', 'ja', 'hi']
python-service  | ERROR:tts_service:Voice cloning failed
python-service  | Traceback (most recent call last):
python-service  |   File "/app/app/main.py", line 611, in _clone_voice
python-service  |     engine.tts_to_file(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 334, in tts_to_file
python-service  |     wav = self.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 276, in tts
python-service  |     wav = self.synthesizer.tts(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/synthesizer.py", line 386, in tts
python-service  |     outputs = self.tts_model.synthesize(
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/models/xtts.py", line 399, in synthesize
python-service  |     "zh-cn" if language == "zh" else language in self.config.languages
python-service  | AssertionError:  ❗ Language fa is not supported. Supported languages are ['en', 'es', 'fr', 'de', 'it', 'pt', 'pl', 'tr', 'ru', 'nl', 'cs', 'ar', 'zh-cn', 'hu', 'ko', 'ja', 'hi']
python-service  |
python-service  | The above exception was the direct cause of the following exception:
python-service  |
python-service  | Traceback (most recent call last):
python-service  |   File "/app/app/main.py", line 665, in voice_clone
python-service  |     await run_in_threadpool(_clone_voice)
python-service  |   File "/usr/local/lib/python3.10/site-packages/starlette/concurrency.py", line 42, in run_in_threadpool
python-service  |     return await anyio.to_thread.run_sync(func, *args)
python-service  |   File "/usr/local/lib/python3.10/site-packages/anyio/to_thread.py", line 61, in run_sync
python-service  |     return await get_async_backend().run_sync_in_worker_thread(
python-service  |   File "/usr/local/lib/python3.10/site-packages/anyio/_backends/_asyncio.py", line 2525, in run_sync_in_worker_thread
python-service  |     return await future
python-service  |   File "/usr/local/lib/python3.10/site-packages/anyio/_backends/_asyncio.py", line 986, in run
python-service  |     result = context.run(func, *args)
python-service  |   File "/app/app/main.py", line 663, in _clone_voice
python-service  |     raise HTTPException(status_code=500, detail=f"Voice cloning failed: {str(err)}") from err
python-service  | fastapi.exceptions.HTTPException: 500: Voice cloning failed:  ❗ Language fa is not supported. Supported languages are ['en', 'es', 'fr', 'de', 'it', 'pt', 'pl', 'tr', 'ru', 'nl', 'cs', 'ar', 'zh-cn', 'hu', 'ko', 'ja', 'hi']
python-service  | INFO:     172.18.0.4:37982 - "POST /voice-clone HTTP/1.1" 500 Internal Server Error
python-service  | INFO:     127.0.0.1:43022 - "GET /health HTTP/1.1" 200 OK
go-service      | [0.500ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:11 /app/main.go:656 record not found
go-service      | [0.678ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:13 /app/main.go:656 record not found
go-service      | [0.388ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:16 /app/main.go:656 record not found
go-service      | [0.501ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:18 /app/main.go:656 record not found
go-service      | [0.461ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:20 /app/main.go:656 record not found
go-service      | [0.513ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:22 /app/main.go:656 record not found
go-service      | [0.682ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:24 /app/main.go:656 record not found
go-service      | [0.704ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:26 /app/main.go:656 record not found
go-service      | [0.422ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:28 /app/main.go:656 record not found
python-service  | INFO:     127.0.0.1:35816 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60556 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:53846 - "GET /health HTTP/1.1" 200 OK
go-service      | [0.670ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:30 /app/main.go:656 record not found
go-service      | [0.589ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:32 /app/main.go:656 record not found
go-service      | [0.562ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:34 /app/main.go:656 record not found
go-service      | [0.294ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:36 /app/main.go:656 record not found
go-service      | [0.571ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:38 /app/main.go:656 record not found
go-service      | [0.507ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:40 /app/main.go:656 record not found
go-service      | [0.587ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:42 /app/main.go:656 record not found
go-service      | [0.637ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:44 /app/main.go:656 record not found
go-service      | [0.431ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:46 /app/main.go:656 record not found
go-service      | [0.427ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:48 /app/main.go:656 record not found
go-service      | [0.522ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:50 /app/main.go:656 record not found
go-service      | [0.563ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:52 /app/main.go:656 record not found
go-service      | [0.484ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:54 /app/main.go:656 record not found
go-service      | [0.421ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:56 /app/main.go:656 record not found
go-service      | [0.320ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:50:58 /app/main.go:656 record not found
go-service      | [0.526ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:00 /app/main.go:656 record not found
go-service      | [1.379ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:02 /app/main.go:656 record not found
go-service      | [0.728ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:04 /app/main.go:656 record not found
go-service      | [0.433ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:06 /app/main.go:656 record not found
go-service      | [0.570ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:08 /app/main.go:656 record not found
go-service      | [0.553ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:10 /app/main.go:656 record not found
go-service      | [0.528ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:12 /app/main.go:656 record not found
go-service      | [0.418ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:14 /app/main.go:656 record not found
go-service      | [0.511ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:16 /app/main.go:656 record not found
go-service      | [0.378ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:18 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:20 /app/main.go:656 record not found
go-service      | [0.510ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:22 /app/main.go:656 record not found
go-service      | [0.535ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:24 /app/main.go:656 record not found
go-service      | [0.345ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:26 /app/main.go:656 record not found
go-service      | [0.569ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:28 /app/main.go:656 record not found
go-service      | [0.650ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:30 /app/main.go:656 record not found
go-service      | [0.342ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:32 /app/main.go:656 record not found
go-service      | [0.427ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:34 /app/main.go:656 record not found
go-service      | [0.319ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:36 /app/main.go:656 record not found
go-service      | [0.585ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:38 /app/main.go:656 record not found
go-service      | [0.557ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:40 /app/main.go:656 record not found
go-service      | [0.540ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:42 /app/main.go:656 record not found
go-service      | [0.525ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:44 /app/main.go:656 record not found
go-service      | [0.565ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:46 /app/main.go:656 record not found
go-service      | [0.424ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:48 /app/main.go:656 record not found
go-service      | [0.411ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:50 /app/main.go:656 record not found
go-service      | [0.481ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:52 /app/main.go:656 record not found
go-service      | [0.459ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:54 /app/main.go:656 record not found
go-service      | [0.643ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:56 /app/main.go:656 record not found
go-service      | [0.648ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:51:58 /app/main.go:656 record not found
go-service      | [0.292ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:00 /app/main.go:656 record not found
go-service      | [0.499ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:02 /app/main.go:656 record not found
go-service      | [0.575ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:04 /app/main.go:656 record not found
go-service      | [0.450ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:06 /app/main.go:656 record not found
go-service      | [0.571ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:08 /app/main.go:656 record not found
go-service      | [0.356ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:10 /app/main.go:656 record not found
go-service      | [0.593ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:12 /app/main.go:656 record not found
go-service      | [0.411ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:14 /app/main.go:656 record not found
go-service      | [0.460ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:16 /app/main.go:656 record not found
go-service      | [0.440ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:18 /app/main.go:656 record not found
go-service      | [0.466ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:20 /app/main.go:656 record not found
go-service      | [0.501ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:22 /app/main.go:656 record not found
go-service      | [0.492ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:24 /app/main.go:656 record not found
go-service      | [0.513ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:26 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:28 /app/main.go:656 record not found
go-service      | [0.711ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:30 /app/main.go:656 record not found
go-service      | [0.675ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:32 /app/main.go:656 record not found
go-service      | [0.532ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:34 /app/main.go:656 record not found
go-service      | [0.660ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:36 /app/main.go:656 record not found
go-service      | [0.599ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:38 /app/main.go:656 record not found
go-service      | [0.341ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:40 /app/main.go:656 record not found
go-service      | [0.490ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:42 /app/main.go:656 record not found
go-service      | [0.641ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:44 /app/main.go:656 record not found
go-service      | [0.493ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:46 /app/main.go:656 record not found
go-service      | [0.493ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:48 /app/main.go:656 record not found
go-service      | [0.401ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:50 /app/main.go:656 record not found
go-service      | [0.567ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:52 /app/main.go:656 record not found
go-service      | [0.519ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:54 /app/main.go:656 record not found
go-service      | [0.473ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:56 /app/main.go:656 record not found
go-service      | [0.671ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:52:58 /app/main.go:656 record not found
go-service      | [0.557ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:00 /app/main.go:656 record not found
go-service      | [0.651ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:02 /app/main.go:656 record not found
go-service      | [0.546ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:04 /app/main.go:656 record not found
go-service      | [0.461ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:06 /app/main.go:656 record not found
go-service      | [0.288ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:08 /app/main.go:656 record not found
go-service      | [0.737ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:10 /app/main.go:656 record not found
go-service      | [0.465ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:12 /app/main.go:656 record not found
go-service      | [0.396ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:14 /app/main.go:656 record not found
go-service      | [0.553ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:16 /app/main.go:656 record not found
go-service      | [0.523ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:18 /app/main.go:656 record not found
go-service      | [0.540ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:20 /app/main.go:656 record not found
go-service      | [0.301ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:22 /app/main.go:656 record not found
go-service      | [0.541ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:24 /app/main.go:656 record not found
go-service      | [0.512ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:26 /app/main.go:656 record not found
go-service      | [0.545ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:28 /app/main.go:656 record not found
go-service      | [0.438ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:30 /app/main.go:656 record not found
go-service      | [0.483ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:32 /app/main.go:656 record not found
go-service      | [0.474ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:34 /app/main.go:656 record not found
go-service      | [0.537ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:36 /app/main.go:656 record not found
go-service      | [0.469ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:38 /app/main.go:656 record not found
go-service      | [0.533ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:40 /app/main.go:656 record not found
go-service      | [0.496ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:42 /app/main.go:656 record not found
go-service      | [0.583ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:44 /app/main.go:656 record not found
go-service      | [0.558ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:46 /app/main.go:656 record not found
go-service      | [0.440ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:48 /app/main.go:656 record not found
go-service      | [0.340ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:50 /app/main.go:656 record not found
go-service      | [0.510ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:52 /app/main.go:656 record not found
go-service      | [0.625ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:54 /app/main.go:656 record not found
go-service      | [0.426ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:56 /app/main.go:656 record not found
go-service      | [0.293ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:53:58 /app/main.go:656 record not found
go-service      | [0.404ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:00 /app/main.go:656 record not found
go-service      | [0.453ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:02 /app/main.go:656 record not found
go-service      | [0.535ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:04 /app/main.go:656 record not found
go-service      | [0.458ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:06 /app/main.go:656 record not found
go-service      | [0.539ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:08 /app/main.go:656 record not found
go-service      | [0.616ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:10 /app/main.go:656 record not found
go-service      | [0.625ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:12 /app/main.go:656 record not found
go-service      | [0.562ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:14 /app/main.go:656 record not found
go-service      | [0.507ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:16 /app/main.go:656 record not found
go-service      | [0.569ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:18 /app/main.go:656 record not found
go-service      | [0.676ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:20 /app/main.go:656 record not found
go-service      | [0.549ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:22 /app/main.go:656 record not found
go-service      | [0.466ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:24 /app/main.go:656 record not found
go-service      | [0.442ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:26 /app/main.go:656 record not found
go-service      | [0.529ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:28 /app/main.go:656 record not found
go-service      | [0.534ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:30 /app/main.go:656 record not found
go-service      | [0.549ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:32 /app/main.go:656 record not found
go-service      | [0.423ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:34 /app/main.go:656 record not found
go-service      | [0.466ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:36 /app/main.go:656 record not found
go-service      | [0.544ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:38 /app/main.go:656 record not found
go-service      | [0.532ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:40 /app/main.go:656 record not found
go-service      | [0.455ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:42 /app/main.go:656 record not found
go-service      | [0.620ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:44 /app/main.go:656 record not found
go-service      | [0.524ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:46 /app/main.go:656 record not found
go-service      | [0.482ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:48 /app/main.go:656 record not found
go-service      | [0.597ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:50 /app/main.go:656 record not found
go-service      | [0.665ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      | 2025/12/06 16:54:53 voice clone worker: job 10 failed: python service returned 500: {"detail":"Voice cloning failed"}
go-service      |
go-service      | 2025/12/06 16:54:53 /app/main.go:656 record not found
go-service      | [0.433ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:55 /app/main.go:656 record not found
go-service      | [0.570ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:57 /app/main.go:656 record not found
go-service      | [0.395ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:54:59 /app/main.go:656 record not found
go-service      | [0.484ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:01 /app/main.go:656 record not found
go-service      | [0.546ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:03 /app/main.go:656 record not found
go-service      | [0.314ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:05 /app/main.go:656 record not found
go-service      | [0.496ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:07 /app/main.go:656 record not found
go-service      | [0.545ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:09 /app/main.go:656 record not found
go-service      | [0.448ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:11 /app/main.go:656 record not found
go-service      | [0.536ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:13 /app/main.go:656 record not found
go-service      | [0.436ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:15 /app/main.go:656 record not found
go-service      | [0.581ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:17 /app/main.go:656 record not found
go-service      | [0.345ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:19 /app/main.go:656 record not found
go-service      | [0.422ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:21 /app/main.go:656 record not found
go-service      | [0.522ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:23 /app/main.go:656 record not found
go-service      | [0.505ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:25 /app/main.go:656 record not found
go-service      | [0.450ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 16:55:27 /app/main.go:656 record not found
go-service      | [0.430ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
root@Automations:/opt/TTS#


