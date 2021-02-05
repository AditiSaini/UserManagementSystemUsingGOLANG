#!/bin/bash

#Starting mysql and redis services
echo Stopping redis and mysql server
sudo mysql.server stop
sleep 1
brew services stop redis
sleep 1

# Running HTTP and TCP servers with nohup
kill $(lsof -ti:3000,4000,8081)
echo Closed all services at HTTP and TCP ports