#!/bin/bash

#Starting mysql and redis services
echo Starting redis and mysql server
sudo mysql.server start
sleep 1
brew services start redis
sleep 1

# Running HTTP and TCP servers with nohup
kill $(lsof -ti:3000,4000,8081)
cd ../servers/TCPServer
nohup go run . &
sleep 1
cd ../HTTPServer
nohup go run . &
sleep 1
echo The TCP and the HTTP Servers have started

cd ../../client
nohup npm start &
sleep 3
echo Client web app has started
