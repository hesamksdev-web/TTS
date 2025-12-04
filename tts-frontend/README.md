# TTS Fine-tune Dashboard

A modern web dashboard for fine-tuning Text-to-Speech models with custom audio data.

## 🎯 Features

- **Audio Upload**: Upload audio files for training
- **Text Input**: Provide corresponding text for the audio
- **Live Progress**: Real-time training progress tracking
- **Audio Playback**: Listen to generated audio output
- **Beautiful UI**: Built with Tailwind CSS and shadcn/ui components
- **Responsive Design**: Works on desktop and mobile devices
- **Dark Mode Support**: Full dark mode support

## 🛠️ Tech Stack

- **Frontend Framework**: Next.js 16 with TypeScript
- **Styling**: Tailwind CSS v4
- **UI Components**: shadcn/ui
- **Icons**: Lucide React
- **HTTP Client**: Axios
- **Notifications**: Sonner

## 🚀 Getting Started

### Prerequisites

- Node.js 18+ installed
- npm or yarn package manager

### Installation

1. Navigate to the project directory:
```bash
cd tts-frontend
```

2. Install dependencies (already done):
```bash
npm install
```

3. Run the development server:
```bash
npm run dev
```

4. Open [http://localhost:3000](http://localhost:3000) in your browser

## 📝 Configuration

The dashboard expects a backend API at `http://localhost:8000` with the following endpoints:

### POST `/api/finetune`
Upload audio file and text for fine-tuning.

**Request:**
```
Content-Type: multipart/form-data
- audio: File (audio file)
- text: string (corresponding text)
```

**Response:**
```json
{
  "job_id": "unique-job-id"
}
```

### GET `/api/finetune/{job_id}`
Get the status of a fine-tuning job.

**Response:**
```json
{
  "status": "uploading|training|completed|error",
  "progress": 0-100,
  "audio_url": "url-to-generated-audio",
  "error": "error message if status is error"
}
```

## 🎨 UI Components

- **Card**: Display sections with headers and content
- **Button**: Interactive buttons with loading states
- **Input**: File upload and text input fields
- **Progress**: Visual progress bar for training
- **Alert**: Error and status messages
- **Label**: Form labels

## 📦 Build

To create a production build:

```bash
npm run build
npm start
```

## 🔧 Development

The main dashboard component is located in `app/page.tsx`. It includes:

- File upload handling
- Form validation
- API communication with progress tracking
- Audio playback controls
- Status management

## 📱 Responsive Breakpoints

- Mobile: < 640px
- Tablet: 640px - 1024px
- Desktop: > 1024px

## 🌙 Dark Mode

Dark mode is automatically supported through Tailwind CSS. The UI adapts based on system preferences.

## 📄 License

This project is part of the TTS Fine-tune system.
