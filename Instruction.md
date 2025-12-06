root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs -f
go-service  | 2025/12/06 12:43:51 Go Service (SaaS Edition) listening on :8080
go-service  | 2025/12/06 12:43:51 Voice clone worker started
go-service  |
go-service  | 2025/12/06 12:43:51 /app/main.go:656 record not found
go-service  | [1.099ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:43:53 /app/main.go:656 record not found
go-service  | [0.912ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:43:55 /app/main.go:656 record not found
go-service  | [0.813ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:43:57 /app/main.go:656 record not found
go-service  | [0.897ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:43:59 /app/main.go:656 record not found
go-service  | [0.778ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:01 /app/main.go:656 record not found
go-service  | [0.837ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:03 /app/main.go:656 record not found
go-service  | [0.340ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:05 /app/main.go:656 record not found
go-service  | [0.451ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:07 /app/main.go:656 record not found
go-service  | [0.696ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:09 /app/main.go:656 record not found
go-service  | [0.478ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:11 /app/main.go:656 record not found
go-service  | [0.628ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:13 /app/main.go:656 record not found
go-service  | [0.719ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:15 /app/main.go:656 record not found
go-service  | [0.478ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:17 /app/main.go:656 record not found
go-service  | [0.492ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:19 /app/main.go:656 record not found
go-service  | [0.458ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:21 /app/main.go:656 record not found
go-service  | [0.471ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:23 /app/main.go:656 record not found
go-service  | [0.464ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:25 /app/main.go:656 record not found
go-service  | [0.421ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:27 /app/main.go:656 record not found
go-service  | [0.638ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 12:44:29 /app/main.go:656 record not found
go-service  | [0.635ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
tts-frontend  |
tts-frontend  | > frontend@0.1.0 start
tts-frontend  | > next start
tts-frontend  |
tts-frontend  |    ▲ Next.js 16.0.6
tts-frontend  |    - Local:         http://localhost:3000
tts-frontend  |    - Network:       http://172.18.0.5:3000
tts-frontend  |
tts-frontend  |  ✓ Starting...
tts-frontend  |  ✓ Ready in 503ms
python-service  | /usr/local/lib/python3.10/site-packages/jieba/_compat.py:18: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   import pkg_resources
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  | ERROR:tts_service:Failed to pre-load engines: 'cv'
python-service  | INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  | INFO:     127.0.0.1:36704 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50532 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37296 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59446 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:55074 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Downloading model to /root/.local/share/tts/tts/tts_models--multilingual--multi-dataset--xtts_v2
100%|█████████▉| 1.87G/1.87G [00:21<00:00, 90.9MiB/s]
caddy-proxy     | {"level":"info","ts":1765025037.095264,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy     | {"level":"info","ts":1765025037.0954766,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1765025037.0960805,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1765025037.0970464,"msg":"adapted config to JSON","adapter":"caddyfile"}
tts-db          |
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
100%|██████████| 1.87G/1.87G [00:21<00:00, 85.1MiB/s]
4.37kiB [00:00, 16.8kiB/s]
361kiB [00:00, 1.71MiB/s]
100%|██████████| 32.0/32.0 [00:00<00:00, 111iB/s]
python-service  |  > Model's license - CPML
python-service  |  > Check https://coqui.ai/cpml.txt for more info.
python-service  |  > Using model: xtts
python-service  |  > Text splitted to sentences.
python-service  | INFO:     127.0.0.1:51062 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:32878 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48276 - "GET /health HTTP/1.1" 200 OK
python-service  | ['The town library sits where the coastline curls like a question mark, and in it the air always smells faintly of rain and old paper.']
python-service  | INFO:     127.0.0.1:35574 - "GET /health HTTP/1.1" 200 OK
caddy-proxy     | {"level":"warn","ts":1765025037.0970578,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1765025037.0973358,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy     | {"level":"info","ts":1765025037.0974524,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy     | {"level":"info","ts":1765025037.0975559,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc000702480"}
caddy-proxy     | {"level":"warn","ts":1765025037.0979633,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"warn","ts":1765025037.0979743,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"info","ts":1765025037.0979774,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy     | {"level":"info","ts":1765025037.0981393,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1765025037.0981479,"msg":"serving initial configuration"}
caddy-proxy     | {"level":"info","ts":1765025037.1007378,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765111437.100736,"try_again_in":86399.999999519}
caddy-proxy     | {"level":"info","ts":1765025037.1009345,"logger":"tls","msg":"finished cleaning storage units"}
tts-db          | 2025-12-06 12:43:20.310 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 15.2.0) 15.2.0, 64-bit
tts-db          | 2025-12-06 12:43:20.310 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-06 12:43:20.310 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-06 12:43:20.313 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-06 12:43:20.317 UTC [29] LOG:  database system was shut down at 2025-12-06 12:42:32 UTC
tts-db          | 2025-12-06 12:43:20.322 UTC [1] LOG:  database system is ready to accept connections
python-service  | INFO:     127.0.0.1:59332 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34554 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 36.13625407218933
python-service  |  > Real-time factor: 4.111308110561869
python-service  |  > Text splitted to sentences.
python-service  | ['Nia kept the place steady, like a lighthouse keeper who never expects the sea to listen but keeps the lamp lit anyway.']
python-service  | INFO:     127.0.0.1:43974 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54382 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48820 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 34.21591830253601
python-service  |  > Real-time factor: 3.9608410256768116
python-service  |  > Text splitted to sentences.
python-service  | ['She shelved books, brewed bad coffee that somehow tasted better than it deserved, and watched the town breathe in and out through stories.']
python-service  | INFO:     127.0.0.1:36142 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:60386 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59182 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42526 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:44528 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 43.88517713546753
python-service  |  > Real-time factor: 4.042461048045999
python-service  |  > Text splitted to sentences.
python-service  | ['One ordinary Thursday, when the sky wore a sheet of grey as if afraid to spill its color, Nia found a box tucked behind a stack of atlases she hadn’t touched in years.']
python-service  | INFO:     127.0.0.1:59168 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:49170 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:44548 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58032 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40250 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 53.76938438415527
python-service  |  > Real-time factor: 4.030510353789175
python-service  |  > Text splitted to sentences.
python-service  | ['The box was a quiet thing, leather-bound, its stitching frayed by someone who handled it with care and not hurry Inside lay a single notebook: weathered, the corners softened as if someone had patted them with affection.']
python-service  | INFO:     127.0.0.1:42818 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42952 - "GET /health HTTP/1.1" 200 OK
tts-db          | 2025-12-06 12:48:20.354 UTC [27] LOG:  checkpoint starting: time
tts-db          | 2025-12-06 12:48:20.971 UTC [27] LOG:  checkpoint complete: wrote 9 buffers (0.1%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.605 s, sync=0.004 s, total=0.617 s; sync files=8, longest=0.002 s, average=0.001 s; distance=10 kB, estimate=10 kB
python-service  | INFO:     127.0.0.1:59670 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40806 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50722 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56238 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47614 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 73.05132675170898
python-service  |  > Real-time factor: 4.142104903505408
python-service  |  > Text splitted to sentences.
python-service  | ['On the cover, in gold leaf that had begun to flake, were the words: Book of Unsent Letters.']
python-service  | INFO:     127.0.0.1:36186 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34156 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46662 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34388 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59110 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 44.75097703933716
python-service  |  > Real-time factor: 4.010693908586624
python-service  |  > Text splitted to sentences.
python-service  | ['A note accompanied the box, written in the same careful hand she’d learned to recognize from her grandmother, a woman who had written more letters than anyone she knew and then never mailed them The note read, in part: “Dear you.']
python-service  | INFO:     127.0.0.1:52854 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43310 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45306 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58290 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40718 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:55742 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54316 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 75.46030759811401
python-service  |  > Real-time factor: 4.147638352357152
python-service  |  > Text splitted to sentences.
python-service  | ['If you’re reading this, you’re the listening sort Write to the ones who never hear back The world is rumor; letters are proof we cared enough to speak anyway ”', 'Nia frowned at the odd sentiment, then opened the notebook.']
python-service  | INFO:     127.0.0.1:47894 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45370 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:39514 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56292 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43898 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47428 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 59.064385175704956
python-service  |  > Real-time factor: 4.034002667274675
python-service  |  > Text splitted to sentences.
python-service  | ['The pages were blank, as blank as a stretch of winter sky She turned to the first page and found a single line, penned in a script she didn’t recognize at first It said: Dear you That was all.']
python-service  | INFO:     127.0.0.1:53334 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52226 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52208 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43956 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36404 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45346 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48508 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:44460 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 77.56732130050659
python-service  |  > Real-time factor: 4.126518612903325
python-service  |  > Text splitted to sentences.
python-service  | ['“Grandmother?” she whispered, because the voice in the back of her head sounded like hers but older, warmer, more certain.']
python-service  | INFO:     127.0.0.1:54490 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50818 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42884 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37344 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37698 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 47.64845609664917
python-service  |  > Real-time factor: 3.99595500262853
python-service  |  > Text splitted to sentences.
python-service  | ['The next moment the page warmed under her fingertips, and the letters appeared—not printed, but written, as if someone had tapped them into existence with a light touch.']
python-service  | INFO:     127.0.0.1:45334 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42870 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35680 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56780 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52538 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34984 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56680 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 75.04346537590027
python-service  |  > Real-time factor: 4.0907905431415905
python-service  |  > Text splitted to sentences.
python-service  | ['The ink glowed faintly, the lines written in a rhythm she could feel in her bones: Dear you I hope you’re listening Her breath caught.']
python-service  | INFO:     127.0.0.1:57990 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58640 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52058 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46972 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54008 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38452 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 56.46030640602112
python-service  |  > Real-time factor: 3.9631925719858327
python-service  |  > Text splitted to sentences.
python-service  | ['The notebook wouldn’t spill its secrets all at once, not like a book that promised adventures but offered only a page of dust.']
python-service  | INFO:     127.0.0.1:52980 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:33240 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35564 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 37.908650636672974
python-service  |  > Real-time factor: 3.9101009773718243
python-service  |  > Text splitted to sentences.
python-service  | ['She closed her eyes, pressed the book to her chest for a heartbeat, and when she opened them again the room’s silence had shifted to a soft, almost conspiratorial noise—the sound of someone nearby who wished to be found.']
python-service  | INFO:     127.0.0.1:54202 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48494 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35282 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37258 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:33842 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:41336 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48866 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35332 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45298 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48666 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 98.4654610157013
python-service  |  > Real-time factor: 4.085188712474436
python-service  |  > Text splitted to sentences.
python-service  | ['That night, after the last customer had drifted out and the town had settled into its habit of quiet, Nia opened the notebook once more and read the words she had “written.']
python-service  | INFO:     127.0.0.1:39792 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34148 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36752 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:57604 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43662 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42052 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45096 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 73.46302366256714
python-service  |  > Real-time factor: 4.095353322477866
python-service  | INFO:     127.0.0.1:47778 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Text splitted to sentences.
python-service  | ['” They weren’t addressed to her, not exactly; they felt more like messages flung across a wide harbor to boats that didn’t know they were part of a voyage yet.']
python-service  | INFO:     127.0.0.1:47172 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42426 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43198 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 39.39248609542847
python-service  |  > Real-time factor: 3.940463809266339
python-service  |  > Text splitted to sentences.
python-service  | ['She read aloud to the empty room, whispering a name she had barely ever spoken aloud: Mrs Calder, who sold seashells on the pier every Sunday with a smile that crept only at the edges of her lips.']
python-service  | INFO:     127.0.0.1:44806 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45634 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43308 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36416 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43932 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:51720 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42182 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 64.05183601379395
python-service  |  > Real-time factor: 4.038588850551759
python-service  |  > Text splitted to sentences.
python-service  | ['Dear you, she read, and the ink seemed to brighten I’ve been waiting for the courage to say your name aloud The sea remembers you, and so do we She blinked, then looked toward the door as if Mrs.']
python-service  | INFO:     127.0.0.1:52354 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40682 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50878 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:32810 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47352 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38628 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:33898 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42896 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 86.19009923934937
python-service  |  > Real-time factor: 4.11048657349304
python-service  |  > Text splitted to sentences.
python-service  | ['Calder might appear in a moment, seashells clacking in a tin bucket But of course the door stayed shut, and the shells stayed in the bucket.']
python-service  | INFO:     127.0.0.1:55338 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42592 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50736 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59422 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 37.16489815711975
python-service  |  > Real-time factor: 3.9130472360593367
python-service  |  > Text splitted to sentences.
python-service  | ['Still, the words trembled in the air as if they had just stepped off the page and onto the wooden floor The next morning, a letter, or perhaps a whisper of one, found its way into her routine.']
python-service  | INFO:     127.0.0.1:37456 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45370 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43928 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36170 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56570 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54512 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 55.33890128135681
python-service  |  > Real-time factor: 3.981826519520172
python-service  |  > Text splitted to sentences.
python-service  | ['The bakery at the end of the street ran out of croissants a half hour early because Mrs.']
python-service  | INFO:     127.0.0.1:60276 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:57886 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 24.846752643585205
python-service  |  > Real-time factor: 3.9335934505388694
python-service  |  > Text splitted to sentences.
python-service  | ['Calder herself stopped by the library to tell Nia that a small miracle had just occurred on the pier: a man she’d once helped with a broken boat had returned to claim a memory he’d forgotten to collect—the memory of his late wife.']
python-service  | INFO:     127.0.0.1:53524 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:55818 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38396 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38764 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37164 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52106 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36220 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 74.13705492019653
python-service  |  > Real-time factor: 4.1222565588822215
python-service  |  > Text splitted to sentences.
python-service  | ['a memory Mrs Calder had once woven into a shell necklace for him “He laughed,” she said, “and cried at the same time It was strange and true.']
python-service  | INFO:     127.0.0.1:39630 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40782 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37052 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56712 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:49232 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58050 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47306 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37812 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 76.9583044052124
python-service  |  > Real-time factor: 4.06150818590102
python-service  |  > Text splitted to sentences.
python-service  | ['” She left the bakery tray behind, and the library gained a new scent—sea salt and warm butter—that clung to the shelves Nia’s breath found its own rhythm, and she kept writing.']
python-service  | INFO:     127.0.0.1:33678 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43780 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40352 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37586 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54380 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:36458 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45304 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:42476 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 76.28046917915344
python-service  |  > Real-time factor: 4.113960066822715
python-service  |  > Text splitted to sentences.
python-service  | ['She decided to test whether the notebook really connected strangers through their words She wrote a letter for a neighbor she hardly spoke to, a quiet man who repaired bikes and never spoke of his own life beyond “the usual.']
python-service  | INFO:     127.0.0.1:51242 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:33428 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34908 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54816 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:51690 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38326 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46412 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 71.32333183288574
python-service  |  > Real-time factor: 4.128374424889565
python-service  |  > Text splitted to sentences.
python-service  | ['” She wrote to him about a dog that needed a home and a child who would love a bicycle more than a new toy that squealed when pushed.']
python-service  | INFO:     127.0.0.1:45616 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56254 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:53708 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 29.83959937095642
python-service  |  > Real-time factor: 4.021779744068393
python-service  |  > Text splitted to sentences.
python-service  | ['She left the letter open on the counter of the bicycle repair shop, tucked between spare tires and a jar of greasy rags That afternoon, the man returned with a bicycle that sounded like a tired tune when he pedaled.']
python-service  | INFO:     127.0.0.1:42300 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59382 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:51250 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43356 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40728 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:41338 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 63.93143677711487
python-service  |  > Real-time factor: 4.042835374132126
python-service  |  > Text splitted to sentences.
python-service  | ['Behind him walked a little girl, cheeks flushed with the kind of excitement only a first proper ride can grant She held up a bike bell she’d painted red with black stripes to mimic a dragon, and she laughed when the bell jingled.']
python-service  | INFO:     127.0.0.1:45182 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:32906 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:54748 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:38714 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34818 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:49788 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34612 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:35204 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 81.51434063911438
python-service  |  > Real-time factor: 4.105732637450368
python-service  |  > Text splitted to sentences.
python-service  | ['The man saw the letter on the counter and his eyes widened with recognition as if he’d always known someone had been thinking about him, even from miles away.']
python-service  | INFO:     127.0.0.1:49138 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45322 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59798 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:40720 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Processing time: 39.783212423324585
python-service  |  > Real-time factor: 3.9383837095678613
python-service  |  > Text splitted to sentences.
python-service  | ['He pockets the note as if it’s a charm and carefully empties a second, unused shelf in the shop, which suddenly seems brighter, more orderly than it did before.']
python-service  | INFO:     127.0.0.1:41054 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:52952 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:48832 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43972 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:46348 - "GET /health HTTP/1.1" 200 OK
go-service      | 2025/12/06 13:14:31 voice clone worker: job 4 failed: python service request failed: Post "http://python-service:5000/voice-clone": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
go-service      |
go-service      | 2025/12/06 13:14:31 /app/main.go:656 record not found
go-service      | [0.375ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:33 /app/main.go:656 record not found
go-service      | [0.327ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:35 /app/main.go:656 record not found
go-service      | [0.363ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:37 /app/main.go:656 record not found
go-service      | [0.411ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  |  > Processing time: 54.28208041191101
python-service  |  > Real-time factor: 3.87985540520019
python-service  | INFO:     127.0.0.1:46830 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:14:39 /app/main.go:656 record not found
go-service      | [0.439ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:41 /app/main.go:656 record not found
go-service      | [0.563ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:43 /app/main.go:656 record not found
go-service      | [0.349ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:45 /app/main.go:656 record not found
go-service      | [0.380ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:47 /app/main.go:656 record not found
go-service      | [0.562ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:49 /app/main.go:656 record not found
go-service      | [0.587ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:54992 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:14:51 /app/main.go:656 record not found
go-service      | [0.616ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:53 /app/main.go:656 record not found
go-service      | [0.603ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:55 /app/main.go:656 record not found
go-service      | [0.357ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:57 /app/main.go:656 record not found
go-service      | [0.343ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:14:59 /app/main.go:656 record not found
go-service      | [0.674ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:34320 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:15:01 /app/main.go:656 record not found
go-service      | [0.627ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:03 /app/main.go:656 record not found
go-service      | [0.570ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:05 /app/main.go:656 record not found
go-service      | [0.639ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:07 /app/main.go:656 record not found
go-service      | [0.457ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:09 /app/main.go:656 record not found
go-service      | [0.581ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:56378 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:15:11 /app/main.go:656 record not found
go-service      | [1.097ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:13 /app/main.go:656 record not found
go-service      | [0.697ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:15 /app/main.go:656 record not found
go-service      | [0.564ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:17 /app/main.go:656 record not found
go-service      | [0.630ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:19 /app/main.go:656 record not found
go-service      | [0.591ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:53510 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:15:21 /app/main.go:656 record not found
go-service      | [0.512ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:23 /app/main.go:656 record not found
go-service      | [0.532ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:25 /app/main.go:656 record not found
go-service      | [0.314ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:27 /app/main.go:656 record not found
go-service      | [0.485ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:29 /app/main.go:656 record not found
go-service      | [0.497ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:37822 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:15:31 /app/main.go:656 record not found
go-service      | [0.690ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:33 /app/main.go:656 record not found
go-service      | [0.505ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:35 /app/main.go:656 record not found
go-service      | [0.556ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:37 /app/main.go:656 record not found
go-service      | [0.583ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:39 /app/main.go:656 record not found
go-service      | [0.579ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:57034 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:15:41 /app/main.go:656 record not found
go-service      | [0.606ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:43 /app/main.go:656 record not found
go-service      | [0.477ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:45 /app/main.go:656 record not found
go-service      | [0.502ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:47 /app/main.go:656 record not found
go-service      | [0.716ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:49 /app/main.go:656 record not found
go-service      | [0.505ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:54586 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:15:51 /app/main.go:656 record not found
go-service      | [0.661ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:53 /app/main.go:656 record not found
go-service      | [1.115ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:55 /app/main.go:656 record not found
go-service      | [0.520ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:57 /app/main.go:656 record not found
go-service      | [0.631ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:15:59 /app/main.go:656 record not found
go-service      | [0.520ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:42206 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:16:01 /app/main.go:656 record not found
go-service      | [0.669ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:16:03 /app/main.go:656 record not found
go-service      | [0.602ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:16:05 /app/main.go:656 record not found
go-service      | [0.564ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:16:07 /app/main.go:656 record not found
go-service      | [0.681ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:16:09 /app/main.go:656 record not found
go-service      | [0.469ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:37212 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 13:16:11 /app/main.go:656 record not found
go-service      | [0.597ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:16:13 /app/main.go:656 record not found
go-service      | [0.536ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 13:16:15 /app/main.go:656 record not found
go-service      | [0.423ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED