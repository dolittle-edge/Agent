#!/bin/bash
env GOOS=linux GOARCH=amd64 GOARM=7
go build -buildmode=exe -o output/DolittleEdgeAgent *.go