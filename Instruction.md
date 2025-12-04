root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml up -d
[+] Running 7/7
 ✔ Network tts_tts-network   Created                                                                                                                                            0.0s
 ✔ Volume tts_nginx_cache    Created                                                                                                                                            0.0s
 ✔ Container python-service  Started                                                                                                                                            0.6s
 ✔ Container tts-db          Healthy                                                                                                                                           11.1s
 ✔ Container go-service      Started                                                                                                                                           11.2s
 ✔ Container tts-frontend    Started                                                                                                                                           11.3s
 ✔ Container nginx-proxy     Started                                                                                                                                           11.4s
root@Automations:/opt/TTS# docker ps
CONTAINER ID   IMAGE                       COMMAND                  CREATED         STATUS                   PORTS                                                                          NAMES
e85ac3100332   nginx:alpine                "/docker-entrypoint.…"   4 minutes ago   Up 4 minutes             0.0.0.0:80->80/tcp, [::]:80->80/tcp, 0.0.0.0:443->443/tcp, [::]:443->443/tcp   nginx-proxy
551720d5f6c4   tts-frontend:latest         "docker-entrypoint.s…"   4 minutes ago   Up 4 minutes             3000/tcp                                                                       tts-frontend
54598ddb1241   tts-go-service:latest       "./go-service"           4 minutes ago   Up 2 seconds             8080/tcp                                                                       go-service
1548135cba72   postgres:15-alpine          "docker-entrypoint.s…"   4 minutes ago   Up 4 minutes (healthy)   5432/tcp                                                                       tts-db
d1e62c5d0aba   tts-python-service:latest   "uvicorn app.main:ap…"   4 minutes ago   Up 4 minutes             5000/tcp                                                                       python-service
root@Automations:/opt/TTS# client_loop: send disconnect: Connection reset

C:\Windows\System32>docker-compose logs
'docker-compose' is not recognized as an internal or external command,
operable program or batch file.

