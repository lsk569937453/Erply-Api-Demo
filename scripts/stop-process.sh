#!/bin/bash

ps -ef | grep erply-api | grep -v grep | awk '{print $2}' | xargs kill
sudo rm -rf /tmp/erply_api/erply*
sudo rm -rf /tmp/erply_api/scripts
sudo rm -rf /tmp/erply_api/appspec.yml
sudo rm -rf /opt/codedeploy-agent/deployment-root/*