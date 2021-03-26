#!/bin/bash

cd /tmp/
unzip erply-api.zip
nohup ./erply-api >>erplyApiGoLog 2>&1 &
