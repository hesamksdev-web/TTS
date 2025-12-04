root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml down
[+] Running 6/6
 ✔ Container caddy-proxy     Removed                                                                                                                                            0.2s
 ✔ Container tts-frontend    Removed                                                                                                                                            0.2s
 ✔ Container go-service      Removed                                                                                                                                            0.2s
 ✔ Container python-service  Removed                                                                                                                                            1.4s
 ✔ Container tts-db          Removed                                                                                                                                            0.2s
 ✔ Network tts_tts-network   Removed                                                                                                                                            0.1s
root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml up -d
[+] Running 6/6
 ✔ Network tts_tts-network   Created                                                                                                                                            0.0s
 ✔ Container python-service  Started                                                                                                                                            0.5s
 ✔ Container tts-db          Healthy                                                                                                                                           11.0s
 ✔ Container go-service      Started                                                                                                                                           11.1s
 ✔ Container tts-frontend    Started                                                                                                                                           11.3s
 ✔ Container caddy-proxy     Started                                                                                                                                           11.4s
root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs
python-service  | /usr/local/lib/python3.10/site-packages/librosa/core/intervals.py:8: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   from pkg_resources import resource_filename
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  | INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
python-service  | INFO:     172.18.0.4:48730 - "POST /tts HTTP/1.1" 400 Bad Request
tts-frontend    |
tts-frontend    | > frontend@0.1.0 start
tts-frontend    | > next start
tts-frontend    |
tts-frontend    |    ▲ Next.js 16.0.6
tts-frontend    |    - Local:         http://localhost:3000
tts-frontend    |    - Network:       http://172.18.0.5:3000
tts-frontend    |
tts-frontend    |  ✓ Starting...
tts-frontend    |  ✓ Ready in 377ms
go-service      | 2025/12/04 15:16:30 Go Service (SaaS Edition) listening on :8080
go-service      | 2025/12/04 15:16:45 Python service response: status=400, body={"detail":"speaker_id 'my_speaker' is not supported. Choose from: ['en_p225', 'de_thorsten', 'fa_cv']"}
caddy-proxy     | {"level":"info","ts":1764861390.3566453,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy     | {"level":"info","ts":1764861390.3569887,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy     | {"level":"info","ts":1764861390.3570225,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy     | {"level":"info","ts":1764861390.3580532,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy     | {"level":"warn","ts":1764861390.3580675,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy     | {"level":"warn","ts":1764861390.3584356,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy     | {"level":"info","ts":1764861390.3586493,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy     | {"level":"info","ts":1764861390.3588552,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc0006fe200"}
caddy-proxy     | {"level":"warn","ts":1764861390.35935,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"warn","ts":1764861390.3593655,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy     | {"level":"info","ts":1764861390.3593807,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
tts-db          |
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db          |
tts-db          | 2025-12-04 15:16:19.233 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 14.2.0) 14.2.0, 64-bit
tts-db          | 2025-12-04 15:16:19.233 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-04 15:16:19.233 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-04 15:16:19.236 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-04 15:16:19.241 UTC [28] LOG:  database system was shut down at 2025-12-04 15:16:13 UTC
tts-db          | 2025-12-04 15:16:19.249 UTC [1] LOG:  database system is ready to accept connections
caddy-proxy     | {"level":"info","ts":1764861390.3595867,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
caddy-proxy     | {"level":"info","ts":1764861390.3596196,"msg":"serving initial configuration"}
caddy-proxy     | {"level":"info","ts":1764861390.3645191,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1764947790.3645172,"try_again_in":86399.999999239}
caddy-proxy     | {"level":"info","ts":1764861390.3645923,"logger":"tls","msg":"finished cleaning storage units"}
root@Automations:/opt/TTS#