#!/usr/bin/env bash
docker build --target crawlerbuild -t crawlerbuild . -f build/Dockerfile &&
    docker run crawlerbuild go test -v -coverprofile cover.out ./... &&
    go test -v ./...
