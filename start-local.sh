#!/bin/bash

cd frontend
nohup pnpm run dev > logs/frontend.log 2>&1 &
FRONTEND_PID=$!
firefox --new-tab "localhost:5174"
echo "frontend running"
cd ..

read exit

kill $FRONTEND_PID
echo "frontend killed"