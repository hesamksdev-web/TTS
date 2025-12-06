root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs
python-service  | /usr/local/lib/python3.10/site-packages/jieba/_compat.py:18: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   import pkg_resources
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  | ERROR:tts_service:Failed to pre-load engines: 'cv'
python-service  | INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  | INFO:     127.0.0.1:44782 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:49724 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47754 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45440 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59788 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56848 - "GET /health HTTP/1.1" 200 OK
tts-db          |
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db          |
tts-db          | 2025-12-06 11:38:22.601 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 15.2.0) 15.2.0, 64-bit
tts-db          | 2025-12-06 11:38:22.601 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-06 11:38:22.601 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-06 11:38:22.608 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-06 11:38:22.613 UTC [29] LOG:  database system was shut down at 2025-12-06 11:34:26 UTC
caddy-proxy     | {"level":"info","ts":1765021139.2788,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
tts-frontend    |
tts-frontend    | > frontend@0.1.0 start
caddy-proxy     | {"level":"info","ts":1765021139.2789428,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1765021139.2789896,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1765021139.2801132,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy     | {"level":"warn","ts":1765021139.280128,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1765021139.2804484,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy     | {"level":"info","ts":1765021139.2805667,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
tts-frontend    | > next start
caddy-proxy     | {"level":"info","ts":1765021139.2808456,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc0001b2180"}
tts-frontend    |
caddy-proxy     | {"level":"warn","ts":1765021139.2809296,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
tts-frontend    |    ▲ Next.js 16.0.6
caddy-proxy     | {"level":"warn","ts":1765021139.280935,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
tts-frontend    |    - Local:         http://localhost:3000
caddy-proxy     | {"level":"info","ts":1765021139.2809374,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
tts-frontend    |    - Network:       http://172.18.0.5:3000
caddy-proxy     | {"level":"info","ts":1765021139.2811449,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1765021139.2811635,"msg":"serving initial configuration"}
caddy-proxy     | {"level":"info","ts":1765021139.2830472,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765107539.2830462,"try_again_in":86399.999999649}
caddy-proxy     | {"level":"info","ts":1765021139.283124,"logger":"tls","msg":"finished cleaning storage units"}
go-service      | 2025/12/06 11:38:53 Go Service (SaaS Edition) listening on :8080
go-service      | 2025/12/06 11:38:53 Voice clone worker started
go-service      |
go-service      | 2025/12/06 11:38:53 /app/main.go:656 record not found
go-service      | [1.893ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:38:55 /app/main.go:656 record not found
go-service      | [0.938ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:38:57 /app/main.go:656 record not found
go-service      | [0.738ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:38:59 /app/main.go:656 record not found
go-service      | [0.669ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:01 /app/main.go:656 record not found
go-service      | [0.766ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:03 /app/main.go:656 record not found
go-service      | [0.681ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:05 /app/main.go:656 record not found
go-service      | [0.570ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:07 /app/main.go:656 record not found
go-service      | [0.314ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:09 /app/main.go:656 record not found
go-service      | [0.507ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:11 /app/main.go:656 record not found
go-service      | [0.543ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:13 /app/main.go:656 record not found
go-service      | [0.537ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:15 /app/main.go:656 record not found
go-service      | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:17 /app/main.go:656 record not found
go-service      | [0.568ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:19 /app/main.go:656 record not found
go-service      | [0.334ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:21 /app/main.go:656 record not found
go-service      | [0.517ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:23 /app/main.go:656 record not found
go-service      | [0.496ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:25 /app/main.go:656 record not found
go-service      | [0.483ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:27 /app/main.go:656 record not found
go-service      | [0.527ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:29 /app/main.go:656 record not found
go-service      | [0.627ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
tts-db          | 2025-12-06 11:38:22.618 UTC [1] LOG:  database system is ready to accept connections
tts-db          | 2025-12-06 11:43:22.708 UTC [27] LOG:  checkpoint starting: time
tts-db          | 2025-12-06 11:43:28.144 UTC [27] LOG:  checkpoint complete: wrote 57 buffers (0.3%); 0 WAL file(s) added, 0 removed, 0 recycled; write=5.423 s, sync=0.008 s, total=5.437 s; sync files=43, longest=0.002 s, average=0.001 s; distance=195 kB, estimate=195 kB
tts-frontend    |
go-service      |
go-service      | 2025/12/06 11:39:31 /app/main.go:656 record not found
go-service      | [0.455ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
tts-frontend    |  ✓ Starting...
python-service  | INFO:     127.0.0.1:37832 - "GET /health HTTP/1.1" 200 OK
tts-frontend    |  ✓ Ready in 305ms
python-service  | INFO:     127.0.0.1:34554 - "GET /health HTTP/1.1" 200 OK
go-service      |
go-service      | 2025/12/06 11:39:33 /app/main.go:656 record not found
go-service      | [0.530ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:35 /app/main.go:656 record not found
go-service      | [0.528ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:37 /app/main.go:656 record not found
go-service      | [0.418ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
python-service  | INFO:     127.0.0.1:50128 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Downloading model to /root/.local/share/tts/tts/tts_models--multilingual--multi-dataset--xtts_v2
100%|█████████▉| 1.86G/1.87G [00:20<00:00, 87.9MiB/s]
go-service      |
go-service      | 2025/12/06 11:39:39 /app/main.go:656 record not found
go-service      | [0.586ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:41 /app/main.go:656 record not found
go-service      | [0.522ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:43 /app/main.go:656 record not found
go-service      | [0.470ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:45 /app/main.go:656 record not found
go-service      | [0.408ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:47 /app/main.go:656 record not found
go-service      | [0.436ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:49 /app/main.go:656 record not found
go-service      | [0.693ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:51 /app/main.go:656 record not found
go-service      | [0.453ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:53 /app/main.go:656 record not found
go-service      | [0.559ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:55 /app/main.go:656 record not found
go-service      | [0.513ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:57 /app/main.go:656 record not found
go-service      | [0.412ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:59 /app/main.go:656 record not found
go-service      | [0.647ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:01 /app/main.go:656 record not found
go-service      | [0.455ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:03 /app/main.go:656 record not found
go-service      | [0.494ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:05 /app/main.go:656 record not found
go-service      | [0.533ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:07 /app/main.go:656 record not found
go-service      | [0.623ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:09 /app/main.go:656 record not found
go-service      | [0.382ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:11 /app/main.go:656 record not found
go-service      | [0.550ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:13 /app/main.go:656 record not found
go-service      | [0.477ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:15 /app/main.go:656 record not found
go-service      | [0.630ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:17 /app/main.go:656 record not found
go-service      | [0.520ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      | 2025/12/06 11:43:19 voice clone worker: job 1 failed: python service request failed: Post "http://python-service:5000/voice-clone": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
go-service      |
go-service      | 2025/12/06 11:43:19 /app/main.go:656 record not found
go-service      | [0.281ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:21 /app/main.go:656 record not found
go-service      | [0.304ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:23 /app/main.go:656 record not found
go-service      | [0.645ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:25 /app/main.go:656 record not found
go-service      | [0.450ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:27 /app/main.go:656 record not found
go-service      | [0.301ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE100%|██████████| 1.87G/1.87G [00:21<00:00, 88.7MiB/s]
4.37kiB [00:00, 21.1kiB/s]
361kiB [00:00, 1.71MiB/s]
100%|██████████| 32.0/32.0 [00:00<00:00, 131iB/s]
python-service  |  > Model's license - CPML
python-service  |  > Check https://coqui.ai/cpml.txt for more info.
python-service  |  > Using model: xtts
python-service  |  > Text splitted to sentences.
python-service  | INFO:     127.0.0.1:55742 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:39494 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45622 - "GET /health HTTP/1.1" 200 OK
python-service  | ['The town library sits where the coastline curls like a question mark, and in it the air always smells faintly of rain and old paper.', 'Nia kept the place steady, like a lighthouse keeper who never expects the sea to listen but keeps the lamp lit anyway.', 'She shelved books, brewed bad coffee that somehow tasted better than it deserved, and watched the town breathe in and out through stories.', 'One ordinary Thursday, when the sky wore a sheet of grey as if afraid to spill its color, Nia found a box tucked behind a stack of atlases she hadn’t touched in years.', 'The box was a quiet thing, leather-bound, its stitching frayed by someone who handled it with care and not hurry.', 'Inside lay a single notebook: weathered, the corners softened as if someone had patted them with affection.', 'On the cover, in gold leaf that had begun to flake, were the words: Book of Unsent Letters.', 'A note accompanied the box, written in the same careful hand she’d learned to recognize from her grandmother, a woman who had written more letters than anyone she knew and then never mailed them.', 'The note read, in part: “Dear you. If you’re reading this, you’re the listening sort. Write to the ones who never hear back. The world is rumor; letters are proof we cared enough to speak anyway.”', 'Nia frowned at the odd sentiment, then opened the notebook.', 'The pages were blank, as blank as a stretch of winter sky.', 'She turned to the first page and found a single line, penned in a script she didn’t recognize at first.', 'It said: Dear you.', 'That was all.', '“Grandmother?” she whispered, because the voice in the back of her head sounded like hers but older, warmer, more certain.', 'The next moment the page warmed under her fingertips, and the letters appeared—not printed, but written, as if someone had tapped them into existence with a light touch.', 'The ink glowed faintly, the lines written in a rhythm she could feel in her bones: Dear you.', 'I hope you’re listening.', 'Her breath caught.', 'The notebook wouldn’t spill its secrets all at once, not like a book that promised adventures but offered only a page of dust.', 'She closed her eyes, pressed the book to her chest for a heartbeat, and when she opened them again the room’s silence had shifted to a soft, almost conspiratorial noise—the sound of someone nearby who wished to be found.', 'That night, after the last customer had drifted out and the town had settled into its habit of quiet, Nia opened the notebook once more and read the words she had “written.”', 'They weren’t addressed to her, not exactly; they felt more like messages flung across a wide harbor to boats that didn’t know they were part of a voyage yet.', 'She read aloud to the empty room, whispering a name she had barely ever spoken aloud: Mrs. Calder, who sold seashells on the pier every Sunday with a smile that crept only at the edges of her lips.', 'Dear you, she read, and the ink seemed to brighten.', 'I’ve been waiting for the courage to say your name aloud.', 'The sea remembers you, and so do we.', 'She blinked, then looked toward the door as if Mrs. Calder might appear in a moment, seashells clacking in a tin bucket.', 'But of course the door stayed shut, and the shells stayed in the bucket.', 'Still, the words trembled in the air as if they had just stepped off the page and onto the wooden floor.', 'The next morning, a letter, or perhaps a whisper of one, found its way into her routine.', 'The bakery at the end of the street ran out of croissants a half hour early because Mrs. Calder herself stopped by the library to tell Nia that a small miracle had just occurred on the pier: a man she’d once helped with a broken boat had returned to claim a memory he’d forgotten to collect—the memory of his late wife, a memory Mrs. Calder had once woven into a shell necklace for him.', '“He laughed,” she said, “and cried at the same time. It was strange and true.”', 'She left the bakery tray behind, and the library gained a new scent—sea salt and warm butter—that clung to the shelves.', 'Nia’s breath found its own rhythm, and she kept writing.', 'She decided to test whether the notebook really connected strangers through their words.', 'She wrote a letter for a neighbor she hardly spoke to, a quiet man who repaired bikes and never spoke of his own life beyond “the usual.”', 'She wrote to him about a dog that needed a home and a child who would love a bicycle more than a new toy that squealed when pushed.', 'She left the letter open on the counter of the bicycle repair shop, tucked between spare tires and a jar of greasy rags.', 'That afternoon, the man returned with a bicycle that sounded like a tired tune when he pedaled.', 'Behind him walked a little girl, cheeks flushed with the kind of excitement only a first proper ride can grant.', 'She held up a bike bell she’d painted red with black stripes to mimic a dragon, and she laughed when the bell jingled.', 'The man saw the letter on the counter and his eyes widened with recognition as if he’d always known someone had been thinking about him, even from miles away.', 'He pockets the note as if it’s a charm and carefully empties a second, unused shelf in the shop, which suddenly seems brighter, more orderly than it did before.']
python-service  | INFO:     127.0.0.1:44682 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:53724 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:57420 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47544 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43906 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45048 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58984 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43210 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59020 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58462 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:55842 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37656 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:51458 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47308 - "GET /health HTTP/1.1" 200 OK
root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs
go-service  | 2025/12/06 11:38:53 Go Service (SaaS Edition) listening on :8080
go-service  | 2025/12/06 11:38:53 Voice clone worker started
go-service  |
go-service  | 2025/12/06 11:38:53 /app/main.go:656 record not found
go-service  | [1.893ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:38:55 /app/main.go:656 record not found
go-service  | [0.938ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:38:57 /app/main.go:656 record not found
go-service  | [0.738ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:38:59 /app/main.go:656 record not found
go-service  | [0.669ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:01 /app/main.go:656 record not found
go-service  | [0.766ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:03 /app/main.go:656 record not found
go-service  | [0.681ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:05 /app/main.go:656 record not found
go-service  | [0.570ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:07 /app/main.go:656 record not found
go-service  | [0.314ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:09 /app/main.go:656 record not found
go-service  | [0.507ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:11 /app/main.go:656 record not found
go-service  | [0.543ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:13 /app/main.go:656 record not found
go-service  | [0.537ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:15 /app/main.go:656 record not found
go-service  | [0.576ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:17 /app/main.go:656 record not found
go-service  | [0.568ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service  |
go-service  | 2025/12/06 11:39:19 /app/main.go:656 record not found
python-service  | /usr/local/lib/python3.10/site-packages/jieba/_compat.py:18: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   import pkg_resources
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  | ERROR:tts_service:Failed to pre-load engines: 'cv'
python-service  | INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  | INFO:     127.0.0.1:44782 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:49724 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47754 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45440 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:59788 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:56848 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:37832 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:34554 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:50128 - "GET /health HTTP/1.1" 200 OK
python-service  |  > Downloading model to /root/.local/share/tts/tts/tts_models--multilingual--multi-dataset--xtts_v2
100%|█████████▉| 1.86G/1.87G [00:20<00:00, 87.9MiB/s]
caddy-proxy     | {"level":"info","ts":1765021139.2788,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy     | {"level":"info","ts":1765021139.2789428,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1765021139.2789896,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1765021139.2801132,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy     | {"level":"warn","ts":1765021139.280128,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1765021139.2804484,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy     | {"level":"info","ts":1765021139.2805667,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy     | {"level":"info","ts":1765021139.2808456,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc0001b2180"}
caddy-proxy     | {"level":"warn","ts":1765021139.2809296,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"warn","ts":1765021139.280935,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"info","ts":1765021139.2809374,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy     | {"level":"info","ts":1765021139.2811449,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1765021139.2811635,"msg":"serving initial configuration"}
caddy-proxy     | {"level":"info","ts":1765021139.2830472,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765107539.2830462,"try_again_in":86399.999999649}
100%|██████████| 1.87G/1.87G [00:21<00:00, 88.7MiB/s]124,"logger":"tls","msg":"finished cleaning storage units"}
4.37kiB [00:00, 21.1kiB/s]
361kiB [00:00, 1.71MiB/s]
100%|██████████| 32.0/32.0 [00:00<00:00, 131iB/s]
python-service  |  > Model's license - CPML
python-service  |  > Check https://coqui.ai/cpml.txt for more info.
python-service  |  > Using model: xtts
python-service  |  > Text splitted to sentences.
python-service  | INFO:     127.0.0.1:55742 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:39494 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45622 - "GET /health HTTP/1.1" 200 OK
python-service  | ['The town library sits where the coastline curls like a question mark, and in it the air always smells faintly of rain and old paper.', 'Nia kept the place steady, like a lighthouse keeper who never expects the sea to listen but keeps the lamp lit anyway.', 'She shelved books, brewed bad coffee that somehow tasted better than it deserved, and watched the town breathe in and out through stories.', 'One ordinary Thursday, when the sky wore a sheet of grey as if afraid to spill its color, Nia found a box tucked behind a stack of atlases she hadn’t touched in years.', 'The box was a quiet thing, leather-bound, its stitching frayed by someone who handled it with care and not hurry.', 'Inside lay a single notebook: weathered, the corners softened as if someone had patted them with affection.', 'On the cover, in gold leaf that had begun to flake, were the words: Book of Unsent Letters.', 'A note accompanied the box, written in the same careful hand she’d learned to recognize from her grandmother, a woman who had written more letters than anyone she knew and then never mailed them.', 'The note read, in part: “Dear you. If you’re reading this, you’re the listening sort. Write to the ones who never hear back. The world is rumor; letters are proof we cared enough to speak anyway.”', 'Nia frowned at the odd sentiment, then opened the notebook.', 'The pages were blank, as blank as a stretch of winter sky.', 'She turned to the first page and found a single line, penned in a script she didn’t recognize at first.', 'It said: Dear you.', 'That was all.', '“Grandmother?” she whispered, because the voice in the back of her head sounded like hers but older, warmer, more certain.', 'The next moment the page warmed under her fingertips, and the letters appeared—not printed, but written, as if someone had tapped them into existence with a light touch.', 'The ink glowed faintly, the lines written in a rhythm she could feel in her bones: Dear you.', 'I hope you’re listening.', 'Her breath caught.', 'The notebook wouldn’t spill its secrets all at once, not like a book that promised adventures but offered only a page of dust.', 'She closed her eyes, pressed the book to her chest for a heartbeat, and when she opened them again the room’s silence had shifted to a soft, almost conspiratorial noise—the sound of someone nearby who wished to be found.', 'That night, after the last customer had drifted out and the town had settled into its habit of quiet, Nia opened the notebook once more and read the words she had “written.”', 'They weren’t addressed to her, not exactly; they felt more like messages flung across a wide harbor to boats that didn’t know they were part of a voyage yet.', 'She read aloud to the empty room, whispering a name she had barely ever spoken aloud: Mrs. Calder, who sold seashells on the pier every Sunday with a smile that crept only at the edges of her lips.', 'Dear you, she read, and the ink seemed to brighten.', 'I’ve been waiting for the courage to say your name aloud.', 'The sea remembers you, and so do we.', 'She blinked, then looked toward the door as if Mrs. Calder might appear in a moment, seashells clacking in a tin bucket.', 'But of course the door stayed shut, and the shells stayed in the bucket.', 'Still, the words trembled in the air as if they had just stepped off the page and onto the wooden floor.', 'The next morning, a letter, or perhaps a whisper of one, found its way into her routine.', 'The bakery at the end of the street ran out of croissants a half hour early because Mrs. Calder herself stopped by the library to tell Nia that a small miracle had just occurred on the pier: a man she’d once helped with a broken boat had returned to claim a memory he’d forgotten to collect—the memory of his late wife, a memory Mrs. Calder had once woven into a shell necklace for him.', '“He laughed,” she said, “and cried at the same time. It was strange and true.”', 'She left the bakery tray behind, and the library gained a new scent—sea salt and warm butter—that clung to the shelves.', 'Nia’s breath found its own rhythm, and she kept writing.', 'She decided to test whether the notebook really connected strangers through their words.', 'She wrote a letter for a neighbor she hardly spoke to, a quiet man who repaired bikes and never spoke of his own life beyond “the usual.”', 'She wrote to him about a dog that needed a home and a child who would love a bicycle more than a new toy that squealed when pushed.', 'She left the letter open on the counter of the bicycle repair shop, tucked between spare tires and a jar of greasy rags.', 'That afternoon, the man returned with a bicycle that sounded like a tired tune when he pedaled.', 'Behind him walked a little girl, cheeks flushed with the kind of excitement only a first proper ride can grant.', 'She held up a bike bell she’d painted red with black stripes to mimic a dragon, and she laughed when the bell jingled.', 'The man saw the letter on the counter and his eyes widened with recognition as if he’d always known someone had been thinking about him, even from miles away.', 'He pockets the note as if it’s a charm and carefully empties a second, unused shelf in the shop, which suddenly seems brighter, more orderly than it did before.']
python-service  | INFO:     127.0.0.1:44682 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:53724 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:57420 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47544 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43906 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:45048 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58984 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:43210 - "GET /health HTTP/1.1" 200 OK
tts-frontend    |
tts-frontend    | > frontend@0.1.0 start
tts-frontend    | > next start
tts-frontend    |
go-service      | [0.334ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:21 /app/main.go:656 record not found
python-service  | INFO:     127.0.0.1:59020 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:58462 - "GET /health HTTP/1.1" 200 OK
tts-frontend    |    ▲ Next.js 16.0.6
python-service  | INFO:     127.0.0.1:55842 - "GET /health HTTP/1.1" 200 OK
tts-frontend    |    - Local:         http://localhost:3000
python-service  | INFO:     127.0.0.1:37656 - "GET /health HTTP/1.1" 200 OK
tts-frontend    |    - Network:       http://172.18.0.5:3000
python-service  | INFO:     127.0.0.1:51458 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:47308 - "GET /health HTTP/1.1" 200 OK
python-service  | INFO:     127.0.0.1:51460 - "GET /health HTTP/1.1" 200 OK
tts-frontend    |
tts-frontend    |  ✓ Starting...
tts-frontend    |  ✓ Ready in 305ms
go-service      | [0.517ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:23 /app/main.go:656 record not found
go-service      | [0.496ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:25 /app/main.go:656 record not found
go-service      | [0.483ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:27 /app/main.go:656 record not found
go-service      | [0.527ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:29 /app/main.go:656 record not found
go-service      | [0.627ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:31 /app/main.go:656 record not found
go-service      | [0.455ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:33 /app/main.go:656 record not found
go-service      | [0.530ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:35 /app/main.go:656 record not found
go-service      | [0.528ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:37 /app/main.go:656 record not found
go-service      | [0.418ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:39 /app/main.go:656 record not found
go-service      | [0.586ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:41 /app/main.go:656 record not found
go-service      | [0.522ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:43 /app/main.go:656 record not found
go-service      | [0.470ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:45 /app/main.go:656 record not found
go-service      | [0.408ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:47 /app/main.go:656 record not found
go-service      | [0.436ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:49 /app/main.go:656 record not found
go-service      | [0.693ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:51 /app/main.go:656 record not found
go-service      | [0.453ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:53 /app/main.go:656 record not found
go-service      | [0.559ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:55 /app/main.go:656 record not found
go-service      | [0.513ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:57 /app/main.go:656 record not found
go-service      | [0.412ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:39:59 /app/main.go:656 record not found
go-service      | [0.647ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:01 /app/main.go:656 record not found
go-service      | [0.455ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:03 /app/main.go:656 record not found
go-service      | [0.494ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:05 /app/main.go:656 record not found
go-service      | [0.533ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:07 /app/main.go:656 record not found
go-service      | [0.623ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:09 /app/main.go:656 record not found
go-service      | [0.382ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:11 /app/main.go:656 record not found
go-service      | [0.550ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:13 /app/main.go:656 record not found
go-service      | [0.477ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:15 /app/main.go:656 record not found
go-service      | [0.630ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:40:17 /app/main.go:656 record not found
go-service      | [0.520ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      | 2025/12/06 11:43:19 voice clone worker: job 1 failed: python service request failed: Post "http://python-service:5000/voice-clone": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
go-service      |
go-service      | 2025/12/06 11:43:19 /app/main.go:656 record not found
go-service      | [0.281ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:21 /app/main.go:656 record not found
go-service      | [0.304ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:23 /app/main.go:656 record not found
go-service      | [0.645ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:25 /app/main.go:656 record not found
go-service      | [0.450ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:27 /app/main.go:656 record not found
go-service      | [0.301ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
go-service      |
go-service      | 2025/12/06 11:43:29 /app/main.go:656 record not found
go-service      | [0.506ms] [rows:0] SELECT * FROM "voice_clone_jobs" WHERE status = 'pending' AND "voice_clone_jobs"."deleted_at" IS NULL ORDER BY created_at ASC LIMIT 1 FOR UPDATE SKIP LOCKED
tts-db          |
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db          |
tts-db          | 2025-12-06 11:38:22.601 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 15.2.0) 15.2.0, 64-bit
tts-db          | 2025-12-06 11:38:22.601 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-06 11:38:22.601 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-06 11:38:22.608 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-06 11:38:22.613 UTC [29] LOG:  database system was shut down at 2025-12-06 11:34:26 UTC
tts-db          | 2025-12-06 11:38:22.618 UTC [1] LOG:  database system is ready to accept connections
tts-db          | 2025-12-06 11:43:22.708 UTC [27] LOG:  checkpoint starting: time
tts-db          | 2025-12-06 11:43:28.144 UTC [27] LOG:  checkpoint complete: wrote 57 buffers (0.3%); 0 WAL file(s) added, 0 removed, 0 recycled; write=5.423 s, sync=0.008 s, total=5.437 s; sync files=43, longest=0.002 s, average=0.001 s; distance=195 kB, estimate=195 kB
root@Automations:/opt/TTS#