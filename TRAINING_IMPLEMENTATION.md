# Training Feature Implementation Summary

## Overview

A complete training system has been added to the admin panel, allowing admins to upload MP3 files with transcripts and fine-tune TTS models with their own voice data.

## Changes Made

### 1. Go Service (`/go-service/main.go`)

#### New Endpoints Added

**Admin Training Upload**
- Route: `POST /api/v1/admin/training/upload`
- Function: `adminTrainingUploadHandler`
- Features:
  - Accepts multipart form data with audio file and transcript
  - Supports MP3, WAV, M4A formats
  - Automatically converts to WAV (22050 Hz, mono) using FFmpeg
  - Stores files in `/app/data/{dataset_name}/wavs/`
  - Updates metadata.csv with filename and transcript
  - Admin-only access

**Admin Training Start**
- Route: `POST /api/v1/admin/training/start`
- Function: `adminTrainingStartHandler`
- Features:
  - Accepts dataset_name, epochs, batch_size
  - Proxies to Python service /train endpoint
  - Returns job_id for tracking
  - Admin-only access

**Admin List Datasets**
- Route: `GET /api/v1/admin/training/datasets`
- Function: `adminListDatasetsHandler`
- Features:
  - Lists all datasets in `/app/data/`
  - Shows audio file count for each dataset
  - Shows creation timestamp
  - Admin-only access

**Admin List Training Jobs**
- Route: `GET /api/v1/admin/training/jobs`
- Function: `adminListTrainingJobsHandler`
- Features:
  - Proxies to Python service /training-jobs endpoint
  - Shows job status (training/completed)
  - Shows model availability
  - Admin-only access

### 2. Python Service (`/python-service/app/main.py`)

#### New Endpoints Added

**List Training Jobs**
- Route: `GET /training-jobs`
- Function: `list_training_jobs`
- Features:
  - Lists all training jobs in `/app/training_runs/`
  - Shows job status (training/completed)
  - Shows if model is available
  - Shows log file path
  - Returns job creation timestamp

### 3. Frontend Admin Panel (`/frontend/app/admin/page.tsx`)

#### New Features

**Tab Navigation**
- Added "Users" and "Training" tabs
- Allows switching between user management and training

**Training Tab Components**

1. **Upload Training Data Card**
   - Dataset name input
   - Add/remove audio file entries
   - File picker for each audio file
   - Transcript input for each file
   - Individual upload buttons
   - Status indicators (Uploaded/Pending)

2. **Training Configuration Card**
   - Epochs input (1-1000)
   - Batch size input (1-64)
   - Start Training button
   - Validation before submission

3. **Datasets Card**
   - Lists all available datasets
   - Shows audio file count
   - Shows creation date
   - Scrollable list for many datasets

4. **Training Jobs Card**
   - Lists all training jobs
   - Shows job ID
   - Shows creation timestamp
   - Shows status with color coding:
     - Blue: Training in progress
     - Green: Completed
     - Yellow: Pending
   - Auto-refreshes every 5 seconds when on Training tab

#### State Management

New state variables:
- `activeTab`: Current tab (users/training)
- `datasets`: List of available datasets
- `datasetName`: Current dataset name
- `trainingFiles`: Array of file entries to upload
- `uploading`: Upload in progress flag
- `training`: Training start in progress flag
- `epochs`: Number of training epochs
- `batchSize`: Batch size for training
- `trainingJobs`: List of training jobs

#### Functions

- `fetchDatasets()`: Get list of datasets
- `fetchTrainingJobs()`: Get list of training jobs
- `handleAddTrainingFile()`: Add new file entry
- `handleRemoveTrainingFile()`: Remove file entry
- `handleFileChange()`: Update file selection
- `handleTranscriptChange()`: Update transcript
- `handleUploadTrainingFile()`: Upload single file
- `handleStartTraining()`: Start training job

## Workflow

### Upload Workflow

1. Admin enters dataset name
2. Admin clicks "Add Audio File" multiple times
3. For each file:
   - Select MP3 file
   - Enter transcript
   - Click Upload
   - File is sent to Go service
   - Go service converts MP3 to WAV
   - WAV is stored in dataset directory
   - Metadata is updated
4. Dataset appears in "Datasets" panel

### Training Workflow

1. All files uploaded to a dataset
2. Admin sets epochs and batch size
3. Admin clicks "Start Training"
4. Request sent to Go service
5. Go service proxies to Python service
6. Python service:
   - Validates dataset
   - Downloads base model
   - Starts fine-tuning job
   - Returns job_id
7. Job appears in "Training Jobs" panel
8. Status updates automatically every 5 seconds
9. When complete, model is available for synthesis

## File Structure

### Data Directory
```
/app/data/
└── {dataset_name}/
    ├── metadata.csv
    ├── wavs/
    │   ├── audio_1.wav
    │   ├── audio_2.wav
    │   └── ...
    └── tmp/
```

### Training Output Directory
```
/app/training_runs/
└── {dataset_name}_{timestamp}/
    ├── train.log
    ├── best_model_step_*.pth
    ├── finetune_config.json
    └── ...
```

## API Flow

```
Frontend (Admin Panel)
    ↓
Go Service (/api/v1/admin/training/*)
    ↓
Python Service (/train, /training-jobs)
    ↓
File System & TTS Model
```

## Security

- All training endpoints require admin authentication
- Admin middleware validates JWT token and role
- File uploads are validated for format and size
- Dataset names are sanitized to prevent directory traversal

## MP3 Support

- Frontend accepts MP3, WAV, M4A files
- Go service uses FFmpeg to convert to WAV
- Conversion parameters:
  - Codec: pcm_s16le
  - Sample rate: 22050 Hz
  - Channels: mono (if stereo, converted to mono)
- Temporary MP3 files are deleted after conversion

## Performance Considerations

- File uploads are sequential (one at a time)
- Training jobs run in background
- Frontend auto-refreshes training status every 5 seconds
- Large datasets may take time to upload (102 files ~3-8 hours)
- Training time depends on:
  - Number of epochs
  - Batch size
  - Hardware (GPU recommended)
  - Dataset size

## Error Handling

- File upload validation
- Dataset validation before training
- Proper error messages to user
- Toast notifications for success/failure
- Graceful handling of network errors

## Testing Checklist

- [ ] Upload single MP3 file with transcript
- [ ] Upload multiple files in batch
- [ ] Verify files appear in Datasets panel
- [ ] Start training with valid dataset
- [ ] Monitor training job status
- [ ] Verify training completes successfully
- [ ] Use trained model for synthesis
- [ ] Test with 102 files
- [ ] Verify MP3 conversion works
- [ ] Test error cases (invalid files, missing transcripts)

## Future Enhancements

- Batch upload UI (drag and drop)
- Progress bar for multi-file uploads
- Dataset editing/deletion
- Training job cancellation
- Training metrics visualization
- Model evaluation interface
- Synthesis testing with trained models
