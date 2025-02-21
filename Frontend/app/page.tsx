"use client";

import { useEffect, useState } from "react";
import TaskDashboard from "../components/TaskDashboard";
import useTaskUpdates from "../hooks/useTaskUpdates";

export default function Home() {
  const [tasks, setTasks] = useState<any[]>([]);
  const updates = useTaskUpdates();

  // Fetch tasks from the backend API
  useEffect(() => {
    fetch("http://localhost:8000/tasks", {
      headers: {
        Authorization: "Bearer " + localStorage.getItem("token"),
      },
    })
      .then((res) => res.json())
      .then((data) => setTasks(data));
  }, []);

  // Append real-time updates from WebSocket
  useEffect(() => {
    if (updates.length > 0) {
      setTasks((prev) => [...prev, updates[updates.length - 1]]);
    }
  }, [updates]);

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Task Dashboard</h1>
      <TaskDashboard tasks={tasks} />
    </div>
  );
}
