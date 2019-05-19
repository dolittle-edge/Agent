#!/bin/bash
#env GOOS=linux GOARCH=amd64 GOARM=7
#go build -buildmode=exe -o output/DolittleEdgeAgent *.go
gox -osarch="linux/amd64" -ldflags="-s -w" -gcflags="-s -w" -output="output/DolittleEdgeAgent"