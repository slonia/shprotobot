#! /bin/bash

go build shprotobot.go
ssh -t server@ams_server "sudo service shprotobot stop"
scp ./shprotobot server@ams_server:/home/server/apps/shproto_bot/
scp ./config.json server@ams_server:/home/server/apps/shproto_bot/config.json
ssh -t server@ams_server "sudo service shprotobot start"
rm shprotobot
