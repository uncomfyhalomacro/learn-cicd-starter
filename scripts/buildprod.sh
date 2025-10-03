#!/bin/bash

if [[ $(arch) == "x86_64" ]]; then
    export GOARCH="amd64"
else
    export GOARCH="arm64"
fi

CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build -o notely
