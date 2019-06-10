#!/usr/bin/env bash
echo 'curl --data-urlencode "url=https://www.goncalopereira.com" http://localhost:8080`'
go build -o ./bin/crawler cmd/crawler/main.go
./bin/crawler
