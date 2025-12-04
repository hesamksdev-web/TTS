# TTS Model Training Guide

## Overview

The admin panel now includes a complete training system that allows you to fine-tune TTS models with your own voice data. This guide explains how to use the training feature to train a model with 102 MP3 files.

## Features

- **Batch Audio Upload**: Upload multiple MP3 files (or WAV/M4A) with their corresponding transcripts
- **Automatic Format Conversion**: MP3 files are automatically converted to WAV format (22050 Hz, mono)
- **Dataset Management**: Organize audio files into named datasets
- **Training Configuration**: Control epochs and batch size for training
- **Real-time Status**: Monitor training jobs as they progress
- **Admin-Only Access**: Training features are restricted to admin users

## Getting Started

### 1. Access the Admin Panel

1. Log in to the application with your admin account
2. Navigate to the Admin Dashboard
3. Click on the "Training" tab

### 2. Prepare Your Data

Before uploading, organize your data:

- **Audio Files**: MP3, WAV, or M4A format (MP3 recommended for smaller file sizes)
- **Transcripts**: Plain text of what is spoken in each audio file
- **Dataset Name**: A unique name for your training dataset (e.g., `my_voice_dataset`)

### 3. Upload Training Data

#### Step-by-Step Upload Process

1. **Set Dataset Name**
   - Enter a unique name in the "Dataset Name" field
   - This name will be used to organize all your audio files
   - Example: `my_voice_dataset`

2. **Add Audio Files**
   - Click "Add Audio File" button to add a new file entry
   - You can add multiple files at once
   - For 102 files, click the button 102 times (or add them in batches)

3. **For Each Audio File**
   - Click the file input to select your MP3 file
   - Enter the exact transcript (what is being said in the audio)
   - Click "Upload" button
   - Wait for the success message before moving to the next file

4. **Monitor Upload Progress**
   - Each file shows "Upload" button while pending
   - Shows "Uploaded" with a checkmark when complete
   - The "Datasets" panel shows the total audio count for your dataset

### 4. Configure Training Parameters

Once all files are uploaded:

1. **Set Epochs**
   - Number of training iterations
   - Recommended: 10-100 for fine-tuning
   - More epochs = better quality but longer training time
   - Default: 10

2. **Set Batch Size**
   - Number of samples processed per iteration
   - Recommended: 4-16 depending on GPU memory
   - Smaller batch size = more stable but slower training
   - Default: 4

### 5. Start Training

1. Click "Start Training" button
2. The system will:
   - Validate your dataset
   - Download the base TTS model
   - Begin fine-tuning with your voice data
   - Create a training job

3. Monitor progress in the "Training Jobs" panel:
   - **Training...** (blue): Currently training
   - **Completed** (green): Training finished successfully
   - **Pending** (yellow): Waiting to start

## Example Workflow for 102 Files

### Batch Upload Strategy

To efficiently upload 102 files:

1. **Organize in groups of 10**
   - Add 10 file entries
   - Upload all 10
   - Repeat 10 times
   - Upload remaining 2 files

2. **Prepare transcripts in advance**
   - Have all transcripts ready in a text file
   - Copy-paste each transcript as you upload

3. **Use consistent naming**
   - Name files: `voice_001.mp3`, `voice_002.mp3`, etc.
   - Makes it easier to track progress

### Estimated Time

- **Upload**: ~2-5 minutes per file (depends on file size and network)
  - Total for 102 files: ~3-8 hours
  - Recommended: Upload in batches over multiple sessions
  
- **Training**: 30 minutes to several hours
  - Depends on epochs, batch size, and hardware
  - GPU recommended for faster training

## API Endpoints

### Admin Training Endpoints

All endpoints require admin authentication (Bearer token).

#### Upload Training File
```
POST /api/v1/admin/training/upload
Content-Type: multipart/form-data

Parameters:
- dataset_name: string (required) - Name of the dataset
- transcript: string (required) - Text transcript of the audio
- audio: file (required) - MP3, WAV, or M4A file

Response:
{
  "status": "stored",
  "dataset": "my_voice_dataset",
  "wav_filename": "audio_1234567890.wav",
  "metadata_path": "/app/data/my_voice_dataset/metadata.csv"
}
```

