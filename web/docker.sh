#!/bin/bash
if [ -d ./dist ]; then rm -r ./dist; fi
mkdir -p ./dist
docker build -t ricardobalk/go-osmand-tracker .
docker run --rm --user "$(id -u):$(id -g)" \
  --mount type=bind,source="$(pwd)"/src/,target=/home/node/go-osmand-tracker/src/,readonly \
  --mount type=bind,source="$(pwd)"/public/,target=/home/node/go-osmand-tracker/public/,readonly \
  --mount type=bind,source="$(pwd)"/.env.local,target=/home/node/go-osmand-tracker/.env.local,readonly \
  --mount type=bind,source="$(pwd)"/dist/,target=/home/node/go-osmand-tracker/dist/ \
  ricardobalk/go-osmand-tracker "build"
