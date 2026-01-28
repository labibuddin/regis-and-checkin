#!/bin/bash

# Function to kill child processes on exit
trap 'kill $(jobs -p)' EXIT

echo "🚀 Starting Database (Docker)..."
docker compose up -d db

echo "⏳ Waiting for Database to be ready..."
sleep 5 # Simple wait, could be more robust with pg_isready

echo "🚀 Starting Backend (Go Fiber)..."
cd backend
# Check if .env exists
if [ ! -f .env ]; then
  echo "❌ .env file missing in backend!"
  exit 1
fi
go run main.go &
BACKEND_PID=$!
cd ..

echo "🚀 Starting Frontend (Vite)..."
cd frontend
npm run dev &
FRONTEND_PID=$!
cd ..

echo "✅ All services started!"
echo "Backend: http://localhost:8080"
echo "Frontend: http://localhost:1966"
echo "Database: localhost:5432"
echo ""
echo "Press Ctrl+C to stop everything."

# Wait for both processes
wait $BACKEND_PID $FRONTEND_PID
