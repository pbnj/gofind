#!/bin/bash

version=${1}

go clean
# Build: Linux
GOOS=linux go build .
mv gofind releases/gofind-${version}-linux
# Build: Darwin
go build .
mv gofind releases/gofind-${version}-darwin

# Install: Darwin
# go install
