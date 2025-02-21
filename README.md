# AI-Powered Task Management System

## Overview
A real-time AI-powered task management system that allows users to create, assign, and track tasks. The system features JWT-based authentication, AI-powered task suggestions, and real-time updates via WebSockets.

## Features
- **User Authentication**: JWT-based signup and login.
- **Task Management**: Create, assign, update, and track tasks.
- **AI-Powered Task Suggestions**: Smart breakdowns and recommendations via OpenAI/Gemini API.
- **Real-Time Updates**: WebSocket integration for instant notifications.
- **Deployment**: Backend on Render/Fly.io, Frontend on Vercel.

## Tech Stack
### Backend
- Golang (Gin)
- PostgreSQL/MongoDB
- JWT Authentication
- WebSockets (Goroutines)
- OpenAI/Gemini API

### Frontend
- Next.js (TypeScript)
- Tailwind CSS
- Client-side JWT handling
- WebSockets for real-time updates

### Deployment
- Backend: Render/Fly.io
- Frontend: Vercel

## Installation & Setup
### Prerequisites
- **Backend**: Go 1.21+, PostgreSQL/MongoDB
- **Frontend**: Node.js 16+

### Clone the Repository
```sh
 git clone https://github.com/your-username/ai-task-manager.git
 cd ai-task-manager
```

### Backend Setup
1. Navigate to `backend/`:
   ```sh
   cd backend
   ```
2. Create a `.env` file:
   ```sh
   MONGO_URI=mongodb://your-mongo-uri
   SECRET_KEY=your_jwt_secret_key
   OPENAI_API_KEY=your_openai_api_key
   ```
3. Run the backend:
   ```sh
   go run main.go
   ```

### Frontend Setup
1. Navigate to `frontend/`:
   ```sh
   cd frontend
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Start the development server:
   ```sh
   npm run dev
   ```
4. Open [http://localhost:3000](http://localhost:3000) in your browser.

## Deployment
### Backend
- Deploy using Render/Fly.io.
- Set up environment variables in the respective cloud platform.

### Frontend
- Deploy using Vercel.
- Configure environment variables via Vercel Dashboard.

## AI Utilization
- **GitHub Copilot**: Assisted with code generation.
- **ChatGPT/AutoGPT**: Helped with architecture design and issue resolution.
- **Rapid prototyping**: Enabled fast development within 4 hours.


