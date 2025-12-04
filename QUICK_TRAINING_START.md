# Quick Start: Training with 102 MP3 Files

## Your Scenario

You have:
- 102 MP3 files of your voice
- Transcripts for all files
- Want to train a custom TTS model

## Step-by-Step Guide

### Step 1: Prepare Your Files

1. Organize your 102 MP3 files
2. Have transcripts ready (text of what's spoken in each file)
3. Note: Files can be any format (MP3, WAV, M4A) - they'll be auto-converted

### Step 2: Access Admin Training Panel

1. Log in with admin account
2. Go to Admin Dashboard
3. Click "Training" tab

### Step 3: Upload Files (Batch Method)

**Recommended: Upload in groups of 10**

#### Batch 1 (Files 1-10)
1. Enter dataset name: `my_voice_dataset` (or your preferred name)
2. Click "Add Audio File" 10 times
3. For each file:
   - Click file input → select MP3 file
   - Enter transcript
   - Click "Upload" button
   - Wait for "Uploaded" confirmation
4. Move to next batch

#### Repeat for Batches 2-10
- Same process for files 11-20, 21-30, etc.

#### Final Batch (Files 101-102)
- Add 2 file entries
- Upload both files

**Total Time**: ~3-8 hours (depending on file sizes and network)

### Step 4: Verify All Files Uploaded

1. Look at "Datasets" panel on the right
2. Should show: `my_voice_dataset` with **102 audio files**
3. If count is less, continue uploading remaining files

### Step 5: Configure Training

1. **Set Epochs**: 
   - Start with 10 (recommended for 102 files)
   - Can increase to 20-50 for better quality
   
2. **Set Batch Size**: 
   - Keep at 4 (good for most systems)
   - Reduce to 2 if running out of memory

### Step 6: Start Training

1. Click "Start Training" button
2. Wait for success message
3. Job ID will be shown in toast notification
4. Example: `my_voice_dataset_20251204145511`

### Step 7: Monitor Training

1. Check "Training Jobs" panel
2. Status will show:
   - **Training...** (blue) - Currently training
   - **Completed** (green) - Done!
3. Panel auto-refreshes every 5 seconds
4. Training typically takes 30 minutes to several hours

### Step 8: Use Your Trained Model

Once training completes (green "Completed" status):

```bash
# Using the API
curl -X POST http://localhost:8080/api/v1/synthesize-trained \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "text": "Hello, this is my trained voice!",
    "job_id": "my_voice_dataset_20251204145511"
  }'
```

Or use the frontend to test synthesis with your trained model.

## Troubleshooting

### Upload Stuck?
- Check internet connection
- Try uploading smaller batch (5 files instead of 10)
- Refresh page and continue

### File Upload Fails?
- Verify MP3 file is valid
- Check transcript is not empty
- Try different MP3 file to test

### Training Won't Start?
- Verify all 102 files show in Datasets panel
- Check that dataset name matches what you entered
- Ensure sufficient disk space (need ~1-2 GB)

### Training Fails?
- Check training job log for errors
- Try with fewer epochs (5 instead of 10)
- Verify audio files are valid

## Expected Results

After training with 102 files:
- Custom TTS model trained on your voice
- Can synthesize speech in your voice
- Quality depends on:
  - Audio quality
  - Transcript accuracy
  - Number of epochs
  - Diversity of content

## Tips for Best Results

1. **Audio Quality**
   - Record in quiet environment
   - Use consistent microphone
   - Avoid background noise
   - Normalize volume levels

2. **Transcripts**
   - Be 100% accurate
   - Include punctuation
   - Use consistent capitalization
   - Match audio exactly

3. **Content Variety**
   - Include different phonemes
   - Vary sentence structure
   - Include numbers, special characters
   - Cover different emotions/styles

4. **Training**
   - Start with 10 epochs
   - Increase if quality is poor
   - Monitor training progress
   - Save job ID for reference

## File Size Estimates

- 102 MP3 files: ~50-200 MB (depends on quality)
- Converted to WAV: ~200-800 MB
- Training output: ~100-300 MB
- **Total space needed**: ~1-2 GB

## Timeline

- **Upload**: 3-8 hours (can do in multiple sessions)
- **Training**: 30 minutes - 2 hours
- **Total**: 4-10 hours (spread over time)

## Next Steps

1. Upload your 102 files
2. Start training
3. Monitor progress
4. Test with your trained model
5. Adjust parameters if needed
6. Deploy your custom TTS model

## Support

For detailed information, see:
- `TRAINING_GUIDE.md` - Complete training documentation
- `TRAINING_IMPLEMENTATION.md` - Technical implementation details

Good luck with your training! 🎉
