root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml down && docker-compose -f docker-compose.prod.yml up -d
[+] Running 6/6
 ✔ Container caddy-proxy     Removed                                                                                                                                            0.3s
 ✔ Container tts-frontend    Removed                                                                                                                                            0.2s
 ✔ Container go-service      Removed                                                                                                                                            0.3s
 ✔ Container tts-db          Removed                                                                                                                                            0.2s
 ✔ Container python-service  Removed                                                                                                                                            0.0s
 ✔ Network tts_tts-network   Removed                                                                                                                                            0.1s
[+] Running 6/6
 ✔ Network tts_tts-network   Created                                                                                                                                            0.0s
 ✔ Container python-service  Started                                                                                                                                            0.4s
 ✔ Container tts-db          Healthy                                                                                                                                           10.9s
 ✔ Container go-service      Started                                                                                                                                           11.1s
 ✔ Container tts-frontend    Started                                                                                                                                           11.2s
 ✔ Container caddy-proxy     Started                                                                                                                                           11.3s
root@Automations:/opt/TTS# docker-compose -f docker-compose.prod.yml logs
caddy-proxy  | {"level":"info","ts":1764947982.368705,"msg":"maxprocs: Leaving GOMAXPROCS=4: CPU quota undefined"}
caddy-proxy  | {"level":"info","ts":1764947982.369333,"msg":"GOMEMLIMIT is updated","package":"github.com/KimMachineGun/automemlimit/memlimit","GOMEMLIMIT":7314967756,"previous":9223372036854775807}
caddy-proxy  | {"level":"info","ts":1764947982.3696702,"msg":"using config from file","file":"/etc/caddy/Caddyfile"}
caddy-proxy  | {"level":"info","ts":1764947982.3722565,"msg":"adapted config to JSON","adapter":"caddyfile"}
caddy-proxy  | {"level":"warn","ts":1764947982.3723595,"msg":"Caddyfile input is not formatted; run 'caddy fmt --overwrite' to fix inconsistencies","adapter":"caddyfile","file":"/etc/caddy/Caddyfile","line":3}
caddy-proxy  | {"level":"warn","ts":1764947982.3728707,"logger":"admin","msg":"admin endpoint disabled"}
caddy-proxy  | {"level":"info","ts":1764947982.3766334,"logger":"http.auto_https","msg":"automatic HTTPS is completely disabled for server","server_name":"srv0"}
caddy-proxy  | {"level":"info","ts":1764947982.376884,"logger":"tls.cache.maintenance","msg":"started background certificate maintenance","cache":"0xc00056c500"}
caddy-proxy  | {"level":"warn","ts":1764947982.3785892,"logger":"http","msg":"HTTP/2 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy   | {"level":"warn","ts":1764947982.378835,"logger":"http","msg":"HTTP/3 skipped because it requires TLS","network":"tcp","addr":":80"}
caddy-proxy   | {"level":"info","ts":1764947982.3789039,"logger":"http.log","msg":"server running","name":"srv0","protocols":["h1","h2","h3"]}
caddy-proxy   | {"level":"info","ts":1764947982.3795302,"msg":"autosaved config (load with --resume flag)","file":"/config/caddy/autosave.json"}
tts-frontend  |
tts-frontend  | > frontend@0.1.0 start
tts-frontend  | > next start
tts-frontend  |
tts-frontend  |    ▲ Next.js 16.0.6
tts-frontend  |    - Local:         http://localhost:3000
tts-frontend  |    - Network:       http://172.18.0.5:3000
tts-frontend  |
tts-frontend  |  ✓ Starting...
tts-frontend  |  ✓ Ready in 506ms
caddy-proxy   | {"level":"info","ts":1764947982.3797007,"msg":"serving initial configuration"}
caddy-proxy   | {"level":"info","ts":1764947982.383792,"logger":"tls","msg":"storage cleaning happened too recently; skipping for now","storage":"FileStorage:/data/caddy","instance":"d8065e38-ac18-4c56-9718-355cfc33b1f3","try_again":1765034382.3837879,"try_again_in":86399.999999318}
go-service    | 2025/12/05 15:19:41 Go Service (SaaS Edition) listening on :8080
go-service    | 2025/12/05 15:20:10 python service request failed: Post "http://python-service:5000/voice-clone": dial tcp: lookup python-service on 127.0.0.11:53: server misbehaving
caddy-proxy   | {"level":"info","ts":1764947982.3840804,"logger":"tls","msg":"finished cleaning storage units"}
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
tts-db          |
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
tts-db          | PostgreSQL Database directory appears to contain a database; Skipping initialization
tts-db          |
tts-db          | 2025-12-05 15:19:31.109 UTC [1] LOG:  starting PostgreSQL 15.15 on x86_64-pc-linux-musl, compiled by gcc (Alpine 14.2.0) 14.2.0, 64-bit
tts-db          | 2025-12-05 15:19:31.110 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
tts-db          | 2025-12-05 15:19:31.110 UTC [1] LOG:  listening on IPv6 address "::", port 5432
tts-db          | 2025-12-05 15:19:31.112 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
tts-db          | 2025-12-05 15:19:31.117 UTC [29] LOG:  database system was shut down at 2025-12-05 15:19:30 UTC
tts-db          | 2025-12-05 15:19:31.124 UTC [1] LOG:  database system is ready to accept connections
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
python-service  | /usr/local/lib/python3.10/site-packages/scipy/__init__.py:132: UserWarning: A NumPy version >=1.21.6 and <1.28.0 is required for this version of SciPy (detected version 2.2.6)
python-service  |   warnings.warn(f"A NumPy version >={np_minversion} and <{np_maxversion}"
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 37, in <module>
python-service  |     patch_phoneme_dataset()
python-service  |   File "/app/app/audio_patch.py", line 46, in patch_phoneme_dataset
python-service  |     from TTS.tts.datasets.dataset import PhonemeDataset
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/__init__.py", line 9, in <module>
python-service  |     from TTS.tts.datasets.dataset import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/tts/datasets/dataset.py", line 13, in <module>
python-service  |     from TTS.utils.audio import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  |
python-service  | A module that was compiled using NumPy 1.x cannot be run in
python-service  | NumPy 2.2.6 as it may crash. To support both 1.x and 2.x
python-service  | versions of NumPy, modules must be compiled with NumPy 2.0.
python-service  | Some module may need to rebuild instead e.g. with 'pybind11>=2.12'.
python-service  |
python-service  | If you are a user of the module, the easiest solution will be to
python-service  | downgrade to 'numpy<2' or try to upgrade the affected module.
python-service  | We expect that some modules will need time to support NumPy 2.
python-service  |
python-service  | Traceback (most recent call last):  File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | AttributeError: _ARRAY_API not found
python-service  | Traceback (most recent call last):
python-service  |   File "/usr/local/bin/uvicorn", line 7, in <module>
python-service  |     sys.exit(main())
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1485, in __call__
python-service  |     return self.main(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1406, in main
python-service  |     rv = self.invoke(ctx)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 1269, in invoke
python-service  |     return ctx.invoke(self.callback, **ctx.params)
python-service  |   File "/usr/local/lib/python3.10/site-packages/click/core.py", line 824, in invoke
python-service  |     return callback(*args, **kwargs)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 410, in main
python-service  |     run(
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/main.py", line 577, in run
python-service  |     server.run()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 65, in run
python-service  |     return asyncio.run(self.serve(sockets=sockets))
python-service  |   File "/usr/local/lib/python3.10/asyncio/runners.py", line 44, in run
python-service  |     return loop.run_until_complete(main)
python-service  |   File "uvloop/loop.pyx", line 1518, in uvloop.loop.Loop.run_until_complete
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 69, in serve
python-service  |     await self._serve(sockets)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/server.py", line 76, in _serve
python-service  |     config.load()
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/config.py", line 434, in load
python-service  |     self.loaded_app = import_from_string(self.app)
python-service  |   File "/usr/local/lib/python3.10/site-packages/uvicorn/importer.py", line 19, in import_from_string
python-service  |     module = importlib.import_module(module_str)
python-service  |   File "/usr/local/lib/python3.10/importlib/__init__.py", line 126, in import_module
python-service  |     return _bootstrap._gcd_import(name[level:], package, level)
python-service  |   File "<frozen importlib._bootstrap>", line 1050, in _gcd_import
python-service  |   File "<frozen importlib._bootstrap>", line 1027, in _find_and_load
python-service  |   File "<frozen importlib._bootstrap>", line 1006, in _find_and_load_unlocked
python-service  |   File "<frozen importlib._bootstrap>", line 688, in _load_unlocked
python-service  |   File "<frozen importlib._bootstrap_external>", line 883, in exec_module
python-service  |   File "<frozen importlib._bootstrap>", line 241, in _call_with_frames_removed
python-service  |   File "/app/app/main.py", line 43, in <module>
python-service  |     from TTS.api import TTS
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/api.py", line 9, in <module>
python-service  |     from TTS.utils.audio.numpy_transforms import save_wav
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/__init__.py", line 1, in <module>
python-service  |     from TTS.utils.audio.processor import AudioProcessor
python-service  |   File "/usr/local/lib/python3.10/site-packages/TTS/utils/audio/processor.py", line 6, in <module>
python-service  |     import scipy.io.wavfile
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/__init__.py", line 97, in <module>
python-service  |     from .matlab import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/__init__.py", line 45, in <module>
python-service  |     from ._mio import loadmat, savemat, whosmat
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio.py", line 9, in <module>
python-service  |     from ._mio4 import MatFile4Reader, MatFile4Writer
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/io/matlab/_mio4.py", line 8, in <module>
python-service  |     import scipy.sparse
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/__init__.py", line 274, in <module>
python-service  |     from ._csr import *
python-service  |   File "/usr/local/lib/python3.10/site-packages/scipy/sparse/_csr.py", line 11, in <module>
python-service  |     from ._sparsetools import (csr_tocsc, csr_tobsr, csr_count_blocks,
python-service  | ImportError: numpy.core.multiarray failed to import
root@Automations:/opt/TTS#