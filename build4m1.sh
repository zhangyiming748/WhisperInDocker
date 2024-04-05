#!/usr/bin/env bash

echo 编译二进制文件forM1

CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o srt main.go