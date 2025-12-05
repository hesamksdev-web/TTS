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