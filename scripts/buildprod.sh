#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=$(arch) go build -o notely
