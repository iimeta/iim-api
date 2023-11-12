#!/bin/bash

docker pull iimeta/iim-api:1.1.0

mkdir -p /data/iim-api/manifest/config

wget -P /data/iim-api/manifest/config https://github.com/iimeta/iim-api/raw/docker/manifest/config/config.yaml
wget https://github.com/iimeta/iim-api/raw/docker/bin/start.sh
