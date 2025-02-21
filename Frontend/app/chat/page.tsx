"use client";
import { useState } from "react";

export default function ChatPage() {
  const [messages, setMessages] = useState<{ role: string; content: string }[]>([]);
  const [input, setInput] = useState("");

  const sendMessage = async (e: React.FormEvent) => {
    e.preventDefault();
    const userMessage = { role: "user", content: input };
    setMessages((prev) => [...prev, userMessage]);
    setInput("");

    // Call backend AI endpoint for task suggestions
    const res = await fetch("http://localhost:8000/tasks/suggestions", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: "Bearer " + localStorage.getItem("token"),
      },
      body: JSON.stringify({ description: input }),
    });
    const data = await res.json();
    const aiMessage = { role: "assistant", content: data.suggestion };
    setMessages((prev) => [...prev, aiMessage]);
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">AI Task Chat</h1>
      <div className="border p-4 h-64 overflow-y-scroll mb-4">
        {messages.map((msg, i) => (
          <div key={i} className={msg.role === "user" ? "text-right" : "text-left"}>
            <p>{msg.content}</p>
          </div>
        ))}
      </div>
      <form onSubmit={sendMessage} className="flex">
        <input
          type="text"
          className="border flex-grow p-2"
          value={input}
          onChange={(e) => setInput(e.target.value)}
          placeholder="Ask for task suggestions..."
        />
        <button type="submit" className="bg-blue-500 text-white px-4 py-2 ml-2">
          Send
        </button>
      </form>
    </div>
  );
}
