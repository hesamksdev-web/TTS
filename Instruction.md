root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs -f
tts-db  |
tts-db  | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db  |
tts-db  | 2025-12-06 13:20:43.714 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 15.2.0) 15.2.0, 64-bit
tts-db  | 2025-12-06 13:20:43.714 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db  | 2025-12-06 13:20:43.714 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db  | 2025-12-06 13:20:43.717 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db  | 2025-12-06 13:20:43.722 UTC [28] LOG:  database system was shut down at 2025-12-06 13:20:20 UTC
tts-db  | 2025-12-06 13:20:43.728 UTC [1] LOG:  database system is ready to accept connections
tts-db  | 2025-12-06 13:25:43.779 UTC [26] LOG:  checkpoint starting: time
tts-db  | 2025-12-06 13:25:44.598 UTC [26] LOG:  checkpoint complete: wrote 11 buffers (0.1%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.806 s, sync=0.006 s, total=0.819 s; sync files=10, longest=0.003 s, average=0.001 s; distance=7 kB, estimate=7 kB
caddy-proxy  | {"level":"info","ts":1765027321.4437437,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy     | {"level":"info","ts":1765027321.444239,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1765027321.4443083,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1765027321.4463603,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy     | {"level":"warn","ts":1765027321.4464238,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1765027321.4471397,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy     | {"level":"info","ts":1765027321.4474816,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy     | {"level":"info","ts":1765027321.4476237,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc000333e80"}
caddy-proxy     | {"level":"warn","ts":1765027321.4484377,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"warn","ts":1765027321.4484577,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"info","ts":1765027321.4484625,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy     | {"level":"info","ts":1765027321.4488297,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1765027321.4488618,"msg":"serving initial configuration"}
caddy-proxy     | {"level":"info","ts":1765027321.4511123,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765113721.4511085,"try_again_in":86399.999999058}
go-service      | 2025/12/06 13:21:55 Go Service (SaaS Edition) listening on :8080
python-service  | /usr/local/lib/python3.10/site-packages/jieba/_compat.py:18: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   import pkg_resources
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  |  > Downloading model to /root/.local/share/tts/tts/tts_models--multilingual--multi-dataset--xtts_v2
caddy-proxy     | {"level":"info","ts":1765027321.4512866,"logger":"tls","msg":"finished cleaning storage units"}
100%|██████████| 1.87G/1.87G [00:22<00:00, 84.3MiB/s]
4.37kiB [00:00, 21.5kiB/s]
361kiB [00:00, 1.82MiB/s]
100%|██████████| 32.0/32.0 [00:00<00:00, 143iB/s]
python-service  |  > Model's license - CPML
python-service  |  > Check https://coqui.ai/cpml.txt for more info.
python-service  |  > Using model: xtts
100%|██████████| 7.75M/7.75M [00:17<00:00, 55.6MiB/s]INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  | INFO:     127.0.0.1:34954 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37486 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54326 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45156 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34100 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59590 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Text splitted to sentences.
python-service  | ['The town library sits where the coastline curls like a question mark, and in it the air always smells faintly of rain and old paper.']
python-service  | INFO:     127.0.0.1:55776 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60008 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38714 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48446 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 38.43591284751892
python-service  |  > Real-time factor: 4.071757428932817
python-service  |  > Text splitted to sentences.
python-service  | ['Nia kept the place steady, like a lighthouse keeper who never expects the sea to listen but keeps the lamp lit anyway.']
python-service  | INFO:     127.0.0.1:35696 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:53926 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56048 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:49804 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 43.57833933830261
python-service  |  > Real-time factor: 3.905599200142959
python-service  |  > Text splitted to sentences.
python-service  | ['She shelved books, brewed bad coffee that somehow tasted better than it deserved, and watched the town breathe in and out through stories.']
python-service  | INFO:     127.0.0.1:33986 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56040 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47270 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52568 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52236 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 49.07546544075012
python-service  |  > Real-time factor: 3.9502438999202014
python-service  | INFO:     172.18.0.4:57462 - "POST /voice-clone HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46302 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47634 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60416 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60136 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52348 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:39022 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45814 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50280 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50550 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36648 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36888 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35802 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59058 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34828 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47308 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34636 - "GET /health HTTP/1.1" 200 OK
tts-frontend  |
tts-frontend    | > frontend@0.1.0 start
tts-frontend    | > next start
tts-frontend    |
tts-frontend    |    ▲ Next.js 16.0.6
tts-frontend    |    - Local:         http://localhost:3000
tts-frontend    |    - Network:       http://172.18.0.5:3000
tts-frontend    |
tts-frontend    |  ✓ Starting...
tts-frontend    |  ✓ Ready in 428ms
go-service      | 2025/12/06 13:21:55 Voice clone worker started
go-service      |
go-service      | 2025/12/06 13:21:55 /app/main.go:656 record not found
go-service      | [1.141ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:21:57 /app/main.go:656 record not found
go-service      | [0.751ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:21:59 /app/main.go:656 record not found
go-service      | [0.960ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:01 /app/main.go:656 record not found
go-service      | [0.712ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:03 /app/main.go:656 record not found
go-service      | [0.700ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:05 /app/main.go:656 record not found
go-service      | [0.799ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:07 /app/main.go:656 record not found
go-service      | [0.615ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:09 /app/main.go:656 record not found
go-service      | [0.526ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:11 /app/main.go:656 record not found
go-service      | [0.382ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:13 /app/main.go:656 record not found
go-service      | [0.688ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:15 /app/main.go:656 record not found
go-service      | [0.346ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:17 /app/main.go:656 record not found
go-service      | [0.449ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:19 /app/main.go:656 record not found
go-service      | [0.796ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:21 /app/main.go:656 record not found
go-service      | [0.489ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:23 /app/main.go:656 record not found
go-service      | [0.582ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:25 /app/main.go:656 record not found
go-service      | [0.582ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:27 /app/main.go:656 record not found
go-service      | [0.391ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:29 /app/main.go:656 record not found
go-service      | [0.414ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:31 /app/main.go:656 record not found
go-service      | [0.478ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:33 /app/main.go:656 record not found
go-service      | [0.343ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:35 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:37 /app/main.go:656 record not found
go-service      | [0.460ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:39 /app/main.go:656 record not found
go-service      | [0.409ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:41 /app/main.go:656 record not found
go-service      | [0.486ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:43 /app/main.go:656 record not found
go-service      | [0.492ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:22:45 /app/main.go:656 record not found
go-service      | [0.463ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:24:59 /app/main.go:656 record not found
go-service      | [0.402ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:01 /app/main.go:656 record not found
go-service      | [0.580ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:03 /app/main.go:656 record not found
go-service      | [0.569ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:05 /app/main.go:656 record not found
go-service      | [0.483ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:07 /app/main.go:656 record not found
go-service      | [0.424ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:09 /app/main.go:656 record not found
go-service      | [0.513ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:11 /app/main.go:656 record not found
go-service      | [0.568ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:13 /app/main.go:656 record not found
go-service      | [0.832ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:15 /app/main.go:656 record not found
go-service      | [0.624ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:17 /app/main.go:656 record not found
go-service      | [0.591ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:19 /app/main.go:656 record not found
go-service      | [0.535ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:21 /app/main.go:656 record not found
go-service      | [0.466ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:23 /app/main.go:656 record not found
go-service      | [0.329ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:25 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:27 /app/main.go:656 record not found
go-service      | [0.489ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:29 /app/main.go:656 record not found
go-service      | [0.405ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:31 /app/main.go:656 record not found
go-service      | [0.364ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:33 /app/main.go:656 record not found
go-service      | [0.583ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:35 /app/main.go:656 record not found
go-service      | [0.600ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:37 /app/main.go:656 record not found
go-service      | [0.480ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:39 /app/main.go:656 record not found
go-service      | [0.462ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:41 /app/main.go:656 record not found
go-service      | [0.390ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:43 /app/main.go:656 record not found
go-service      | [0.494ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:45 /app/main.go:656 record not found
go-service      | [0.472ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:47 /app/main.go:656 record not found
go-service      | [0.389ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:49 /app/main.go:656 record not found
go-service      | [0.523ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:51 /app/main.go:656 record not found
go-service      | [0.353ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:53 /app/main.go:656 record not found
go-service      | [0.293ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:55 /app/main.go:656 record not found
go-service      | [0.517ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:57 /app/main.go:656 record not found
go-service      | [0.618ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:25:59 /app/main.go:656 record not found
go-service      | [0.738ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:01 /app/main.go:656 record not found
go-service      | [0.521ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:03 /app/main.go:656 record not found
go-service      | [0.560ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:05 /app/main.go:656 record not found
go-service      | [0.521ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:07 /app/main.go:656 record not found
go-service      | [0.502ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:09 /app/main.go:656 record not found
go-service      | [0.557ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:11 /app/main.go:656 record not found
go-service      | [0.534ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:13 /app/main.go:656 record not found
go-service      | [0.479ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:15 /app/main.go:656 record not found
go-service      | [0.475ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:17 /app/main.go:656 record not found
go-service      | [0.526ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:19 /app/main.go:656 record not found
go-service      | [0.600ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:21 /app/main.go:656 record not found
go-service      | [0.448ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:23 /app/main.go:656 record not found
go-service      | [0.471ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:25 /app/main.go:656 record not found
go-service      | [0.429ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:27 /app/main.go:656 record not found
go-service      | [0.345ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:29 /app/main.go:656 record not found
go-service      | [0.460ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:31 /app/main.go:656 record not found
go-service      | [0.506ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:33 /app/main.go:656 record not found
go-service      | [0.642ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:35 /app/main.go:656 record not found
go-service      | [0.537ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:37 /app/main.go:656 record not found
go-service      | [0.614ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:39 /app/main.go:656 record not found
go-service      | [0.550ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:41 /app/main.go:656 record not found
go-service      | [0.346ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:43 /app/main.go:656 record not found
go-service      | [0.572ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:45 /app/main.go:656 record not found
go-service      | [0.480ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:47 /app/main.go:656 record not found
go-service      | [0.609ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:49 /app/main.go:656 record not found
go-service      | [0.584ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:51 /app/main.go:656 record not found
go-service      | [0.341ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:53 /app/main.go:656 record not found
go-service      | [0.577ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:55 /app/main.go:656 record not found
go-service      | [0.580ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:57 /app/main.go:656 record not found
go-service      | [0.492ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:26:59 /app/main.go:656 record not found
go-service      | [0.548ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:01 /app/main.go:656 record not found
go-service      | [0.451ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:03 /app/main.go:656 record not found
go-service      | [0.710ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:05 /app/main.go:656 record not found
go-service      | [0.571ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:07 /app/main.go:656 record not found
go-service      | [0.537ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:09 /app/main.go:656 record not found
go-service      | [0.445ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:11 /app/main.go:656 record not found
go-service      | [0.541ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:13 /app/main.go:656 record not found
go-service      | [0.544ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:15 /app/main.go:656 record not found
go-service      | [0.502ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:17 /app/main.go:656 record not found
go-service      | [0.871ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:19 /app/main.go:656 record not found
go-service      | [0.531ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:21 /app/main.go:656 record not found
go-service      | [0.455ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:23 /app/main.go:656 record not found
go-service      | [0.696ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:25 /app/main.go:656 record not found
go-service      | [0.350ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:27 /app/main.go:656 record not found
go-service      | [0.461ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:29 /app/main.go:656 record not found
go-service      | [0.303ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:31 /app/main.go:656 record not found
go-service      | [0.451ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:33 /app/main.go:656 record not found
go-service      | [0.628ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:35 /app/main.go:656 record not found
go-service      | [0.554ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:37 /app/main.go:656 record not found
go-service      | [0.419ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:39 /app/main.go:656 record not found
go-service      | [0.545ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:41 /app/main.go:656 record not found
go-service      | [0.638ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:43 /app/main.go:656 record not found
go-service      | [0.414ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:45 /app/main.go:656 record not found
go-service      | [0.391ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:47 /app/main.go:656 record not found
go-service      | [0.464ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:59402 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:27:49 /app/main.go:656 record not found
go-service      | [0.471ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:51 /app/main.go:656 record not found
go-service      | [0.535ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:53 /app/main.go:656 record not found
go-service      | [0.434ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:55 /app/main.go:656 record not found
go-service      | [0.350ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:27:57 /app/main.go:656 record not found
go-service      | [0.364ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:32992 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:27:59 /app/main.go:656 record not found
go-service      | [0.448ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:01 /app/main.go:656 record not found
go-service      | [0.398ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:03 /app/main.go:656 record not found
go-service      | [0.517ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:05 /app/main.go:656 record not found
go-service      | [0.534ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:07 /app/main.go:656 record not found
go-service      | [0.400ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:44246 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:28:09 /app/main.go:656 record not found
go-service      | [0.485ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:11 /app/main.go:656 record not found
go-service      | [0.465ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:13 /app/main.go:656 record not found
go-service      | [0.509ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:15 /app/main.go:656 record not found
go-service      | [0.490ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:17 /app/main.go:656 record not found
go-service      | [0.367ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:45120 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:28:19 /app/main.go:656 record not found
go-service      | [0.652ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:21 /app/main.go:656 record not found
go-service      | [0.486ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:23 /app/main.go:656 record not found
go-service      | [0.471ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:25 /app/main.go:656 record not found
go-service      | [0.525ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:27 /app/main.go:656 record not found
go-service      | [0.284ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:47850 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:28:29 /app/main.go:656 record not found
go-service      | [0.425ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:31 /app/main.go:656 record not found
go-service      | [0.523ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:28:33 /app/main.go:656 record not found
go-service      | [0.692ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
^C
root@Automations:/opt/TTS#