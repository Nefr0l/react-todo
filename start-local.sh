#!/bin/bash

cd frontend
nohup pnpm run dev > logs/frontend.log 2>&1 &
FRONTEND_PID=$!
echo "frontend running"
cd ..

cd backend
go build -o api .
echo "backend running"

read exit

kill $FRONTEND_PID
echo "frontend and backend killed"