"""
Monkey-patch torchaudio.load to use soundfile instead of torchcodec.
Also patch TTS phoneme dataset to provide default cache path.
This must be imported before any TTS code.
"""
import os
import soundfile as sf
import torch


def load_with_soundfile(filepath, frame_offset=0, num_frames=-1, **kwargs):
    """
    Load audio using soundfile instead of torchaudio's torchcodec backend.
    Returns (waveform, sample_rate) to match torchaudio.load signature.
    """
    # Read audio file with soundfile
    data, sr = sf.read(filepath, dtype='float32')
    
    # Handle frame offset and num_frames
    if frame_offset > 0 or num_frames >= 0:
        start = frame_offset
        end = frame_offset + num_frames if num_frames >= 0 else len(data)
        data = data[start:end]
    
    # Convert to torch tensor and reshape to (channels, samples)
    # soundfile returns (samples, channels) for multi-channel or (samples,) for mono
    if len(data.shape) == 1:
        # Mono: reshape to (1, samples)
        waveform = torch.from_numpy(data).unsqueeze(0)
    else:
        # Multi-channel: transpose to (channels, samples)
        waveform = torch.from_numpy(data).t()
    
    return waveform, sr


def patch_torchaudio():
    """Apply the soundfile patch to torchaudio.load"""
    import torchaudio
    torchaudio.load = load_with_soundfile


def patch_phoneme_dataset():
    """Patch TTS phoneme dataset to provide default cache path if None"""
    try:
        from TTS.tts.datasets.dataset import PhonemeDataset
        original_init = PhonemeDataset.__init__
        
        def patched_init(self, *args, **kwargs):
            original_init(self, *args, **kwargs)
            # If cache_path is None, set it to a temp directory
            if self.cache_path is None:
                self.cache_path = os.path.join(os.path.expanduser("~"), ".cache", "tts_phonemes")
                os.makedirs(self.cache_path, exist_ok=True)
        
        PhonemeDataset.__init__ = patched_init
    except Exception:
        # If patching fails, continue anyway
        pass


def patch_transforms():
    """Patch TTS transforms to handle empty tensors in rational_quadratic_spline"""
    try:
        from TTS.tts.layers.vits import transforms
        original_rational_quadratic_spline = transforms.rational_quadratic_spline
        
        def patched_rational_quadratic_spline(
            inputs,
            unnormalized_widths,
            unnormalized_heights,
            unnormalized_derivatives,
            inverse=False,
            left=0.0,
            right=1.0,
            bottom=0.0,
            top=1.0,
            min_bin_width=transforms.DEFAULT_MIN_BIN_WIDTH,
            min_bin_height=transforms.DEFAULT_MIN_BIN_HEIGHT,
            min_derivative=transforms.DEFAULT_MIN_DERIVATIVE,
        ):
            # Handle empty tensor case - return empty outputs and logabsdet
            if inputs.numel() == 0:
                return inputs, torch.zeros_like(inputs)
            
            # Original validation check with proper dimension handling
            if inputs.numel() > 0:
                if torch.min(inputs) < left or torch.max(inputs) > right:
                    raise ValueError("Input to a transform is not within its domain")
            
            # Call the rest of the original function
            return original_rational_quadratic_spline(
                inputs=inputs,
                unnormalized_widths=unnormalized_widths,
                unnormalized_heights=unnormalized_heights,
                unnormalized_derivatives=unnormalized_derivatives,
                inverse=inverse,
                left=left,
                right=right,
                bottom=bottom,
                top=top,
                min_bin_width=min_bin_width,
                min_bin_height=min_bin_height,
                min_derivative=min_derivative,
            )
        
        transforms.rational_quadratic_spline = patched_rational_quadratic_spline
    except Exception:
        # If patching fails, continue anyway
        pass
