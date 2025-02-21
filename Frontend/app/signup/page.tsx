"use client";
import { useState } from "react";
import { useRouter } from "next/navigation";

export default function SignupPage() {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [phone, setPhone] = useState("");
  const [userType, setUserType] = useState("USER");
  const router = useRouter();

  const handleSignup = async (e: React.FormEvent) => {
    e.preventDefault();
    const res = await fetch("http://localhost:8000/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        first_name: firstName,
        last_name: lastName,
        email,
        password,
        phone,
        user_type: userType,
      }),
    });
    const data = await res.json();
    if (res.ok) {
      alert("User created successfully. Please log in.");
      router.push("/login");
    } else {
      alert(data.error);
    }
  };

  return (
    <div className="container mx-auto p-4">
      <h1 className="text-2xl font-bold mb-4">Sign Up</h1>
      <form onSubmit={handleSignup} className="space-y-4">
        <input
          type="text"
          placeholder="First Name"
          className="border p-2 w-full"
          value={firstName}
          onChange={(e) => setFirstName(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Last Name"
          className="border p-2 w-full"
          value={lastName}
          onChange={(e) => setLastName(e.target.value)}
          required
        />
        <input
          type="email"
          placeholder="Email"
          className="border p-2 w-full"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <input
          type="password"
          placeholder="Password"
          className="border p-2 w-full"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          required
        />
        <input
          type="text"
          placeholder="Phone"
          className="border p-2 w-full"
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
          required
        />
        <select
          value={userType}
          onChange={(e) => setUserType(e.target.value)}
          className="border p-2 w-full"
        >
          <option value="USER">User</option>
          <option value="ADMIN">Admin</option>
        </select>
        <button className="bg-green-500 text-white px-4 py-2">Sign Up</button>
      </form>
    </div>
  );
}
