#!/usr/bin/env bash
echo 'curl --data-urlencode "url=https://www.goncalopereira.com" http://localhost:8080'
docker build -t crawler . -f build/Dockerfile &&
    docker run -p 8080:8080 --env-file ./.env crawler
