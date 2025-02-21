import { useEffect, useState } from "react";

export default function useTaskUpdates() {
  const [updates, setUpdates] = useState<any[]>([]);

  useEffect(() => {
    const ws = new WebSocket("ws://localhost:8000/ws/tasks");
    ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      setUpdates((prev) => [...prev, data]);
    };
    ws.onerror = (err) => {
      console.error("WebSocket error:", err);
    };
    return () => {
      ws.close();
    };
  }, []);

  return updates;
}
