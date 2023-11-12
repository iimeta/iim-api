#!/bin/bash

docker run -d \
  --network host \
  --restart=always \
  -p 11000:11000 \
  -v /etc/localtime:/etc/localtime:ro \
  -v /data/iim-api/manifest/config/config.yaml:/app/manifest/config/config.yaml \
  --name iim-api \
  iimeta/iim-api:1.1.0
