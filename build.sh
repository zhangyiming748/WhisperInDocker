#!/usr/bin/env bash
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o srt main.go