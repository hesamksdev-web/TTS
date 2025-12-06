root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs -f
tts-frontend  |
python-service  | /usr/local/lib/python3.10/site-packages/jieba/_compat.py:18: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   import pkg_resources
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  |  > Downloading model to /root/.local/share/tts/tts/tts_models--multilingual--multi-dataset--xtts_v2
tts-db          |
tts-frontend    | > frontend@0.1.0 start
tts-frontend    | > next start
tts-frontend    |
tts-frontend    |    ▲ Next.js 16.0.6
tts-frontend    |    - Local:         http://localhost:3000
tts-frontend    |    - Network:       http://172.18.0.5:3000
tts-frontend    |
tts-frontend    |  ✓ Starting...
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db          |
tts-db          | 2025-12-06 15:24:39.032 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 15.2.0) 15.2.0, 64-bit
tts-db          | 2025-12-06 15:24:39.032 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-06 15:24:39.032 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-06 15:24:39.036 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-06 15:24:39.043 UTC [29] LOG:  database system was shut down at 2025-12-06 15:24:00 UTC
tts-db          | 2025-12-06 15:24:39.052 UTC [1] LOG:  database system is ready to accept connections
go-service      | 2025/12/06 15:25:56 Voice clone worker started
go-service      | 2025/12/06 15:25:56 Go Service (SaaS Edition) listening on :8080
go-service      |
go-service      | 2025/12/06 15:25:56 /app/main.go:656 record not found
go-service      | [3.113ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:25:58 /app/main.go:656 record not found
go-service      | [0.598ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:00 /app/main.go:656 record not found
go-service      | [0.751ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:02 /app/main.go:656 record not found
go-service      | [0.928ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:04 /app/main.go:656 record not found
go-service      | [0.785ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:06 /app/main.go:656 record not found
go-service      | [0.697ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:08 /app/main.go:656 record not found
go-service      | [0.554ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:10 /app/main.go:656 record not found
go-service      | [0.624ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:12 /app/main.go:656 record not found
go-service      | [0.508ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:14 /app/main.go:656 record not found
go-service      | [0.443ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:16 /app/main.go:656 record not found
go-service      | [0.708ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:18 /app/main.go:656 record not found
go-service      | [0.505ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:20 /app/main.go:656 record not found
go-service      | [0.463ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:22 /app/main.go:656 record not found
go-service      | [0.542ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:24 /app/main.go:656 record not found
go-service      | [0.571ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:26 /app/main.go:656 record not found
go-service      | [0.444ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:28 /app/main.go:656 record not found
go-service      | [0.539ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
tts-frontend    |  ✓ Ready in 438ms
caddy-proxy     | {"level":"info","ts":1765034761.8212342,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
go-service      | 2025/12/06 15:26:30 /app/main.go:656 record not found
go-service      | [0.550ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:32 /app/main.go:656 record not found
go-service      | [0.405ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:34 /app/main.go:656 record not found
go-service      | [0.649ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:36 /app/main.go:656 record not found
go-service      | [0.622ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:38 /app/main.go:656 record not found
go-service      | [0.409ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:40 /app/main.go:656 record not found
go-service      | [0.612ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:42 /app/main.go:656 record not found
go-service      | [0.393ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:44 /app/main.go:656 record not found
go-service      | [0.588ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:46 /app/main.go:656 record not found
go-service      | [0.472ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:48 /app/main.go:656 record not found
go-service      | [0.598ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:50 /app/main.go:656 record not found
go-service      | [0.490ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
100%|██████████| 1.87G/1.87G [00:23<00:00, 81.2MiB/s]
4.37kiB [00:00, 20.9kiB/s]
361kiB [00:00, 1.19MiB/s]
caddy-proxy     | {"level":"info","ts":1765034761.821467,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1765034761.8222067,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1765034761.8236537,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy     | {"level":"warn","ts":1765034761.8236706,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1765034761.8240492,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy     | {"level":"info","ts":1765034761.8242042,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy     | {"level":"info","ts":1765034761.8243587,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc0005c8c00"}
caddy-proxy     | {"level":"warn","ts":1765034761.8252368,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"warn","ts":1765034761.825261,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"info","ts":1765034761.8252823,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy     | {"level":"info","ts":1765034761.825578,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1765034761.8255913,"msg":"serving initial configuration"}
caddy-proxy     | {"level":"info","ts":1765034761.830182,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765121161.8301806,"try_again_in":86399.999999529}
caddy-proxy     | {"level":"info","ts":1765034761.8302634,"logger":"tls","msg":"finished cleaning storage units"}
go-service      | 2025/12/06 15:26:52 /app/main.go:656 record not found
go-service      | [0.451ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:54 /app/main.go:656 record not found
go-service      | [0.480ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:56 /app/main.go:656 record not found
go-service      | [0.625ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:26:58 /app/main.go:656 record not found
go-service      | [0.415ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:00 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:02 /app/main.go:656 record not found
go-service      | [0.530ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:04 /app/main.go:656 record not found
go-service      | [0.450ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:06 /app/main.go:656 record not found
go-service      | [0.588ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:08 /app/main.go:656 record not found
go-service      | [0.322ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:10 /app/main.go:656 record not found
go-service      | [0.487ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:12 /app/main.go:656 record not found
go-service      | [0.468ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:14 /app/main.go:656 record not found
go-service      | [0.441ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:16 /app/main.go:656 record not found
go-service      | [0.660ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:18 /app/main.go:656 record not found
go-service      | [0.558ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:27:20 /app/main.go:656 record not found
go-service      | [0.617ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
100%|██████████| 32.0/32.0 [00:00<00:00, 128iB/s]
python-service  |  > Model's license - CPML
python-service  |  > Check https://coqui.ai/cpml.txt for more info.
python-service  |  > Using model: xtts
100%|██████████| 7.75M/7.75M [00:16<00:00, 49.7MiB/s]INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  | INFO:     127.0.0.1:38188 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:39144 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59494 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:51136 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54530 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59292 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46420 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:44200 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45016 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Text splitted to sentences.
python-service  | ['And so the town continues to breathe through stories, one letter at a time, a little louder, a little brighter, the sea outside listening like a patient confidant.']
python-service  | INFO:     127.0.0.1:53314 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35288 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47056 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48910 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47550 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59094 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:33640 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 66.27225685119629
python-service  |  > Real-time factor: 4.085870083346973
python-service  |  > Text splitted to sentences.
python-service  | ['The library’s lights burn a little longer, not out of obligation, but because a room where people are seen by someone who cares is the kind of light that refuses to dim The attic hums with quiet reassurance and in that hum.']
python-service  | INFO:     127.0.0.1:33534 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:57602 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:41048 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:44862 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:41700 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:53494 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60676 - "GET /health HTTP/1.1" 200 OK
tts-db          | 2025-12-06 15:29:39.048 UTC [27] LOG:  checkpoint starting: time
tts-db          | 2025-12-06 15:29:39.465 UTC [27] LOG:  checkpoint complete: wrote 7 buffers (0.0%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.407 s, sync=0.005 s, total=0.417 s; sync files=6, longest=0.003 s, average=0.001 s; distance=7 kB, estimate=7 kB
python-service  | INFO:     127.0.0.1:45214 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 81.31580924987793
python-service  |  > Real-time factor: 4.105335737351651
python-service  |  > Text splitted to sentences.
python-service  | ['Nia finally feels what her grandmother must have known all along: listening poured out as letters and shared aloud is a kind of kindness that travels farther than any sea-wind can carry.']
python-service  | INFO:     127.0.0.1:58712 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60676 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45128 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:33182 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:49944 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60340 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 60.455384492874146
python-service  |  > Real-time factor: 4.139056920574404
python-service  | INFO:     172.18.0.4:46908 - "POST /voice-clone HTTP/1.1" 200 OK
go-service      | 2025/12/06 15:30:51 Voice clone job 8: successfully wrote 2223372 bytes to /app/data/voice_clone/job_1_1765034840275837249/output.wav
go-service      |
go-service      | 2025/12/06 15:30:51 /app/main.go:656 record not found
go-service      | [0.347ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      | 2025/12/06 15:30:51 Serving voice clone file: /app/data/voice_clone/job_1_1765034840275837249/output.wav (job_id: 8, size: 2223372 bytes)
go-service      |
go-service      | 2025/12/06 15:30:53 /app/main.go:656 record not found
go-service      | [0.462ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:30:55 /app/main.go:656 record not found
go-service      | [0.420ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:30:57 /app/main.go:656 record not found
go-service      | [0.488ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:46128 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 15:30:59 /app/main.go:656 record not found
go-service      | [0.770ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:01 /app/main.go:656 record not found
go-service      | [0.777ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:03 /app/main.go:656 record not found
go-service      | [0.485ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:05 /app/main.go:656 record not found
go-service      | [0.569ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:07 /app/main.go:656 record not found
go-service      | [0.508ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:48550 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 15:31:09 /app/main.go:656 record not found
go-service      | [0.546ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:11 /app/main.go:656 record not found
go-service      | [0.532ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:13 /app/main.go:656 record not found
go-service      | [0.460ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:15 /app/main.go:656 record not found
go-service      | [0.470ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:17 /app/main.go:656 record not found
go-service      | [0.622ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:41714 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 15:31:19 /app/main.go:656 record not found
go-service      | [0.443ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:21 /app/main.go:656 record not found
go-service      | [0.626ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:23 /app/main.go:656 record not found
go-service      | [0.605ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:25 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:27 /app/main.go:656 record not found
go-service      | [0.591ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:33008 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 15:31:29 /app/main.go:656 record not found
go-service      | [0.482ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:31 /app/main.go:656 record not found
go-service      | [0.456ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:33 /app/main.go:656 record not found
go-service      | [0.347ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:35 /app/main.go:656 record not found
go-service      | [0.567ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:37 /app/main.go:656 record not found
go-service      | [0.406ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:58628 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 15:31:39 /app/main.go:656 record not found
go-service      | [0.594ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:41 /app/main.go:656 record not found
go-service      | [0.571ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:43 /app/main.go:656 record not found
go-service      | [0.489ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:45 /app/main.go:656 record not found
go-service      | [0.640ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:47 /app/main.go:656 record not found
go-service      | [0.548ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:40784 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 15:31:49 /app/main.go:656 record not found
go-service      | [0.519ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:51 /app/main.go:656 record not found
go-service      | [0.625ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:53 /app/main.go:656 record not found
go-service      | [0.503ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:55 /app/main.go:656 record not found
go-service      | [0.448ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:31:57 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:52552 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 15:31:59 /app/main.go:656 record not found
go-service      | [0.503ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:32:01 /app/main.go:656 record not found
go-service      | [0.530ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:32:03 /app/main.go:656 record not found
go-service      | [0.581ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 15:32:05 /app/main.go:656 record not found
go-service      | [0.514ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
^C