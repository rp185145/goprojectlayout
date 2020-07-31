#!/bin/bash

set -e
set -x

go test -race -coverprofile=coverage.txt . ./emitter/... ./idutils/...
go tool cover -html=coverage.txt -o coverage.html

for cmd in hello howdy uuid demo gplcli
do
    go build -o "bin/$cmd" "./cmd/$cmd/..."
    "bin/$cmd"
done

echo SUCCESS