#### Start Training
```
POST /api/v1/admin/training/start
Content-Type: application/json

Body:
{
  "dataset_name": "my_voice_dataset",
  "epochs": 10,
  "batch_size": 4
}

Response:
{
  "status": "accepted",
  "job_id": "my_voice_dataset_20251204145511",
  "output_dir": "/app/training_runs/my_voice_dataset_20251204145511"
}
```

#### List Datasets
```
GET /api/v1/admin/training/datasets

Response:
{
  "status": "ok",
  "datasets": [
    {
      "name": "my_voice_dataset",
      "audio_count": 102,
      "created_at": "2025-12-04 14:55:11"
    }
  ]
}
```

#### List Training Jobs
```
GET /api/v1/admin/training/jobs

Response:
{
  "status": "ok",
  "jobs": [
    {
      "job_id": "my_voice_dataset_20251204145511",
      "created_at": "2025-12-04T14:55:11",
      "status": "training",
      "has_model": false,
      "log_file": "/app/training_runs/my_voice_dataset_20251204145511/train.log"
    }
  ]
}
```

## Troubleshooting

### Upload Issues

**Problem**: File upload fails
- **Solution**: Check file format (MP3, WAV, M4A supported)
- **Solution**: Ensure transcript is not empty
- **Solution**: Check dataset name is valid (no special characters)

**Problem**: "Audio conversion failed"
- **Solution**: Ensure FFmpeg is installed in the container
- **Solution**: Try a different MP3 file to test

### Training Issues

**Problem**: Training doesn't start
- **Solution**: Verify all files are uploaded (check audio count in Datasets panel)
- **Solution**: Check that metadata.csv exists in the dataset directory
- **Solution**: Ensure sufficient disk space for training

**Problem**: Training fails or crashes
- **Solution**: Check training job log file for errors
- **Solution**: Try with fewer epochs or smaller batch size
- **Solution**: Ensure audio files are valid and properly formatted

## File Structure

After uploading to dataset `my_voice_dataset`:

```
/app/data/my_voice_dataset/
├── metadata.csv          # Audio file list with transcripts
├── wavs/                 # Converted audio files
│   ├── audio_1.wav
│   ├── audio_2.wav
│   └── ... (102 files)
└── tmp/                  # Temporary files (cleaned up)
```

The metadata.csv format:
```
audio_1.wav|transcript for audio 1
audio_2.wav|transcript for audio 2
...
```

## Training Output

After training completes:

```
/app/training_runs/my_voice_dataset_20251204145511/
├── train.log                    # Training log
├── best_model_step_1000.pth     # Trained model
├── finetune_config.json         # Training configuration
└── ... (other checkpoint files)
```

## Using the Trained Model

After training completes, you can use the trained model for synthesis:

```
POST /api/v1/synthesize-trained
Content-Type: application/json

Body:
{
  "text": "Hello, this is my trained voice!",
  "job_id": "my_voice_dataset_20251204145511"
}

Response: Audio file (WAV)
```

## Best Practices

1. **Audio Quality**
   - Use clear, noise-free audio recordings
   - Consistent microphone and recording settings
   - Avoid background noise

2. **Transcript Accuracy**
   - Ensure transcripts exactly match the audio content
   - Include punctuation for better prosody
   - Use consistent capitalization

3. **Dataset Size**
   - Minimum: 10-20 audio files
   - Recommended: 50-200 files
   - Your 102 files is excellent

4. **Training Parameters**
   - Start with default values (10 epochs, batch size 4)
   - Increase epochs if quality is poor
   - Decrease batch size if running out of memory

5. **Monitoring**
   - Check training job status regularly
   - Monitor system resources during training
   - Save the job ID for future reference

## Support

For issues or questions:
1. Check the training job log file
2. Review the troubleshooting section
3. Verify file formats and transcripts
4. Check system resources (disk space, memory)
