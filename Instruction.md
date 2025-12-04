root@Automations:/opt/TTS# cat training_runs/my_dataset_20251204145511/train.log
/usr/local/lib/python3.10/site-packages/librosa/core/intervals.py:8: UserWarning: pkg_resources is deprecated as an API. See https://setuptools.pypa.io/en/latest/pkg_resources.html. The pkg_resources package is slated for removal as early as 2025-11-30. Refrain from using this package or pin to Setuptools<81.
  from pkg_resources import resource_filename
 | > Found 1 files in /app/data/my_dataset
 > Using model: vits
 > Setting up Audio Processor...
 | > sample_rate:22050
 | > resample:False
 | > num_mels:80
 | > log_func:np.log10
 | > min_level_db:0
 | > frame_shift_ms:None
 | > frame_length_ms:None
 | > ref_level_db:None
 | > fft_size:1024
 | > power:None
 | > preemphasis:0.0
 | > griffin_lim_iters:None
 | > signal_norm:None
 | > symmetric_norm:None
 | > mel_fmin:0
 | > mel_fmax:None
 | > pitch_fmin:None
 | > pitch_fmax:None
 | > spec_gain:20.0
 | > stft_pad_mode:reflect
 | > max_norm:1.0
 | > clip_norm:True
 | > do_trim_silence:False
 | > trim_db:60
 | > do_sound_norm:False
 | > do_amp_to_db_linear:True
 | > do_amp_to_db_mel:True
 | > do_rms_norm:False
 | > db_level:None
 | > stats_path:None
 | > base:10
 | > hop_length:256
 | > win_length:1024
 > Training Environment:
 | > Backend: Torch
 | > Mixed precision: False
 | > Precision: float32
 | > Num. of CPUs: 4
 | > Num. of Torch Threads: 4
 | > Torch seed: 54321
 | > Torch CUDNN: True
 | > Torch CUDNN deterministic: False
 | > Torch CUDNN benchmark: False
 | > Torch TF32 MatMul: False
 > Start Tensorboard: tensorboard --logdir=/app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000
 > Restoring from model_file.pth ...
 > Restoring Model...
 > Partial model initialization...
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.0.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.0.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.1.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.1.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.2.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.2.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.3.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.3.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.4.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.4.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.5.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.5.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.6.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.6.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.7.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.7.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.8.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.8.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.9.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.9.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.10.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.10.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.11.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.11.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.12.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.12.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.13.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.13.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.14.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.14.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.15.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.in_layers.15.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.0.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.0.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.1.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.1.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.2.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.2.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.3.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.3.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.4.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.4.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.5.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.5.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.6.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.6.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.7.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.7.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.8.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.8.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.9.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.9.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.10.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.10.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.11.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.11.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.12.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.12.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.13.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.13.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.14.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.14.weight_v
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.15.weight_g
 | > Layer missing in the model definition: posterior_encoder.enc.res_skip_layers.15.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.in_layers.3.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.0.enc.res_skip_layers.3.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.in_layers.3.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.1.enc.res_skip_layers.3.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.in_layers.3.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.2.enc.res_skip_layers.3.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.in_layers.3.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.0.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.0.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.1.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.1.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.2.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.2.weight_v
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.3.weight_g
 | > Layer missing in the model definition: flow.flows.3.enc.res_skip_layers.3.weight_v
 | > Layer missing in the model definition: waveform_decoder.ups.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.ups.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.ups.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.ups.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.ups.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.ups.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.ups.3.weight_g
 | > Layer missing in the model definition: waveform_decoder.ups.3.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.0.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.1.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.2.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.3.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.4.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.5.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.6.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.7.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.8.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.9.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.10.convs2.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs1.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs1.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs1.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs1.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs1.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs1.2.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs2.0.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs2.0.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs2.1.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs2.1.weight_v
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs2.2.weight_g
 | > Layer missing in the model definition: waveform_decoder.resblocks.11.convs2.2.weight_v
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.5.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.5.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.6.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.6.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.7.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.7.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.8.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.8.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.9.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.9.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.10.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.10.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.11.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.11.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.12.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.12.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.13.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.13.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.14.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.14.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.15.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.in_layers.15.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.5.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.5.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.6.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.6.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.7.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.7.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.8.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.8.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.9.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.9.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.10.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.10.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.11.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.11.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.12.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.12.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.13.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.13.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.14.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.14.parametrizations.weight.original1
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.15.parametrizations.weight.original0
 | > Layer missing in the checkpoint: posterior_encoder.enc.res_skip_layers.15.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.in_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.0.enc.res_skip_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.in_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.1.enc.res_skip_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.in_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.2.enc.res_skip_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.in_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: flow.flows.3.enc.res_skip_layers.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.ups.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.ups.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.ups.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.ups.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.ups.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.ups.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.ups.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.ups.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.0.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.1.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.2.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.3.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.4.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.5.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.6.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.7.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.8.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.9.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.10.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs1.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs1.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs1.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs1.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs1.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs1.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs2.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs2.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs2.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs2.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs2.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: waveform_decoder.resblocks.11.convs2.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.0.convs.0.bias
 | > Layer missing in the checkpoint: disc.nets.0.convs.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.0.convs.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.0.convs.1.bias
 | > Layer missing in the checkpoint: disc.nets.0.convs.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.0.convs.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.0.convs.2.bias
 | > Layer missing in the checkpoint: disc.nets.0.convs.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.0.convs.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.0.convs.3.bias
 | > Layer missing in the checkpoint: disc.nets.0.convs.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.0.convs.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.0.convs.4.bias
 | > Layer missing in the checkpoint: disc.nets.0.convs.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.0.convs.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.0.convs.5.bias
 | > Layer missing in the checkpoint: disc.nets.0.convs.5.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.0.convs.5.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.0.conv_post.bias
 | > Layer missing in the checkpoint: disc.nets.0.conv_post.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.0.conv_post.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.1.convs.0.bias
 | > Layer missing in the checkpoint: disc.nets.1.convs.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.1.convs.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.1.convs.1.bias
 | > Layer missing in the checkpoint: disc.nets.1.convs.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.1.convs.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.1.convs.2.bias
 | > Layer missing in the checkpoint: disc.nets.1.convs.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.1.convs.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.1.convs.3.bias
 | > Layer missing in the checkpoint: disc.nets.1.convs.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.1.convs.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.1.convs.4.bias
 | > Layer missing in the checkpoint: disc.nets.1.convs.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.1.convs.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.1.conv_post.bias
 | > Layer missing in the checkpoint: disc.nets.1.conv_post.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.1.conv_post.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.2.convs.0.bias
 | > Layer missing in the checkpoint: disc.nets.2.convs.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.2.convs.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.2.convs.1.bias
 | > Layer missing in the checkpoint: disc.nets.2.convs.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.2.convs.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.2.convs.2.bias
 | > Layer missing in the checkpoint: disc.nets.2.convs.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.2.convs.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.2.convs.3.bias
 | > Layer missing in the checkpoint: disc.nets.2.convs.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.2.convs.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.2.convs.4.bias
 | > Layer missing in the checkpoint: disc.nets.2.convs.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.2.convs.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.2.conv_post.bias
 | > Layer missing in the checkpoint: disc.nets.2.conv_post.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.2.conv_post.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.3.convs.0.bias
 | > Layer missing in the checkpoint: disc.nets.3.convs.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.3.convs.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.3.convs.1.bias
 | > Layer missing in the checkpoint: disc.nets.3.convs.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.3.convs.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.3.convs.2.bias
 | > Layer missing in the checkpoint: disc.nets.3.convs.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.3.convs.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.3.convs.3.bias
 | > Layer missing in the checkpoint: disc.nets.3.convs.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.3.convs.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.3.convs.4.bias
 | > Layer missing in the checkpoint: disc.nets.3.convs.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.3.convs.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.3.conv_post.bias
 | > Layer missing in the checkpoint: disc.nets.3.conv_post.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.3.conv_post.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.4.convs.0.bias
 | > Layer missing in the checkpoint: disc.nets.4.convs.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.4.convs.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.4.convs.1.bias
 | > Layer missing in the checkpoint: disc.nets.4.convs.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.4.convs.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.4.convs.2.bias
 | > Layer missing in the checkpoint: disc.nets.4.convs.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.4.convs.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.4.convs.3.bias
 | > Layer missing in the checkpoint: disc.nets.4.convs.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.4.convs.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.4.convs.4.bias
 | > Layer missing in the checkpoint: disc.nets.4.convs.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.4.convs.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.4.conv_post.bias
 | > Layer missing in the checkpoint: disc.nets.4.conv_post.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.4.conv_post.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.5.convs.0.bias
 | > Layer missing in the checkpoint: disc.nets.5.convs.0.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.5.convs.0.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.5.convs.1.bias
 | > Layer missing in the checkpoint: disc.nets.5.convs.1.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.5.convs.1.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.5.convs.2.bias
 | > Layer missing in the checkpoint: disc.nets.5.convs.2.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.5.convs.2.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.5.convs.3.bias
 | > Layer missing in the checkpoint: disc.nets.5.convs.3.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.5.convs.3.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.5.convs.4.bias
 | > Layer missing in the checkpoint: disc.nets.5.convs.4.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.5.convs.4.parametrizations.weight.original1
 | > Layer missing in the checkpoint: disc.nets.5.conv_post.bias
 | > Layer missing in the checkpoint: disc.nets.5.conv_post.parametrizations.weight.original0
 | > Layer missing in the checkpoint: disc.nets.5.conv_post.parametrizations.weight.original1
 | > 558 / 949 layers are restored.
 > Model restored from step 1000000
