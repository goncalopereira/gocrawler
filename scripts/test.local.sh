#!/usr/bin/env bash
cp ./scripts/cover.out cover.out
go test -v -coverprofile cover.out ./...
go tool cover -html=cover.out -o ./docs/index.html
