#!/bin/bash

ps -ef | grep erply-api | grep -v grep | awk '{print $2}' | xargs kill
sudo rm -rf /tmp/erply-api/*
sudo rm -rf /opt/codedeploy-agent/deployment-root/*