/usr/local/lib/python3.10/site-packages/trainer/trainer.py:561: FutureWarning: `torch.cuda.amp.GradScaler(args...)` is deprecated. Please use `torch.amp.GradScaler('cuda', args...)` instead.
  self.scaler = torch.cuda.amp.GradScaler()
/usr/local/lib/python3.10/site-packages/torch/cuda/amp/grad_scaler.py:31: UserWarning: torch.cuda.amp.GradScaler is enabled, but CUDA is not available.  Disabling.
  super().__init__(

 > Model has 83068396 parameters

 > EPOCH: 0/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000


> DataLoader initialization
| > Tokenizer:
        | > add_blank: True
        | > use_eos_bos: False
        | > use_phonemes: True
        | > phonemizer:
                | > phoneme language: en-us
                | > phoneme backend: espeak
| > Number of instances : 1
 | > Preprocessing samples
 | > Max text length: 72
 | > Min text length: 72
 | > Avg text length: 72.0
 |
 | > Max audio length: 101991.0
 | > Min audio length: 101991.0
 | > Avg audio length: 101991.0
 | > Num. instances discarded samples: 0
 | > Batch group size: 0.

 > TRAINING (2025-12-04 14:55:26)
/usr/local/lib/python3.10/site-packages/torch/functional.py:681: UserWarning: stft with return_complex=False is deprecated. In a future pytorch release, stft will return complex tensors for all inputs, and return_complex=False will raise an error.
Note: you can still call torch.view_as_real on the complex output to recover the old return format. (Triggered internally at /pytorch/aten/src/ATen/native/SpectralOps.cpp:875.)
  return _VF.stft(  # type: ignore[attr-defined]
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion


> DataLoader initialization
| > Tokenizer:
        | > add_blank: True
        | > use_eos_bos: False
        | > use_phonemes: True
        | > phonemizer:
                | > phoneme language: en-us
                | > phoneme backend: espeak
| > Number of instances : 1
 | > Preprocessing samples
 | > Max text length: 72
 | > Min text length: 72
 | > Avg text length: 72.0
 |
 | > Max audio length: 101991.0
 | > Min audio length: 101991.0
 | > Avg audio length: 101991.0
 | > Num. instances discarded samples: 0
 | > Batch group size: 0.

 > EVALUATION

 | > Synthesizing test sentences.
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1455: UserWarning: The use of `x.T` on tensors of dimension other than 2 to reverse their shape is deprecated and it will throw an error in a future release. Consider `x.mT` to transpose batches of matrices or `x.permute(*torch.arange(x.ndim - 1, -1, -1))` to reverse the dimensions of a tensor. (Triggered internally at /pytorch/aten/src/ATen/native/TensorShape.cpp:4416.)
  test_figures["{}-alignment".format(idx)] = plot_alignment(alignment.T, output_fig=False)

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.006119728088378906 (+0)
     | > avg_loss_disc: 4.5814948081970215 (+0)
     | > avg_loss_disc_real_0: 0.62183678150177 (+0)
     | > avg_loss_disc_real_1: 0.7922453880310059 (+0)
     | > avg_loss_disc_real_2: 0.7934720516204834 (+0)
     | > avg_loss_disc_real_3: 0.7719997763633728 (+0)
     | > avg_loss_disc_real_4: 0.7112516164779663 (+0)
     | > avg_loss_disc_real_5: 0.7622858285903931 (+0)
     | > avg_loss_0: 4.5814948081970215 (+0)
     | > avg_loss_gen: 4.453646659851074 (+0)
     | > avg_loss_kl: 21.429943084716797 (+0)
     | > avg_loss_feat: 0.17062613368034363 (+0)
     | > avg_loss_mel: 30.747154235839844 (+0)
     | > avg_loss_duration: 1.8812894821166992 (+0)
     | > avg_loss_1: 58.68266296386719 (+0)

 > BEST MODEL : /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000/best_model_1000002.pth

 > EPOCH: 1/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:55:59)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.00559544563293457 (-0.0005242824554443359)
     | > avg_loss_disc: 3.410338878631592 (-1.1711559295654297)
     | > avg_loss_disc_real_0: 0.34364357590675354 (-0.2781932055950165)
     | > avg_loss_disc_real_1: 0.5103206634521484 (-0.2819247245788574)
     | > avg_loss_disc_real_2: 0.5063145160675049 (-0.2871575355529785)
     | > avg_loss_disc_real_3: 0.4664575457572937 (-0.3055422306060791)
     | > avg_loss_disc_real_4: 0.39988645911216736 (-0.31136515736579895)
     | > avg_loss_disc_real_5: 0.4908544719219208 (-0.2714313566684723)
     | > avg_loss_0: 3.410338878631592 (-1.1711559295654297)
     | > avg_loss_gen: 2.724057197570801 (-1.7295894622802734)
     | > avg_loss_kl: 13.253870964050293 (-8.176072120666504)
     | > avg_loss_feat: 0.40507030487060547 (+0.23444417119026184)
     | > avg_loss_mel: 34.39110565185547 (+3.643951416015625)
     | > avg_loss_duration: 1.5965768098831177 (-0.28471267223358154)
     | > avg_loss_1: 52.37068176269531 (-6.311981201171875)

 > BEST MODEL : /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000/best_model_1000003.pth

 > EPOCH: 2/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:56:32)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.0498805046081543 (+0.04428505897521973)
     | > avg_loss_disc: 3.4558277130126953 (+0.045488834381103516)
     | > avg_loss_disc_real_0: 0.1574803739786148 (-0.18616320192813873)
     | > avg_loss_disc_real_1: 0.16291925311088562 (-0.3474014103412628)
     | > avg_loss_disc_real_2: 0.1653490960597992 (-0.3409654200077057)
     | > avg_loss_disc_real_3: 0.12681680917739868 (-0.339640736579895)
     | > avg_loss_disc_real_4: 0.09213787317276001 (-0.30774858593940735)
     | > avg_loss_disc_real_5: 0.1723603904247284 (-0.3184940814971924)
     | > avg_loss_0: 3.4558277130126953 (+0.045488834381103516)
     | > avg_loss_gen: 0.8771001100540161 (-1.8469570875167847)
     | > avg_loss_kl: 8.99279499053955 (-4.261075973510742)
     | > avg_loss_feat: 0.20059524476528168 (-0.2044750601053238)
     | > avg_loss_mel: 25.032686233520508 (-9.358419418334961)
     | > avg_loss_duration: 1.6070560216903687 (+0.010479211807250977)
     | > avg_loss_1: 36.71023178100586 (-15.660449981689453)

 > BEST MODEL : /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000/best_model_1000004.pth

 > EPOCH: 3/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:57:05)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.004311561584472656 (-0.04556894302368164)
     | > avg_loss_disc: 3.7104458808898926 (+0.25461816787719727)
     | > avg_loss_disc_real_0: 0.10241949558258057 (-0.05506087839603424)
     | > avg_loss_disc_real_1: 0.06858160346746445 (-0.09433764964342117)
     | > avg_loss_disc_real_2: 0.0756978914141655 (-0.0896512046456337)
     | > avg_loss_disc_real_3: 0.10995330661535263 (-0.01686350256204605)
     | > avg_loss_disc_real_4: 0.14562708139419556 (+0.05348920822143555)
     | > avg_loss_disc_real_5: 0.13080115616321564 (-0.041559234261512756)
     | > avg_loss_0: 3.7104458808898926 (+0.25461816787719727)
     | > avg_loss_gen: 0.636505126953125 (-0.2405949831008911)
     | > avg_loss_kl: 8.577470779418945 (-0.41532421112060547)
     | > avg_loss_feat: 0.37989386916160583 (+0.17929862439632416)
     | > avg_loss_mel: 28.022321701049805 (+2.989635467529297)
     | > avg_loss_duration: 1.6366750001907349 (+0.02961897850036621)
     | > avg_loss_1: 39.25286865234375 (+2.5426368713378906)


 > EPOCH: 4/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:57:32)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.008492231369018555 (+0.0041806697845458984)
     | > avg_loss_disc: 3.309513568878174 (-0.40093231201171875)
     | > avg_loss_disc_real_0: 0.10300500690937042 (+0.000585511326789856)
     | > avg_loss_disc_real_1: 0.12794607877731323 (+0.059364475309848785)
     | > avg_loss_disc_real_2: 0.13300541043281555 (+0.057307519018650055)
     | > avg_loss_disc_real_3: 0.19537030160427094 (+0.0854169949889183)
     | > avg_loss_disc_real_4: 0.2661876082420349 (+0.12056052684783936)
     | > avg_loss_disc_real_5: 0.1986822783946991 (+0.06788112223148346)
     | > avg_loss_0: 3.309513568878174 (-0.40093231201171875)
     | > avg_loss_gen: 1.0342037677764893 (+0.39769864082336426)
     | > avg_loss_kl: 6.739046573638916 (-1.8384242057800293)
     | > avg_loss_feat: 0.4253697693347931 (+0.045475900173187256)
     | > avg_loss_mel: 31.806814193725586 (+3.7844924926757812)
     | > avg_loss_duration: 1.7139593362808228 (+0.07728433609008789)
     | > avg_loss_1: 41.719390869140625 (+2.466522216796875)


 > EPOCH: 5/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:57:59)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.006452083587646484 (-0.0020401477813720703)
     | > avg_loss_disc: 3.1470208168029785 (-0.1624927520751953)
     | > avg_loss_disc_real_0: 0.15595649182796478 (+0.05295148491859436)
     | > avg_loss_disc_real_1: 0.29472264647483826 (+0.16677656769752502)
     | > avg_loss_disc_real_2: 0.2951781749725342 (+0.16217276453971863)
     | > avg_loss_disc_real_3: 0.3437584638595581 (+0.14838816225528717)
     | > avg_loss_disc_real_4: 0.37970733642578125 (+0.11351972818374634)
     | > avg_loss_disc_real_5: 0.34114667773246765 (+0.14246439933776855)
     | > avg_loss_0: 3.1470208168029785 (-0.1624927520751953)
     | > avg_loss_gen: 1.8226349353790283 (+0.7884311676025391)
     | > avg_loss_kl: 6.652072429656982 (-0.0869741439819336)
     | > avg_loss_feat: 0.3849424421787262 (-0.040427327156066895)
     | > avg_loss_mel: 27.028528213500977 (-4.778285980224609)
     | > avg_loss_duration: 1.6306160688400269 (-0.0833432674407959)
     | > avg_loss_1: 37.518795013427734 (-4.200595855712891)


 > EPOCH: 6/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:58:28)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.00570368766784668 (-0.0007483959197998047)
     | > avg_loss_disc: 3.23134708404541 (+0.08432626724243164)
     | > avg_loss_disc_real_0: 0.236568883061409 (+0.08061239123344421)
     | > avg_loss_disc_real_1: 0.40677300095558167 (+0.11205035448074341)
     | > avg_loss_disc_real_2: 0.411217600107193 (+0.11603942513465881)
     | > avg_loss_disc_real_3: 0.4081677794456482 (+0.06440931558609009)
     | > avg_loss_disc_real_4: 0.40221789479255676 (+0.022510558366775513)
     | > avg_loss_disc_real_5: 0.41955456137657166 (+0.078407883644104)
     | > avg_loss_0: 3.23134708404541 (+0.08432626724243164)
     | > avg_loss_gen: 2.298447370529175 (+0.4758124351501465)
     | > avg_loss_kl: 5.661383152008057 (-0.9906892776489258)
     | > avg_loss_feat: 0.4114537835121155 (+0.026511341333389282)
     | > avg_loss_mel: 34.57725143432617 (+7.548723220825195)
     | > avg_loss_duration: 1.6677111387252808 (+0.037095069885253906)
     | > avg_loss_1: 44.61624526977539 (+7.097450256347656)


 > EPOCH: 7/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:58:58)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.006796836853027344 (+0.001093149185180664)
     | > avg_loss_disc: 3.213514804840088 (-0.017832279205322266)
     | > avg_loss_disc_real_0: 0.31301257014274597 (+0.07644368708133698)
     | > avg_loss_disc_real_1: 0.4145524203777313 (+0.007779419422149658)
     | > avg_loss_disc_real_2: 0.4209626019001007 (+0.009745001792907715)
     | > avg_loss_disc_real_3: 0.388883501291275 (-0.01928427815437317)
     | > avg_loss_disc_real_4: 0.35483893752098083 (-0.04737895727157593)
     | > avg_loss_disc_real_5: 0.4048769176006317 (-0.014677643775939941)
     | > avg_loss_0: 3.213514804840088 (-0.017832279205322266)
     | > avg_loss_gen: 2.308960437774658 (+0.010513067245483398)
     | > avg_loss_kl: 6.33433723449707 (+0.6729540824890137)
     | > avg_loss_feat: 0.40203434228897095 (-0.009419441223144531)
     | > avg_loss_mel: 25.961559295654297 (-8.615692138671875)
     | > avg_loss_duration: 1.6798441410064697 (+0.012133002281188965)
     | > avg_loss_1: 36.68673324584961 (-7.929512023925781)

 > BEST MODEL : /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000/best_model_1000009.pth

 > EPOCH: 8/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 14:59:34)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.0078067779541015625 (+0.0010099411010742188)
     | > avg_loss_disc: 3.1148433685302734 (-0.09867143630981445)
     | > avg_loss_disc_real_0: 0.3623775243759155 (+0.049364954233169556)
     | > avg_loss_disc_real_1: 0.34351059794425964 (-0.07104182243347168)
     | > avg_loss_disc_real_2: 0.3581702709197998 (-0.0627923309803009)
     | > avg_loss_disc_real_3: 0.3101358711719513 (-0.07874763011932373)
     | > avg_loss_disc_real_4: 0.276631623506546 (-0.07820731401443481)
     | > avg_loss_disc_real_5: 0.3360292911529541 (-0.06884762644767761)
     | > avg_loss_0: 3.1148433685302734 (-0.09867143630981445)
     | > avg_loss_gen: 2.0128798484802246 (-0.2960805892944336)
     | > avg_loss_kl: 4.431258678436279 (-1.903078556060791)
     | > avg_loss_feat: 0.4861384928226471 (+0.08410415053367615)
     | > avg_loss_mel: 24.48224449157715 (-1.4793148040771484)
     | > avg_loss_duration: 1.6176711320877075 (-0.06217300891876221)
     | > avg_loss_1: 33.03019332885742 (-3.6565399169921875)

 > BEST MODEL : /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000/best_model_1000010.pth

 > EPOCH: 9/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 15:00:11)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.007678508758544922 (-0.00012826919555664062)
     | > avg_loss_disc: 3.106476306915283 (-0.008367061614990234)
     | > avg_loss_disc_real_0: 0.37629252672195435 (+0.013915002346038818)
     | > avg_loss_disc_real_1: 0.2580568790435791 (-0.08545371890068054)
     | > avg_loss_disc_real_2: 0.27399492263793945 (-0.08417534828186035)
     | > avg_loss_disc_real_3: 0.22784456610679626 (-0.08229130506515503)
     | > avg_loss_disc_real_4: 0.2229522317647934 (-0.053679391741752625)
     | > avg_loss_disc_real_5: 0.25885918736457825 (-0.07717010378837585)
     | > avg_loss_0: 3.106476306915283 (-0.008367061614990234)
     | > avg_loss_gen: 1.6291559934616089 (-0.3837238550186157)
     | > avg_loss_kl: 4.531954288482666 (+0.10069561004638672)
     | > avg_loss_feat: 0.1486668884754181 (-0.337471604347229)
     | > avg_loss_mel: 31.783254623413086 (+7.3010101318359375)
     | > avg_loss_duration: 1.610673427581787 (-0.00699770450592041)
     | > avg_loss_1: 39.70370101928711 (+6.6735076904296875)


 > EPOCH: 10/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 15:00:40)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.005582332611083984 (-0.0020961761474609375)
     | > avg_loss_disc: 3.1368932723999023 (+0.03041696548461914)
     | > avg_loss_disc_real_0: 0.3575332760810852 (-0.01875925064086914)
     | > avg_loss_disc_real_1: 0.17923328280448914 (-0.07882359623908997)
     | > avg_loss_disc_real_2: 0.1938566118478775 (-0.08013831079006195)
     | > avg_loss_disc_real_3: 0.20071308314800262 (-0.02713148295879364)
     | > avg_loss_disc_real_4: 0.2198261022567749 (-0.0031261295080184937)
     | > avg_loss_disc_real_5: 0.20886772871017456 (-0.049991458654403687)
     | > avg_loss_0: 3.1368932723999023 (+0.03041696548461914)
     | > avg_loss_gen: 1.3778800964355469 (-0.251275897026062)
     | > avg_loss_kl: 4.78898286819458 (+0.25702857971191406)
     | > avg_loss_feat: 0.42466890811920166 (+0.27600201964378357)
     | > avg_loss_mel: 35.32855987548828 (+3.5453052520751953)
     | > avg_loss_duration: 1.5563570261001587 (-0.05431640148162842)
     | > avg_loss_1: 43.4764518737793 (+3.7727508544921875)


 > EPOCH: 11/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 15:01:10)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.055733442306518555 (+0.05015110969543457)
     | > avg_loss_disc: 3.096923351287842 (-0.03996992111206055)
     | > avg_loss_disc_real_0: 0.3197193741798401 (-0.03781390190124512)
     | > avg_loss_disc_real_1: 0.1874074935913086 (+0.008174210786819458)
     | > avg_loss_disc_real_2: 0.1869002729654312 (-0.006956338882446289)
     | > avg_loss_disc_real_3: 0.2355770617723465 (+0.03486397862434387)
     | > avg_loss_disc_real_4: 0.2629163861274719 (+0.04309028387069702)
     | > avg_loss_disc_real_5: 0.22708649933338165 (+0.018218770623207092)
     | > avg_loss_0: 3.096923351287842 (-0.03996992111206055)
     | > avg_loss_gen: 1.4399088621139526 (+0.06202876567840576)
     | > avg_loss_kl: 4.7269110679626465 (-0.062071800231933594)
     | > avg_loss_feat: 0.3774986267089844 (-0.047170281410217285)
     | > avg_loss_mel: 26.827089309692383 (-8.501470565795898)
     | > avg_loss_duration: 1.4681248664855957 (-0.08823215961456299)
     | > avg_loss_1: 34.83953094482422 (-8.636920928955078)


 > EPOCH: 12/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 15:01:38)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.

  --> EVAL PERFORMANCE
     | > avg_loader_time: 0.005896568298339844 (-0.04983687400817871)
     | > avg_loss_disc: 3.0344018936157227 (-0.06252145767211914)
     | > avg_loss_disc_real_0: 0.27538806200027466 (-0.04433131217956543)
     | > avg_loss_disc_real_1: 0.2279210239648819 (+0.0405135303735733)
     | > avg_loss_disc_real_2: 0.2260938286781311 (+0.03919355571269989)
     | > avg_loss_disc_real_3: 0.28526681661605835 (+0.04968975484371185)
     | > avg_loss_disc_real_4: 0.30110281705856323 (+0.03818643093109131)
     | > avg_loss_disc_real_5: 0.27526846528053284 (+0.048181965947151184)
     | > avg_loss_0: 3.0344018936157227 (-0.06252145767211914)
     | > avg_loss_gen: 1.639206886291504 (+0.19929802417755127)
     | > avg_loss_kl: 3.890453338623047 (-0.8364577293395996)
     | > avg_loss_feat: 0.5137903690338135 (+0.1362917423248291)
     | > avg_loss_mel: 38.097023010253906 (+11.269933700561523)
     | > avg_loss_duration: 1.533678412437439 (+0.06555354595184326)
     | > avg_loss_1: 45.67415237426758 (+10.83462142944336)


 > EPOCH: 13/10000
 --> /app/training_runs/my_dataset_20251204145511/my_dataset_20251204145511-December-04-2025_02+55PM-0000000

 > TRAINING (2025-12-04 15:02:06)
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1273: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1284: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):
/usr/local/lib/python3.10/site-packages/TTS/tts/models/vits.py:1311: FutureWarning: `torch.cuda.amp.autocast(args...)` is deprecated. Please use `torch.amp.autocast('cuda', args...)` instead.
  with autocast(enabled=False):  # use float32 for the criterion

 > EVALUATION

 | > Synthesizing test sentences.
root@Automations:/opt/TTS#