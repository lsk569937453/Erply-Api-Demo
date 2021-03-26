#!/bin/bash

cd /tmp/erply_api
unzip erply-api.zip
nohup ./erply-api >>erplyApiGoLog 2>&1 &
