#! /bin/bash

go build shprotobot.go
ssh -t berlin@95.85.49.5 "sudo service shprotobot stop"
scp ./shprotobot berlin@95.85.49.5:/home/berlin/shprotobot/
scp ./config.json berlin@95.85.49.5:/home/berlin/shprotobot/config.json
ssh -t berlin@95.85.49.5 "sudo service shprotobot start"
rm shprotobot
