"use client";

import { useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";
import { toast } from "sonner";
import { Upload, Play, Mic, LogOut, Settings, Zap } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "/api/v1";

const getErrorMessage = (error: any): string => {
  if (typeof error === "string") return error;
  if (error?.response?.data?.error) return error.response.data.error;
  if (error?.response?.data?.message) return error.response.data.message;
  if (error?.response?.data?.detail) return error.response.data.detail;
  if (typeof error?.response?.data === "string") return error.response.data;
  return "An error occurred";
};

export default function Dashboard() {
  const router = useRouter();
  const [token, setToken] = useState<string | null>(null);
  const [role, setRole] = useState<string | null>(null);

  const [status, setStatus] = useState("");
  const [textToSpeak, setTextToSpeak] = useState("Hello, this is a test of the voice synthesis system.");
  const [audioUrl, setAudioUrl] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  const [selectedSpeaker, setSelectedSpeaker] = useState("fa_standard");
  const [synthesisMode, setSynthesisMode] = useState<"pretrained" | "trained">("pretrained");
  const [trainedJobId, setTrainedJobId] = useState("");
  const [trainedModels, setTrainedModels] = useState<Array<{ job_id: string; model_path: string }>>([]);
  
  // Voice Clone states
  const [voiceCloneFile, setVoiceCloneFile] = useState<File | null>(null);
  const [voiceCloneLanguage, setVoiceCloneLanguage] = useState("fa");
  const [voiceCloneText, setVoiceCloneText] = useState("سلام، این یک نمونه از صدای شما است.");
  const [voiceCloneSpeed, setVoiceCloneSpeed] = useState(1.0);
  const [clonedAudioUrl, setClonedAudioUrl] = useState<string | null>(null);
  const [voiceCloneJobId, setVoiceCloneJobId] = useState<number | null>(null);
  const [voiceCloneJobStatus, setVoiceCloneJobStatus] = useState<string | null>(null);
  const jobPollTimeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const jobStartTimeRef = useRef<number | null>(null);
  const [notifications, setNotifications] = useState<Array<{ id: number; message: string; read: boolean; created_at: string }>>([]);
  const [unreadNotifications, setUnreadNotifications] = useState<number>(0);

  const speakers = [
    { id: "fa_standard", label: "🇮🇷 Persian (Farsi) - VITS" },
    { id: "en_standard", label: "🇬🇧 English - Tacotron2 DDC" },
    { id: "de_standard", label: "🇩🇪 German - Tacotron2 DCA" },
  ];

  useEffect(() => {
    const t = localStorage.getItem("token");
    const r = localStorage.getItem("role");
    if (!t) {
      router.push("/login");
    } else {
      setToken(t);
      setRole(r);
    }
  }, []);

  useEffect(() => {
    if (token && synthesisMode === "trained") {
      fetchTrainedModels();
    }
  }, [token, synthesisMode]);

  useEffect(() => {
    if (!token) return;
    fetchNotifications();
    const interval = setInterval(fetchNotifications, 15000);
    return () => clearInterval(interval);
  }, [token]);

  useEffect(() => {
    return () => {
      if (jobPollTimeoutRef.current) {
        clearTimeout(jobPollTimeoutRef.current);
      }
    };
  }, []);

  const fetchTrainedModels = async () => {
    try {
      const res = await axios.get(`${API_BASE_URL}/trained-models`, getHeaders());
      if (res.data.models && res.data.models.length > 0) {
        setTrainedModels(res.data.models);
        if (!trainedJobId) {
          setTrainedJobId(res.data.models[0].job_id);
        }
      }
    } catch (error) {
      console.error("Failed to fetch trained models:", error);
    }
  };

  const getHeaders = () => ({
    headers: { Authorization: `Bearer ${token}` },
  });

  const pollVoiceCloneJob = async (jobId: number, delay = 5000) => {
    if (!jobId || !token) return;

    // Check if job has been polling for too long (2 hours = 7200 seconds)
    const elapsedSeconds = jobStartTimeRef.current ? (Date.now() - jobStartTimeRef.current) / 1000 : 0;
    if (elapsedSeconds > 7200) {
      setStatus("❌ Voice cloning job timeout: Processing took too long");
      setVoiceCloneJobStatus("timeout");
      toast.error("Voice cloning job timeout after 2 hours");
      return;
    }

    try {
      const res = await axios.get(`${API_BASE_URL}/voice-clone/job`, {
        ...getHeaders(),
        params: { job_id: jobId },
      });
      const job = res.data;
      console.log("Poll result:", job);
      setVoiceCloneJobStatus(job.status);

      if (job.status === "completed") {
        setStatus("✅ Voice clone ready! Downloading...");
        await downloadVoiceCloneJob(jobId);
        fetchNotifications();
      } else if (job.status === "failed") {
        setStatus(`❌ Voice cloning failed: ${job.error_message || "Unknown error"}`);
        toast.error(job.error_message || "Voice cloning failed");
      } else {
        console.log(`Job ${jobId} still ${job.status}, polling again in ${delay}ms (elapsed: ${elapsedSeconds.toFixed(0)}s)`);
        jobPollTimeoutRef.current = setTimeout(() => pollVoiceCloneJob(jobId), delay);
      }
    } catch (error: any) {
      console.error("Poll error:", error);
      toast.error(getErrorMessage(error));
    }
  };

  const downloadVoiceCloneJob = async (jobId: number) => {
    try {
      const res = await axios.get(`${API_BASE_URL}/voice-clone/job/download`, {
        ...getHeaders(),
        params: { job_id: jobId },
        responseType: "blob",
      });
      const url = window.URL.createObjectURL(new Blob([res.data]));
      setClonedAudioUrl(url);
      setStatus("✅ Voice clone ready!");
      toast.success("Voice clone ready!");
    } catch (error: any) {
      toast.error(getErrorMessage(error));
    }
  };

  const fetchNotifications = async () => {
    if (!token) return;
    try {
      const res = await axios.get(`${API_BASE_URL}/notifications`, getHeaders());
      setNotifications(res.data.notifications || []);
      setUnreadNotifications(res.data.unread_count || 0);
    } catch (error) {
      console.error("Failed to fetch notifications", error);
    }
  };

  const markNotificationsRead = async () => {
    try {
      await axios.post(`${API_BASE_URL}/notifications/read`, {}, getHeaders());
      fetchNotifications();
    } catch (error) {
      console.error("Failed to mark notifications read", error);
    }
  };

  const handleSynthesize = async () => {
    setLoading(true);
    setStatus("🗣️ Generating audio...");
    setAudioUrl(null);

    try {
      let endpoint = `${API_BASE_URL}/synthesize`;
      let payload: any = { text: textToSpeak };

      if (synthesisMode === "trained") {
        if (!trainedJobId) {
          toast.error("Please enter a trained model Job ID");
          setLoading(false);
          return;
        }
        endpoint = `${API_BASE_URL}/synthesize-trained`;
        payload.job_id = trainedJobId;
      } else {
        payload.speaker_id = selectedSpeaker;
      }

      const res = await axios.post(endpoint, payload, {
        ...getHeaders(),
        responseType: "blob",
      });

      const url = window.URL.createObjectURL(new Blob([res.data]));
      setAudioUrl(url);
      setStatus("✅ Audio generated!");
      toast.success("Audio generated successfully!");
    } catch (error: any) {
      setStatus("❌ Generation failed");
      toast.error(getErrorMessage(error));
    } finally {
      setLoading(false);
    }
  };

  const handleVoiceClone = async () => {
    if (!voiceCloneFile) {
      toast.error("Please select an audio file");
      return;
    }

    setLoading(true);
    setStatus("🎙️ Cloning voice (queued)...");
    setClonedAudioUrl(null);
    setVoiceCloneJobStatus(null);
    setVoiceCloneJobId(null);

    try {
      const formData = new FormData();
      formData.append("audio", voiceCloneFile);
      formData.append("text", voiceCloneText);
      formData.append("language", voiceCloneLanguage);
      formData.append("speed", voiceCloneSpeed.toString());

      const res = await axios.post(`${API_BASE_URL}/voice-clone`, formData, getHeaders());
      const jobId = res.data?.job_id;
      setVoiceCloneJobId(jobId);
      setVoiceCloneJobStatus(res.data?.status || "pending");
      jobStartTimeRef.current = Date.now();
      toast.success(`Voice clone job queued #${jobId}`);
      pollVoiceCloneJob(jobId);
    } catch (error: any) {
      setStatus("❌ Voice cloning failed");
      toast.error(getErrorMessage(error));
    } finally {
      setLoading(false);
    }
  };

  const logout = () => {
    localStorage.clear();
    router.push("/login");
  };

  if (!token) return null;

  return (
    <div className="min-h-screen bg-white">
      {/* Header */}
      <div className="border-b border-slate-200 bg-slate-50">
        <div className="max-w-7xl mx-auto px-6 py-6 flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-slate-900">Voice Studio</h1>
            <p className="text-slate-500 text-sm mt-1">Train and synthesize custom voices</p>
          </div>
          <div className="flex gap-3">
            {role === "admin" && (
              <Button
                variant="outline"
                onClick={() => router.push("/admin")}
                className="gap-2"
              >
                <Settings className="w-4 h-4" />
                Admin
              </Button>
            )}
            <Button
              variant="destructive"
              onClick={logout}
              className="gap-2"
            >
              <LogOut className="w-4 h-4" />
              Logout
            </Button>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="max-w-7xl mx-auto px-6 py-8">
        {/* Status Message */}
        {status && (
          <div className={`mb-6 p-4 rounded-lg border-l-4 text-sm font-medium ${
            status.includes("❌")
              ? "bg-red-50 border-red-500 text-red-800"
              : "bg-green-50 border-green-500 text-green-800"
          }`}>
            {status}
          </div>
        )}

        <div className="grid grid-cols-1 gap-8">
          {/* Synthesize Section */}
          <Card className="border-0 shadow-sm">
            <CardHeader className="border-b border-slate-200 bg-slate-50">
              <CardTitle className="flex items-center gap-2 text-lg">
                <Mic className="w-5 h-5 text-green-600" />
                Generate Audio
              </CardTitle>
            </CardHeader>
            <CardContent className="p-6 space-y-4">
              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Synthesis Mode
                </label>
                <div className="flex gap-3">
                  <button
                    onClick={() => setSynthesisMode("pretrained")}
                    className={`flex-1 px-3 py-2 rounded-md text-sm font-medium transition ${
                      synthesisMode === "pretrained"
                        ? "bg-green-600 text-white"
                        : "bg-slate-100 text-slate-700 hover:bg-slate-200"
                    }`}
                  >
                    🌍 Pre-trained
                  </button>
                  <button
                    onClick={() => setSynthesisMode("trained")}
                    className={`flex-1 px-3 py-2 rounded-md text-sm font-medium transition ${
                      synthesisMode === "trained"
                        ? "bg-blue-600 text-white"
                        : "bg-slate-100 text-slate-700 hover:bg-slate-200"
                    }`}
                  >
                    🎓 Trained Model
                  </button>
                </div>
              </div>

              {synthesisMode === "pretrained" ? (
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-2">
                    Speaker Voice
                  </label>
                  <select
                    value={selectedSpeaker}
                    onChange={(e) => setSelectedSpeaker(e.target.value)}
                    className="w-full px-3 py-2 border border-slate-200 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  >
                    {speakers.map((speaker) => (
                      <option key={speaker.id} value={speaker.id}>
                        {speaker.label}
                      </option>
                    ))}
                  </select>
                </div>
              ) : (
                <div>
                  <label className="block text-sm font-medium text-slate-700 mb-2">
                    Select Trained Model
                  </label>
                  {trainedModels.length > 0 ? (
                    <select
                      value={trainedJobId}
                      onChange={(e) => setTrainedJobId(e.target.value)}
                      className="w-full px-3 py-2 border border-slate-200 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                    >
                      {trainedModels.map((model) => (
                        <option key={model.job_id} value={model.job_id}>
                          🎓 {model.job_id}
                        </option>
                      ))}
                    </select>
                  ) : (
                    <div className="p-3 bg-amber-50 border border-amber-200 rounded-md text-sm text-amber-700">
                      No trained models found. Train a model first!
                    </div>
                  )}
                </div>
              )}

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Text to Synthesize
                </label>
                <textarea
                  placeholder="Enter text to convert to speech"
                  value={textToSpeak}
                  onChange={(e) => setTextToSpeak(e.target.value)}
                  className="w-full px-3 py-2 border border-slate-200 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-green-500 focus:border-transparent"
                  rows={4}
                />
              </div>

              <Button
                onClick={handleSynthesize}
                disabled={loading}
                className="w-full bg-green-600 hover:bg-green-700"
              >
                <Play className="w-4 h-4 mr-2" />
                Generate Audio
              </Button>

              {audioUrl && (
                <div className="mt-6 p-4 bg-slate-50 rounded-lg border border-slate-200">
                  <p className="text-xs font-medium text-slate-600 mb-3 uppercase tracking-wider">
                    Generated Audio
                  </p>
                  <audio
                    controls
                    src={audioUrl}
                    className="w-full"
                  />
                </div>
              )}
            </CardContent>
          </Card>

          {/* Voice Clone Section */}
          <Card className="border-0 shadow-sm">
            <CardHeader className="border-b border-slate-200 bg-slate-50">
              <CardTitle className="flex items-center gap-2 text-lg">
                <Mic className="w-5 h-5 text-purple-600" />
                Text to Speech by Language
              </CardTitle>
            </CardHeader>
            <CardContent className="p-6 space-y-4">
              <div className="p-3 bg-blue-50 border border-blue-200 rounded-md text-sm text-blue-700">
                💡 Select a language and enter text to synthesize speech in that language. Voice clone jobs run asynchronously; you will be notified when the output is ready.
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Language
                </label>
                <select
                  value={voiceCloneLanguage}
                  onChange={(e) => setVoiceCloneLanguage(e.target.value)}
                  className="w-full px-3 py-2 border border-slate-200 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent"
                >
                  <option value="fa">🇮🇷 Persian (Farsi)</option>
                  <option value="en">🇬🇧 English</option>
                  <option value="de">🇩🇪 German</option>
                </select>
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Voice Sample Audio File
                </label>
                <Input
                  type="file"
                  accept="audio/*"
                  onChange={(e) => setVoiceCloneFile(e.target.files?.[0] || null)}
                  className="border-slate-200"
                />
                {voiceCloneFile && (
                  <p className="text-xs text-slate-500 mt-2">
                    Selected: {voiceCloneFile.name}
                  </p>
                )}
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Speech Speed: {voiceCloneSpeed.toFixed(1)}x
                </label>
                <input
                  type="range"
                  min="0.5"
                  max="2.0"
                  step="0.1"
                  value={voiceCloneSpeed}
                  onChange={(e) => setVoiceCloneSpeed(parseFloat(e.target.value))}
                  className="w-full"
                />
                <p className="text-xs text-slate-500 mt-1">0.5 (slow) to 2.0 (fast)</p>
              </div>

              <div>
                <label className="block text-sm font-medium text-slate-700 mb-2">
                  Text to Speak
                </label>
                <textarea
                  placeholder="Enter text to be spoken in the selected language"
                  value={voiceCloneText}
                  onChange={(e) => setVoiceCloneText(e.target.value)}
                  className="w-full px-3 py-2 border border-slate-200 rounded-md text-sm focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-transparent"
                  rows={4}
                />
              </div>

              <Button
                onClick={handleVoiceClone}
                disabled={loading}
                className="w-full bg-purple-600 hover:bg-purple-700"
              >
                <Mic className="w-4 h-4 mr-2" />
                Synthesize Speech
              </Button>

              {voiceCloneJobId && (
                <div className={`p-3 border rounded-md text-sm ${
                  voiceCloneJobStatus === "completed"
                    ? "bg-green-50 border-green-200 text-green-700"
                    : voiceCloneJobStatus === "failed" || voiceCloneJobStatus === "timeout"
                    ? "bg-red-50 border-red-200 text-red-700"
                    : "bg-slate-50 border-slate-200 text-slate-700"
                }`}>
                  <div className="flex items-center justify-between">
                    <span>Job ID: #{voiceCloneJobId}</span>
                    <span className="font-semibold">
                      Status: {voiceCloneJobStatus || "pending"}
                    </span>
                  </div>
                  {(voiceCloneJobStatus === "pending" || voiceCloneJobStatus === "processing") && (
                    <p className="text-xs text-slate-500 mt-1">We will notify you when the job completes.</p>
                  )}
                  {voiceCloneJobStatus === "completed" && (
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => downloadVoiceCloneJob(voiceCloneJobId)}
                      className="mt-2"
                    >
                      <Upload className="w-4 h-4 mr-2" />
                      Download
                    </Button>
                  )}
                  {(voiceCloneJobStatus === "failed" || voiceCloneJobStatus === "timeout") && (
                    <Button
                      variant="outline"
                      size="sm"
                      onClick={() => {
                        setVoiceCloneJobId(null);
                        setVoiceCloneJobStatus(null);
                        if (jobPollTimeoutRef.current) {
                          clearTimeout(jobPollTimeoutRef.current);
                        }
                      }}
                      className="mt-2"
                    >
                      Clear
                    </Button>
                  )}
                </div>
              )}

              {clonedAudioUrl && (
                <div className="mt-6 p-4 bg-purple-50 rounded-lg border border-purple-200">
                  <p className="text-xs font-medium text-purple-600 mb-3 uppercase tracking-wider">
                    Synthesized Audio
                  </p>
                  <audio
                    controls
                    src={clonedAudioUrl}
                    className="w-full"
                  />
                </div>
              )}
            </CardContent>
          </Card>

          {/* Notifications */}
          <Card className="border-0 shadow-sm">
            <CardHeader className="border-b border-slate-200 bg-slate-50 flex justify-between items-center">
              <CardTitle className="text-lg">Notifications</CardTitle>
              <div className="flex items-center gap-3 text-sm text-slate-500">
                <span>{unreadNotifications} unread</span>
                <Button variant="outline" size="sm" onClick={markNotificationsRead}>
                  Mark all read
                </Button>
              </div>
            </CardHeader>
            <CardContent className="p-6 space-y-3">
              {notifications.length === 0 ? (
                <p className="text-sm text-slate-500">No notifications yet.</p>
              ) : (
                notifications.slice(0, 5).map((notification) => (
                  <div
                    key={notification.id}
                    className={`p-3 rounded border ${
                      notification.read ? "bg-white border-slate-100" : "bg-slate-50 border-slate-200"
                    }`}
                  >
                    <p className="text-sm text-slate-800">{notification.message}</p>
                    <p className="text-xs text-slate-400 mt-1">
                      {notification.created_at ? new Date(notification.created_at).toLocaleString() : "Just now"}
                    </p>
                  </div>
                ))
              )}
            </CardContent>
          </Card>
        </div>
      </div>
    </div>
  );
}
