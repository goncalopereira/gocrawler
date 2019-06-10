#!/usr/bin/env bash
echo 'race flag in use'
echo 'curl --data-urlencode "url=https://www.goncalopereira.com" http://localhost:8080`'
go run -race cmd/crawler/main.go
