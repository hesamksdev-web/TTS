"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";
import { toast } from "sonner";
import { LogOut, Users, ArrowLeft, Upload, Play, Trash2, Plus } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "/api/v1";

interface User {
  id: number;
  email: string;
  role: string;
  approved: boolean;
  created_at: string;
}

interface Dataset {
  name: string;
  audio_count: number;
  created_at: string;
}

interface TrainingFile {
  id: string;
  filename: string;
  transcript: string;
  uploaded: boolean;
}

export default function AdminPanel() {
  const router = useRouter();
  const [token, setToken] = useState<string | null>(null);
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [activeTab, setActiveTab] = useState<"users" | "training">("users");
  
  // Training state
  const [datasets, setDatasets] = useState<Dataset[]>([]);
  const [datasetName, setDatasetName] = useState("");
  const [trainingFiles, setTrainingFiles] = useState<TrainingFile[]>([]);
  const [uploading, setUploading] = useState(false);
  const [training, setTraining] = useState(false);
  const [epochs, setEpochs] = useState(10);
  const [batchSize, setBatchSize] = useState(4);
  const [trainingJobs, setTrainingJobs] = useState<any[]>([]);

  useEffect(() => {
    const t = localStorage.getItem("token");
    const r = localStorage.getItem("role");

    console.log("Admin page - Token:", t ? "exists" : "missing");
    console.log("Admin page - Role:", r);

    if (!t) {
      console.log("No token, redirecting to login");
      router.push("/login");
      return;
    }

    if (r !== "admin") {
      console.log("Not admin, redirecting to dashboard");
      router.push("/");
      return;
    }

    setToken(t);
    fetchUsers(t);
    fetchDatasets(t);
    fetchTrainingJobs(t);
  }, [router]);

  // Refresh training jobs every 5 seconds
  useEffect(() => {
    if (!token || activeTab !== "training") return;
    
    const interval = setInterval(() => {
      fetchTrainingJobs(token);
    }, 5000);

    return () => clearInterval(interval);
  }, [token, activeTab]);

  const fetchUsers = async (authToken: string) => {
    try {
      const res = await axios.get(`${API_BASE_URL}/admin/users`, {
        headers: { Authorization: `Bearer ${authToken}` },
      });
      console.log("Users response:", res.data);
      const userData = Array.isArray(res.data) ? res.data : [];
      setUsers(userData);
    } catch (error: any) {
      console.error("Error fetching users:", error);
      toast.error("Failed to load users");
    } finally {
      setLoading(false);
    }
  };

  const fetchDatasets = async (authToken: string) => {
    try {
      const res = await axios.get(`${API_BASE_URL}/admin/training/datasets`, {
        headers: { Authorization: `Bearer ${authToken}` },
      });
      const datasetList = res.data.datasets || [];
      setDatasets(datasetList);
    } catch (error: any) {
      console.error("Error fetching datasets:", error);
    }
  };

  const fetchTrainingJobs = async (authToken: string) => {
    try {
      const res = await axios.get(`${API_BASE_URL}/admin/training/jobs`, {
        headers: { Authorization: `Bearer ${authToken}` },
      });
      const jobs = res.data.jobs || [];
      setTrainingJobs(jobs);
    } catch (error: any) {
      console.error("Error fetching training jobs:", error);
    }
  };

  const approveUser = async (userId: number) => {
    try {
      await axios.post(
        `${API_BASE_URL}/admin/approve`,
        { user_id: userId },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      toast.success("User approved!");
      fetchUsers(token!);
    } catch (error: any) {
      toast.error("Failed to approve user");
      console.error(error);
    }
  };

  const logout = () => {
    localStorage.clear();
    router.push("/login");
  };

  const handleAddTrainingFile = () => {
    const newFile: TrainingFile = {
      id: Date.now().toString(),
      filename: "",
      transcript: "",
      uploaded: false,
    };
    setTrainingFiles([...trainingFiles, newFile]);
  };

  const handleRemoveTrainingFile = (id: string) => {
    setTrainingFiles(trainingFiles.filter((f) => f.id !== id));
  };

  const handleFileChange = (id: string, filename: string) => {
    setTrainingFiles(
      trainingFiles.map((f) =>
        f.id === id ? { ...f, filename } : f
      )
    );
  };

  const handleTranscriptChange = (id: string, transcript: string) => {
    setTrainingFiles(
      trainingFiles.map((f) =>
        f.id === id ? { ...f, transcript } : f
      )
    );
  };

  const handleUploadTrainingFile = async (file: TrainingFile) => {
    if (!file.filename || !file.transcript || !datasetName) {
      toast.error("Please fill in all fields");
      return;
    }

    setUploading(true);
    try {
      const fileInput = document.getElementById(`file-${file.id}`) as HTMLInputElement;
      if (!fileInput?.files?.[0]) {
        toast.error("Please select a file");
        setUploading(false);
        return;
      }

      const formData = new FormData();
      formData.append("dataset_name", datasetName);
      formData.append("transcript", file.transcript);
      formData.append("audio", fileInput.files[0]);

      await axios.post(
        `${API_BASE_URL}/admin/training/upload`,
        formData,
        {
          headers: {
            Authorization: `Bearer ${token}`,
            "Content-Type": "multipart/form-data",
          },
        }
      );

      toast.success(`File uploaded: ${file.filename}`);
      setTrainingFiles(
        trainingFiles.map((f) =>
          f.id === file.id ? { ...f, uploaded: true } : f
        )
      );
      fetchDatasets(token!);
    } catch (error: any) {
      console.error("Upload error:", error);
      toast.error(error.response?.data?.error || "Upload failed");
    } finally {
      setUploading(false);
    }
  };

  const handleStartTraining = async () => {
    if (!datasetName || epochs <= 0 || batchSize <= 0) {
      toast.error("Please fill in all training parameters");
      return;
    }

    setTraining(true);
    try {
      const res = await axios.post(
        `${API_BASE_URL}/admin/training/start`,
        {
          dataset_name: datasetName,
          epochs,
          batch_size: batchSize,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      toast.success(`Training started! Job ID: ${res.data.job_id}`);
      setTrainingFiles([]);
      setDatasetName("");
      setEpochs(10);
      setBatchSize(4);
      fetchDatasets(token!);
    } catch (error: any) {
      console.error("Training error:", error);
      toast.error(error.response?.data?.error || "Failed to start training");
    } finally {
      setTraining(false);
    }
  };

  if (!token) return null;

  const adminCount = users.filter((u) => u.role === "admin").length;
  const userCount = users.filter((u) => u.role === "user").length;

  return (
    <div className="min-h-screen bg-white">
      {/* Header */}
      <div className="border-b border-slate-200 bg-slate-50">
        <div className="max-w-7xl mx-auto px-6 py-6 flex justify-between items-center">
          <div>
            <h1 className="text-3xl font-bold text-slate-900">Admin Dashboard</h1>
            <p className="text-slate-500 text-sm mt-1">Manage users and model training</p>
          </div>
          <div className="flex gap-3">
            <Button
              variant="outline"
              onClick={() => router.push("/")}
              className="gap-2"
            >
              <ArrowLeft className="w-4 h-4" />
              Back
            </Button>
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

      {/* Tabs */}
      <div className="border-b border-slate-200 bg-white">
        <div className="max-w-7xl mx-auto px-6">
          <div className="flex gap-8">
            <button
              onClick={() => setActiveTab("users")}
              className={`py-4 px-2 font-medium text-sm border-b-2 transition-colors ${
                activeTab === "users"
                  ? "border-blue-600 text-blue-600"
                  : "border-transparent text-slate-600 hover:text-slate-900"
              }`}
            >
              <Users className="w-4 h-4 inline mr-2" />
              Users
            </button>
            <button
              onClick={() => setActiveTab("training")}
              className={`py-4 px-2 font-medium text-sm border-b-2 transition-colors ${
                activeTab === "training"
                  ? "border-blue-600 text-blue-600"
                  : "border-transparent text-slate-600 hover:text-slate-900"
              }`}
            >
              <Upload className="w-4 h-4 inline mr-2" />
              Training
            </button>
          </div>
        </div>
      </div>

      {/* Main Content */}
      <div className="max-w-7xl mx-auto px-6 py-8">
        {activeTab === "users" ? (
          <>
            {/* Stats */}
            <div className="grid grid-cols-3 gap-4 mb-8">
              <Card className="border-0 shadow-sm">
                <CardContent className="pt-6">
                  <div className="text-center">
                    <div className="text-4xl font-bold text-slate-900">
                      {users.length}
                    </div>
                    <p className="text-slate-500 text-sm mt-2">Total Users</p>
                  </div>
                </CardContent>
              </Card>

              <Card className="border-0 shadow-sm">
                <CardContent className="pt-6">
                  <div className="text-center">
                    <div className="text-4xl font-bold text-red-600">
                      {adminCount}
                    </div>
                    <p className="text-slate-500 text-sm mt-2">Administrators</p>
                  </div>
                </CardContent>
              </Card>

              <Card className="border-0 shadow-sm">
                <CardContent className="pt-6">
                  <div className="text-center">
                    <div className="text-4xl font-bold text-blue-600">
                      {userCount}
                    </div>
                    <p className="text-slate-500 text-sm mt-2">Regular Users</p>
                  </div>
                </CardContent>
              </Card>
            </div>

            {/* Users Table */}
            <Card className="border-0 shadow-sm">
              <CardHeader className="border-b border-slate-200 bg-slate-50">
                <CardTitle className="flex items-center gap-2 text-lg">
                  <Users className="w-5 h-5" />
                  Users
                </CardTitle>
              </CardHeader>
              <CardContent className="p-0">
                {loading ? (
                  <div className="text-center py-12 text-slate-500">Loading...</div>
                ) : users.length === 0 ? (
                  <div className="text-center py-12 text-slate-500">
                    No users found
                  </div>
                ) : (
                  <div className="overflow-x-auto">
                    <table className="w-full text-sm">
                      <thead className="border-b border-slate-200 bg-slate-50">
                        <tr>
                          <th className="text-left py-4 px-6 font-semibold text-slate-700">
                            Email
                          </th>
                          <th className="text-left py-4 px-6 font-semibold text-slate-700">
                            Role
                          </th>
                          <th className="text-left py-4 px-6 font-semibold text-slate-700">
                            Status
                          </th>
                          <th className="text-left py-4 px-6 font-semibold text-slate-700">
                            Joined
                          </th>
                          <th className="text-left py-4 px-6 font-semibold text-slate-700">
                            Action
                          </th>
                        </tr>
                      </thead>
                      <tbody>
                        {users.map((user: User) => (
                          <tr
                            key={`user-${user.id}`}
                            className="border-b border-slate-100 hover:bg-slate-50 transition-colors"
                          >
                            <td className="py-4 px-6 font-medium text-slate-900">
                              {user.email}
                            </td>
                            <td className="py-4 px-6">
                              <span
                                className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                                  user.role === "admin"
                                    ? "bg-red-100 text-red-800"
                                    : "bg-blue-100 text-blue-800"
                                }`}
                              >
                                {user.role === "admin" ? "Admin" : "User"}
                              </span>
                            </td>
                            <td className="py-4 px-6">
                              <span
                                className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${
                                  user.approved
                                    ? "bg-green-100 text-green-800"
                                    : "bg-yellow-100 text-yellow-800"
                                }`}
                              >
                                {user.approved ? "Approved" : "Pending"}
                              </span>
                            </td>
                            <td className="py-4 px-6 text-slate-500 text-sm">
                              {user.created_at}
                            </td>
                            <td className="py-4 px-6">
                              {!user.approved && user.role !== "admin" && (
                                <Button
                                  size="sm"
                                  onClick={() => approveUser(user.id)}
                                  className="bg-green-600 hover:bg-green-700 text-white text-xs"
                                >
                                  Approve
                                </Button>
                              )}
                            </td>
                          </tr>
                        ))}
                      </tbody>
                    </table>
                  </div>
                )}
              </CardContent>
            </Card>
          </>
        ) : (
          <>
            {/* Training Section */}
            <div className="grid grid-cols-3 gap-6">
              {/* Upload Section */}
              <Card className="border-0 shadow-sm">
                <CardHeader className="border-b border-slate-200 bg-slate-50">
                  <CardTitle className="text-lg">Upload Training Data</CardTitle>
                </CardHeader>
                <CardContent className="pt-6 space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-slate-700 mb-2">
                      Dataset Name
                    </label>
                    <Input
                      placeholder="e.g., my_voice_dataset"
                      value={datasetName}
                      onChange={(e) => setDatasetName(e.target.value)}
                      className="w-full"
                    />
                  </div>

                  <div className="space-y-3 max-h-96 overflow-y-auto">
                    {trainingFiles.map((file) => (
                      <div
                        key={file.id}
                        className="p-3 border border-slate-200 rounded-lg bg-slate-50"
                      >
                        <div className="space-y-2">
                          <div className="flex gap-2">
                            <Input
                              type="file"
                              id={`file-${file.id}`}
                              accept=".mp3,.wav,.m4a"
                              onChange={(e) =>
                                handleFileChange(file.id, e.target.files?.[0]?.name || "")
                              }
                              className="flex-1 text-xs"
                              disabled={file.uploaded}
                            />
                            <Button
                              size="sm"
                              variant="ghost"
                              onClick={() => handleRemoveTrainingFile(file.id)}
                              className="text-red-600 hover:text-red-700"
                            >
                              <Trash2 className="w-4 h-4" />
                            </Button>
                          </div>
                          <Input
                            placeholder="Transcript text for this audio"
                            value={file.transcript}
                            onChange={(e) =>
                              handleTranscriptChange(file.id, e.target.value)
                            }
                            className="w-full text-sm"
                            disabled={file.uploaded}
                          />
                          <Button
                            size="sm"
                            onClick={() => handleUploadTrainingFile(file)}
                            disabled={uploading || file.uploaded}
                            className="w-full bg-blue-600 hover:bg-blue-700 text-white"
                          >
                            {file.uploaded ? (
                              <>
                                <Play className="w-3 h-3 mr-1" />
                                Uploaded
                              </>
                            ) : (
                              <>
                                <Upload className="w-3 h-3 mr-1" />
                                Upload
                              </>
                            )}
                          </Button>
                        </div>
                      </div>
                    ))}
                  </div>

                  <Button
                    onClick={handleAddTrainingFile}
                    variant="outline"
                    className="w-full gap-2"
                  >
                    <Plus className="w-4 h-4" />
                    Add Audio File
                  </Button>
                </CardContent>
              </Card>

              {/* Training Config Section */}
              <Card className="border-0 shadow-sm">
                <CardHeader className="border-b border-slate-200 bg-slate-50">
                  <CardTitle className="text-lg">Training Configuration</CardTitle>
                </CardHeader>
                <CardContent className="pt-6 space-y-4">
                  <div>
                    <label className="block text-sm font-medium text-slate-700 mb-2">
                      Epochs
                    </label>
                    <Input
                      type="number"
                      min="1"
                      max="1000"
                      value={epochs}
                      onChange={(e) => setEpochs(parseInt(e.target.value) || 10)}
                      className="w-full"
                    />
                    <p className="text-xs text-slate-500 mt-1">
                      Number of training iterations (10-100 recommended)
                    </p>
                  </div>

                  <div>
                    <label className="block text-sm font-medium text-slate-700 mb-2">
                      Batch Size
                    </label>
                    <Input
                      type="number"
                      min="1"
                      max="64"
                      value={batchSize}
                      onChange={(e) => setBatchSize(parseInt(e.target.value) || 4)}
                      className="w-full"
                    />
                    <p className="text-xs text-slate-500 mt-1">
                      Samples per batch (4-16 recommended)
                    </p>
                  </div>

                  <Button
                    onClick={handleStartTraining}
                    disabled={training || !datasetName || trainingFiles.length === 0}
                    className="w-full bg-green-600 hover:bg-green-700 text-white"
                  >
                    {training ? "Starting..." : "Start Training"}
                  </Button>
                </CardContent>
              </Card>

              {/* Datasets List */}
              <Card className="border-0 shadow-sm">
                <CardHeader className="border-b border-slate-200 bg-slate-50">
                  <CardTitle className="text-lg">Datasets</CardTitle>
                </CardHeader>
                <CardContent className="pt-6 max-h-96 overflow-y-auto">
                  {datasets.length === 0 ? (
                    <p className="text-sm text-slate-500 text-center py-4">
                      No datasets yet
                    </p>
                  ) : (
                    <div className="space-y-2">
                      {datasets.map((dataset) => (
                        <div
                          key={dataset.name}
                          className="p-3 border border-slate-200 rounded-lg hover:bg-slate-50 transition-colors"
                        >
                          <div className="flex justify-between items-start">
                            <div>
                              <p className="font-medium text-slate-900 text-sm">
                                {dataset.name}
                              </p>
                              <p className="text-xs text-slate-500 mt-1">
                                {dataset.audio_count} audio files
                              </p>
                            </div>
                            <span className="text-xs text-slate-400">
                              {dataset.created_at}
                            </span>
                          </div>
                        </div>
                      ))}
                    </div>
                  )}
                </CardContent>
              </Card>

              {/* Training Jobs */}
              <Card className="border-0 shadow-sm">
                <CardHeader className="border-b border-slate-200 bg-slate-50">
                  <CardTitle className="text-lg">Training Jobs</CardTitle>
                </CardHeader>
                <CardContent className="pt-6 max-h-96 overflow-y-auto">
                  {trainingJobs.length === 0 ? (
                    <p className="text-sm text-slate-500 text-center py-4">
                      No training jobs yet
                    </p>
                  ) : (
                    <div className="space-y-2">
                      {trainingJobs.map((job) => (
                        <div
                          key={job.job_id}
                          className="p-3 border border-slate-200 rounded-lg hover:bg-slate-50 transition-colors"
                        >
                          <div className="flex justify-between items-start mb-2">
                            <div className="flex-1">
                              <p className="font-medium text-slate-900 text-sm break-all">
                                {job.job_id}
                              </p>
                              <p className="text-xs text-slate-500 mt-1">
                                {job.created_at}
                              </p>
                            </div>
                            <span
                              className={`inline-flex items-center px-2 py-1 rounded-full text-xs font-medium whitespace-nowrap ml-2 ${
                                job.status === "training"
                                  ? "bg-blue-100 text-blue-800"
                                  : job.has_model
                                  ? "bg-green-100 text-green-800"
                                  : "bg-yellow-100 text-yellow-800"
                              }`}
                            >
                              {job.status === "training"
                                ? "Training..."
                                : job.has_model
                                ? "Completed"
                                : "Pending"}
                            </span>
                          </div>
                        </div>
                      ))}
                    </div>
                  )}
                </CardContent>
              </Card>
            </div>
          </>
        )}
      </div>
    </div>
  );
}
