"use client";

import { useState, useRef } from "react";
import axios from "axios";
import { toast } from "sonner";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Label } from "@/components/ui/label";
import { Progress } from "@/components/ui/progress";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { Upload, Play, Pause, Volume2, Loader2, CheckCircle2, AlertCircle } from "lucide-react";

interface TrainingStatus {
  status: "idle" | "uploading" | "training" | "completed" | "error";
  progress: number;
  message: string;
  audioUrl?: string;
}

export default function Dashboard() {
  const [audioFile, setAudioFile] = useState<File | null>(null);
  const [textInput, setTextInput] = useState("");
  const [trainingStatus, setTrainingStatus] = useState<TrainingStatus>({
    status: "idle",
    progress: 0,
    message: "آماده برای شروع",
  });
  const [isPlaying, setIsPlaying] = useState(false);
  const audioRef = useRef<HTMLAudioElement>(null);

  const handleAudioFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      if (file.type.startsWith("audio/")) {
        setAudioFile(file);
        toast.success("فایل صوتی انتخاب شد");
      } else {
        toast.error("لطفاً یک فایل صوتی انتخاب کنید");
      }
    }
  };

  const handleStartTraining = async () => {
    if (!audioFile || !textInput.trim()) {
      toast.error("لطفاً فایل صوتی و متن را وارد کنید");
      return;
    }

    const formData = new FormData();
    formData.append("audio", audioFile);
    formData.append("text", textInput);

    try {
      setTrainingStatus({
        status: "uploading",
        progress: 10,
        message: "در حال آپلود فایل...",
      });

      // Simulate API call - replace with your actual backend URL
      const response = await axios.post(
        "http://localhost:8000/api/finetune",
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
          onUploadProgress: (progressEvent) => {
            const percentCompleted = Math.round(
              (progressEvent.loaded * 100) / (progressEvent.total || 1)
            );
            setTrainingStatus((prev) => ({
              ...prev,
              progress: Math.min(percentCompleted, 90),
              message: `آپلود: ${percentCompleted}%`,
            }));
          },
        }
      );

      setTrainingStatus({
        status: "training",
        progress: 50,
        message: "در حال Fine-tune کردن مدل...",
      });

      // Poll for training completion
      let attempts = 0;
      const maxAttempts = 120; // 2 minutes with 1-second intervals

      while (attempts < maxAttempts) {
        const statusResponse = await axios.get(
          `http://localhost:8000/api/finetune/${response.data.job_id}`
        );

        if (statusResponse.data.status === "completed") {
          setTrainingStatus({
            status: "completed",
            progress: 100,
            message: "تمام شد! مدل آماده است",
            audioUrl: statusResponse.data.audio_url,
          });
          toast.success("Fine-tune با موفقیت تمام شد!");
          break;
        } else if (statusResponse.data.status === "error") {
          setTrainingStatus({
            status: "error",
            progress: 0,
            message: `خطا: ${statusResponse.data.error}`,
          });
          toast.error("خطا در Fine-tune");
          break;
        }

        const currentProgress = statusResponse.data.progress || 50;
        setTrainingStatus((prev) => ({
          ...prev,
          progress: currentProgress,
          message: `در حال Fine-tune: ${currentProgress}%`,
        }));

        await new Promise((resolve) => setTimeout(resolve, 1000));
        attempts++;
      }

      if (attempts >= maxAttempts) {
        setTrainingStatus({
          status: "error",
          progress: 0,
          message: "زمان انتظار تمام شد",
        });
        toast.error("زمان انتظار تمام شد");
      }
    } catch (error) {
      setTrainingStatus({
        status: "error",
        progress: 0,
        message: "خطا در ارسال درخواست",
      });
      toast.error("خطا در ارسال درخواست");
      console.error(error);
    }
  };

  const togglePlayback = () => {
    if (audioRef.current) {
      if (isPlaying) {
        audioRef.current.pause();
      } else {
        audioRef.current.play();
      }
      setIsPlaying(!isPlaying);
    }
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-slate-50 to-slate-100 dark:from-slate-950 dark:to-slate-900 p-4 md:p-8">
      <div className="max-w-4xl mx-auto">
        {/* Header */}
        <div className="mb-8">
          <h1 className="text-4xl font-bold text-slate-900 dark:text-slate-50 mb-2">
            TTS Fine-tune Dashboard
          </h1>
          <p className="text-lg text-slate-600 dark:text-slate-400">
            مدل Text-to-Speech خود را با داده‌های شخصی بهبود دهید
          </p>
        </div>

        {/* Main Content */}
        <div className="grid gap-6">
          {/* Upload Section */}
          <Card className="border-slate-200 dark:border-slate-800">
            <CardHeader>
              <CardTitle className="flex items-center gap-2">
                <Upload className="w-5 h-5" />
                آپلود فایل‌ها
              </CardTitle>
              <CardDescription>
                فایل صوتی و متن متناظر را آپلود کنید
              </CardDescription>
            </CardHeader>
            <CardContent className="space-y-6">
              {/* Audio File Upload */}
              <div className="space-y-2">
                <Label htmlFor="audio-file" className="text-base font-semibold">
                  فایل صوتی
                </Label>
                <div className="flex items-center gap-3">
                  <Input
                    id="audio-file"
                    type="file"
                    accept="audio/*"
                    onChange={handleAudioFileChange}
                    disabled={trainingStatus.status === "uploading" || trainingStatus.status === "training"}
                    className="flex-1"
                  />
                  {audioFile && (
                    <span className="text-sm text-green-600 dark:text-green-400 font-medium">
                      ✓ {audioFile.name}
                    </span>
                  )}
                </div>
              </div>

              {/* Text Input */}
              <div className="space-y-2">
                <Label htmlFor="text-input" className="text-base font-semibold">
                  متن متناظر
                </Label>
                <textarea
                  id="text-input"
                  placeholder="متن صوتی را اینجا وارد کنید..."
                  value={textInput}
                  onChange={(e) => setTextInput(e.target.value)}
                  disabled={trainingStatus.status === "uploading" || trainingStatus.status === "training"}
                  className="w-full p-3 border border-slate-300 dark:border-slate-700 rounded-lg bg-white dark:bg-slate-950 text-slate-900 dark:text-slate-50 placeholder-slate-500 dark:placeholder-slate-400 focus:outline-none focus:ring-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed"
                  rows={4}
                />
              </div>

              {/* Start Button */}
              <Button
                onClick={handleStartTraining}
                disabled={
                  !audioFile ||
                  !textInput.trim() ||
                  trainingStatus.status === "uploading" ||
                  trainingStatus.status === "training"
                }
                className="w-full h-12 text-base font-semibold"
              >
                {trainingStatus.status === "uploading" || trainingStatus.status === "training" ? (
                  <>
                    <Loader2 className="w-4 h-4 mr-2 animate-spin" />
                    در حال پردازش...
                  </>
                ) : (
                  "شروع Fine-tune"
                )}
              </Button>
            </CardContent>
          </Card>

          {/* Status Section */}
          {trainingStatus.status !== "idle" && (
            <Card className={`border-2 ${
              trainingStatus.status === "completed"
                ? "border-green-500 dark:border-green-600"
                : trainingStatus.status === "error"
                ? "border-red-500 dark:border-red-600"
                : "border-blue-500 dark:border-blue-600"
            }`}>
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  {trainingStatus.status === "completed" && (
                    <>
                      <CheckCircle2 className="w-5 h-5 text-green-600 dark:text-green-400" />
                      تمام شد
                    </>
                  )}
                  {trainingStatus.status === "error" && (
                    <>
                      <AlertCircle className="w-5 h-5 text-red-600 dark:text-red-400" />
                      خطا
                    </>
                  )}
                  {(trainingStatus.status === "uploading" || trainingStatus.status === "training") && (
                    <>
                      <Loader2 className="w-5 h-5 text-blue-600 dark:text-blue-400 animate-spin" />
                      در حال پردازش
                    </>
                  )}
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <p className="text-sm text-slate-600 dark:text-slate-400">
                  {trainingStatus.message}
                </p>
                <Progress
                  value={trainingStatus.progress}
                  className="h-2"
                />
              </CardContent>
            </Card>
          )}

          {/* Audio Playback Section */}
          {trainingStatus.status === "completed" && trainingStatus.audioUrl && (
            <Card className="border-green-200 dark:border-green-900 bg-green-50 dark:bg-green-950">
              <CardHeader>
                <CardTitle className="flex items-center gap-2">
                  <Volume2 className="w-5 h-5" />
                  صوت تولید شده
                </CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <audio
                  ref={audioRef}
                  src={trainingStatus.audioUrl}
                  onEnded={() => setIsPlaying(false)}
                  className="hidden"
                />
                <div className="flex items-center gap-4">
                  <Button
                    onClick={togglePlayback}
                    variant="outline"
                    size="lg"
                    className="gap-2"
                  >
                    {isPlaying ? (
                      <>
                        <Pause className="w-4 h-4" />
                        توقف
                      </>
                    ) : (
                      <>
                        <Play className="w-4 h-4" />
                        پخش
                      </>
                    )}
                  </Button>
                  <a
                    href={trainingStatus.audioUrl}
                    download
                    className="text-blue-600 dark:text-blue-400 hover:underline text-sm font-medium"
                  >
                    دانلود فایل
                  </a>
                </div>
              </CardContent>
            </Card>
          )}

          {/* Error Alert */}
          {trainingStatus.status === "error" && (
            <Alert className="border-red-200 dark:border-red-900 bg-red-50 dark:bg-red-950">
              <AlertCircle className="h-4 w-4 text-red-600 dark:text-red-400" />
              <AlertDescription className="text-red-800 dark:text-red-200">
                {trainingStatus.message}
              </AlertDescription>
            </Alert>
          )}
        </div>
      </div>
    </div>
  );
}