C:\Windows\System32>ssh bbauto@46.224.57.222
bbauto@46.224.57.222's password:
Welcome to Ubuntu 24.04.3 LTS (GNU/Linux 6.8.0-88-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/pro

 System information as of Thu Dec  4 12:04:52 PM CET 2025

  System load:  0.08               Processes:             196
  Usage of /:   54.3% of 74.79GB   Users logged in:       1
  Memory usage: 19%                IPv4 address for eth0: 46.224.57.222
  Swap usage:   1%                 IPv6 address for eth0: 2a01:4f8:1c1a:6823::1


Expanded Security Maintenance for Applications is not enabled.

9 updates can be applied immediately.
To see these additional updates run: apt list --upgradable

12 additional security updates can be applied with ESM Apps.
Learn more about enabling ESM Apps service at https://ubuntu.com/esm


Last login: Thu Dec  4 11:55:49 2025 from 4.180.63.136
bbauto@Automations:~$ sudo -i
[sudo] password for bbauto:
root@Automations:~# cd /opt/TTS/
root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs
nginx-proxy  | /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
nginx-proxy  | /docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
nginx-proxy  | /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
nginx-proxy  | 10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
nginx-proxy  | 10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
nginx-proxy  | /docker-entrypoint.sh: Sourcing /docker-entrypoint.d/15-local-resolvers.envsh
nginx-proxy  | /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
nginx-proxy  | /docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
nginx-proxy  | /docker-entrypoint.sh: Configuration complete; ready for start up
python-service  | /usr/local/lib/python3.10/site-packages/librosa/core/intervals.py:8: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
python-service  |   from pkg_resources import resource_filename
python-service  | INFO:     Started server process [1]
python-service  | INFO:     Waiting for application startup.
python-service  | INFO:     Application startup complete.
python-service  | INFO:     Uvicorn running on http://0.0.0.0:5000 (Press CTRL+C to quit)
tts-db          |
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db          |
tts-db          | 2025-12-04 10:59:47.716 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 14.2.0) 14.2.0, 64-bit
tts-db          | 2025-12-04 10:59:47.717 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-04 10:59:47.717 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-04 10:59:47.719 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-04 10:59:47.724 UTC [29] LOG:  database system was shut down at 2025-12-04 10:44:12 UTC
tts-db          | 2025-12-04 10:59:47.730 UTC [1] LOG:  database system is ready to accept connections
tts-db          | 2025-12-04 10:59:58.328 UTC [40] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 10:59:58.328 UTC [40] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:00.335 UTC [41] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:00.335 UTC [41] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:02.342 UTC [42] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:02.342 UTC [42] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:04.350 UTC [44] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:04.350 UTC [44] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:06.358 UTC [45] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:06.358 UTC [45] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:08.366 UTC [53] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:08.366 UTC [53] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:10.375 UTC [54] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:10.375 UTC [54] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:12.383 UTC [55] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:12.383 UTC [55] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:14.392 UTC [56] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:14.392 UTC [56] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:16.401 UTC [57] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:16.401 UTC [57] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:18.726 UTC [66] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:18.726 UTC [66] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:20.734 UTC [67] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:20.734 UTC [67] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:22.743 UTC [68] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:22.743 UTC [68] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:24.752 UTC [69] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:24.752 UTC [69] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:26.760 UTC [70] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:26.760 UTC [70] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:28.766 UTC [78] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:28.766 UTC [78] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:30.773 UTC [79] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:30.773 UTC [79] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:32.781 UTC [81] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:32.781 UTC [81] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:34.788 UTC [82] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:34.788 UTC [82] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:36.794 UTC [83] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:36.794 UTC [83] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:39.151 UTC [91] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:39.151 UTC [91] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:41.159 UTC [92] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:41.159 UTC [92] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:43.166 UTC [93] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:43.166 UTC [93] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:45.175 UTC [94] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:45.175 UTC [94] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:47.183 UTC [95] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:47.183 UTC [95] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:49.189 UTC [104] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:49.189 UTC [104] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:51.197 UTC [105] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:51.197 UTC [105] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:53.205 UTC [106] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:53.205 UTC [106] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:55.213 UTC [107] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:55.213 UTC [107] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:57.222 UTC [108] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:57.222 UTC [108] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:00:59.570 UTC [116] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:00:59.570 UTC [116] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:01.578 UTC [117] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:01.578 UTC [117] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:03.585 UTC [119] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:03.585 UTC [119] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:05.593 UTC [120] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:05.593 UTC [120] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:07.601 UTC [121] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:07.601 UTC [121] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:09.611 UTC [129] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:09.611 UTC [129] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:11.620 UTC [130] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:11.620 UTC [130] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:13.629 UTC [131] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:13.629 UTC [131] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:15.637 UTC [132] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:15.637 UTC [132] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:17.644 UTC [133] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:17.644 UTC [133] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:19.999 UTC [142] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:19.999 UTC [142] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:22.007 UTC [143] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:22.007 UTC [143] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:24.015 UTC [144] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:24.015 UTC [144] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
go-service      |
go-service      | 2025/12/04 10:59:58 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:00 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:02 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:04 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:06 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:08 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:10 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:12 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:14 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:16 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:00:18 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:18 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:20 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:22 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
tts-frontend    |
tts-frontend    | > frontend@0.1.0 start
tts-frontend    | > next start
go-service      |
tts-frontend    |
go-service      | 2025/12/04 11:00:24 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:26 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:28 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:30 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:32 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:34 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:36 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:00:38 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:39 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:41 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:43 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:45 /app/main.go:76
tts-db          | 2025-12-04 11:01:26.022 UTC [145] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:26.022 UTC [145] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:28.030 UTC [146] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:28.030 UTC [146] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:30.039 UTC [154] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:30.039 UTC [154] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:32.047 UTC [155] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:32.047 UTC [155] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:34.056 UTC [157] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:34.056 UTC [157] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:36.063 UTC [158] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:36.063 UTC [158] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:38.071 UTC [159] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:38.071 UTC [159] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:40.397 UTC [167] FATAL:  password authentication failed for user "admin"
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:47 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:49 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:51 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
tts-db          | 2025-12-04 11:01:40.397 UTC [167] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:42.404 UTC [168] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:42.404 UTC [168] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:44.412 UTC [169] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:44.412 UTC [169] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:46.419 UTC [170] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:46.419 UTC [170] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:48.427 UTC [172] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:48.427 UTC [172] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:50.435 UTC [180] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:50.435 UTC [180] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:52.442 UTC [181] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:52.442 UTC [181] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:54.449 UTC [182] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:54.449 UTC [182] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:53 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:55 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
tts-frontend    |    ▲ Next.js 16.0.6
tts-frontend    |    - Local:         http://localhost:3000
tts-frontend    |    - Network:       http://172.18.0.5:3000
tts-frontend    |
tts-frontend    |  ✓ Starting...
tts-frontend    |  ✓ Ready in 317ms
tts-db          | 2025-12-04 11:01:56.458 UTC [183] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:56.458 UTC [183] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:01:58.466 UTC [184] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:01:58.466 UTC [184] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:00.810 UTC [192] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:00.810 UTC [192] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:02.818 UTC [194] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:02.818 UTC [194] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:04.826 UTC [195] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:04.826 UTC [195] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:06.835 UTC [196] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:06.835 UTC [196] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:08.841 UTC [204] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:08.841 UTC [204] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:10.848 UTC [205] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:10.848 UTC [205] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:12.855 UTC [206] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:12.855 UTC [206] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:14.862 UTC [207] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:14.862 UTC [207] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:16.869 UTC [208] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:16.869 UTC [208] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:18.877 UTC [217] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:18.877 UTC [217] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:21.223 UTC [218] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:21.223 UTC [218] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:23.231 UTC [219] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:23.231 UTC [219] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:57 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
tts-db          | 2025-12-04 11:02:25.238 UTC [220] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:25.238 UTC [220] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:27.247 UTC [221] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:27.247 UTC [221] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:29.256 UTC [229] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:29.256 UTC [229] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:31.263 UTC [230] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:31.263 UTC [230] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:33.272 UTC [232] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:33.272 UTC [232] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:35.281 UTC [233] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:35.281 UTC [233] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:37.288 UTC [234] FATAL:  password authentication failed for user "admin"
go-service      | 2025/12/04 11:00:59 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:00:59 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:01 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:03 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:05 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:07 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:09 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:11 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:13 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:15 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:17 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:01:19 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:20 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:22 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:24 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:26 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:28 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:30 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:32 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:34 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:36 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
tts-db          | 2025-12-04 11:02:37.288 UTC [234] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:39.296 UTC [242] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:39.296 UTC [242] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:41.654 UTC [243] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:41.654 UTC [243] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:43.661 UTC [244] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:43.661 UTC [244] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:45.667 UTC [245] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:45.667 UTC [245] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:47.674 UTC [246] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:47.674 UTC [246] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:49.682 UTC [255] FATAL:  password authentication failed for user "admin"
go-service      | 2025/12/04 11:01:38 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
tts-db          | 2025-12-04 11:02:49.682 UTC [255] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:51.690 UTC [256] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:51.690 UTC [256] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:53.696 UTC [257] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:53.696 UTC [257] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:55.704 UTC [258] FATAL:  password authentication failed for user "admin"
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:01:40 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:40 /app/main.go:76
tts-db          | 2025-12-04 11:02:55.704 UTC [258] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:57.712 UTC [259] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:57.712 UTC [259] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:02:59.720 UTC [267] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:02:59.720 UTC [267] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:02.071 UTC [268] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:02.071 UTC [268] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:04.079 UTC [270] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:04.079 UTC [270] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:06.088 UTC [271] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:06.088 UTC [271] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:08.096 UTC [272] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:08.096 UTC [272] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:10.104 UTC [280] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:10.104 UTC [280] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:12.114 UTC [281] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:12.114 UTC [281] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:14.122 UTC [282] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:14.122 UTC [282] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:16.127 UTC [283] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:16.127 UTC [283] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:18.135 UTC [285] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:18.135 UTC [285] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:20.143 UTC [293] FATAL:  password authentication failed for user "admin"
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:42 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:44 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:46 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
tts-db          | 2025-12-04 11:03:20.143 UTC [293] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:22.479 UTC [294] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:22.479 UTC [294] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:24.487 UTC [295] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:24.487 UTC [295] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:26.494 UTC [296] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:26.494 UTC [296] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:28.504 UTC [297] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:28.504 UTC [297] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:30.509 UTC [305] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:30.509 UTC [305] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:32.518 UTC [306] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:32.518 UTC [306] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:34.525 UTC [308] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:34.525 UTC [308] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:36.534 UTC [309] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:36.534 UTC [309] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:38.543 UTC [310] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:38.543 UTC [310] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:40.551 UTC [318] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:40.551 UTC [318] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:42.892 UTC [319] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:42.892 UTC [319] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:44.901 UTC [320] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:44.901 UTC [320] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:46.908 UTC [321] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:46.908 UTC [321] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:48.917 UTC [323] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:48.917 UTC [323] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:50.924 UTC [331] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:50.924 UTC [331] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:52.930 UTC [332] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:52.930 UTC [332] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:54.936 UTC [333] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:54.936 UTC [333] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:56.945 UTC [334] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:56.945 UTC [334] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:03:58.953 UTC [335] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:03:58.953 UTC [335] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:00.962 UTC [343] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:00.962 UTC [343] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:03.298 UTC [345] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:03.298 UTC [345] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:05.306 UTC [346] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:05.306 UTC [346] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:07.314 UTC [347] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:07.314 UTC [347] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:09.321 UTC [348] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:09.321 UTC [348] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:11.329 UTC [356] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:11.329 UTC [356] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:13.337 UTC [357] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:13.337 UTC [357] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:15.347 UTC [358] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:15.347 UTC [358] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:17.355 UTC [359] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:17.355 UTC [359] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:19.363 UTC [361] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:19.363 UTC [361] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:21.371 UTC [369] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:21.371 UTC [369] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:23.716 UTC [370] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:23.716 UTC [370] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:25.725 UTC [371] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:25.725 UTC [371] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:27.734 UTC [372] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:27.734 UTC [372] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:29.742 UTC [378] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:29.742 UTC [378] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:31.751 UTC [380] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:31.751 UTC [380] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:33.759 UTC [382] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:33.759 UTC [382] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:35.766 UTC [383] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:35.766 UTC [383] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:37.774 UTC [384] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:37.774 UTC [384] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:39.783 UTC [385] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:39.783 UTC [385] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:41.791 UTC [393] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:41.791 UTC [393] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:44.145 UTC [394] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:44.145 UTC [394] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:46.153 UTC [395] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:46.153 UTC [395] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:47.804 UTC [27] LOG:  checkpoint starting: time
tts-db          | 2025-12-04 11:04:47.816 UTC [27] LOG:  checkpoint complete: wrote 3 buffers (0.0%); 0 WAL file(s) added, 0 removed, 0 recycled; write=0.003 s, sync=0.002 s, total=0.012 s; sync files=2, longest=0.001 s, average=0.001 s; distance=0 kB, estimate=0 kB
tts-db          | 2025-12-04 11:04:48.161 UTC [397] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:48.161 UTC [397] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:50.169 UTC [405] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:50.169 UTC [405] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:52.175 UTC [406] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:52.175 UTC [406] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:54.184 UTC [407] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:54.184 UTC [407] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:56.193 UTC [408] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:56.193 UTC [408] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:04:58.201 UTC [409] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:04:58.201 UTC [409] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:00.209 UTC [417] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:00.209 UTC [417] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:02.218 UTC [418] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:02.218 UTC [418] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:04.544 UTC [420] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:04.544 UTC [420] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:06.552 UTC [421] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:06.552 UTC [421] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:08.561 UTC [422] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:08.561 UTC [422] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:10.568 UTC [430] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:10.568 UTC [430] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:12.575 UTC [431] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:12.575 UTC [431] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:14.584 UTC [432] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:14.584 UTC [432] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
tts-db          | 2025-12-04 11:05:16.591 UTC [433] FATAL:  password authentication failed for user "admin"
tts-db          | 2025-12-04 11:05:16.591 UTC [433] DETAIL:  Connection matched pg_hba.conf line 100: "host all all all scram-sha-256"
go-service      | 2025/12/04 11:01:48 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:50 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:52 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:54 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:56 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:01:58 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:02:00 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:00 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:02 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:04 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:06 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:08 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:10 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:12 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:14 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:16 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:18 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:02:20 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:21 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:23 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:25 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:27 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:29 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:31 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:33 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:35 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:37 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:39 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:02:41 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:41 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:43 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:45 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:47 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:49 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:51 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:53 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:55 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:57 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:02:59 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:03:01 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:02 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:04 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:06 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:08 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:10 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:12 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:14 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:16 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:18 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:20 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:03:22 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:22 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:24 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:26 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:28 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:30 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:32 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:34 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:36 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:38 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:40 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:03:42 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:42 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:44 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:46 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:48 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:50 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:52 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:54 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:56 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:03:58 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:00 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:04:02 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:03 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:05 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:07 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:09 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:11 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:13 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:15 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:17 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:19 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:21 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:04:23 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:23 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:25 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:27 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:29 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:31 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:33 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:35 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:37 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:39 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:41 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:04:43 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:44 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:46 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:48 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:50 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:52 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:54 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:56 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:04:58 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:00 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:02 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | 2025/12/04 11:05:04 Failed to connect to database:failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:04 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:06 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:08 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:10 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:12 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:14 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      |
go-service      | 2025/12/04 11:05:16 /app/main.go:76
go-service      | [error] failed to initialize database, got error failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
go-service      | Waiting for Database... failed to connect to `user=admin database=tts_db`: 172.18.0.3:5432 (postgres): failed SASL auth: FATAL: password authentication failed for user "admin" (SQLSTATE 28P01)
root@Automations:/opt/TTS#
