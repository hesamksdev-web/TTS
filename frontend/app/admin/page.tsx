"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import axios from "axios";
import { toast } from "sonner";
import { LogOut, Users, ArrowLeft } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "/api/v1";

interface User {
  id: number;
  email: string;
  role: string;
  approved: boolean;
  created_at: string;
}

export default function AdminPanel() {
  const router = useRouter();
  const [token, setToken] = useState<string | null>(null);
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);

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
  }, [router]);

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
            <p className="text-slate-500 text-sm mt-1">Manage users and system settings</p>
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

      {/* Main Content */}
      <div className="max-w-7xl mx-auto px-6 py-8">
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
      </div>
    </div>
  );
}
