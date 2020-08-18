#!/bin/bash

# Exit on non-zero (basically, an error) status
set -e

# Trace commands (echo command to stdout)
set -x

# Enforce the gofmt style on all code. At its most
# basic, these commands force a non-zero status if
# gofmt emits output, meaning it had to modify code
# to meet the standard. This is how it's done from
# a POSIX shell. Not sure how Windws PowerShell or
# cmd.exe would do this.
diff -u <(echo -n) <(gofmt -d -s .)

# Check code for suspicious constructs using go vet
go vet ./...

# Execute unit tests with:
#   - Race detector enabled
#   - Output coverage info to coverage.txt
go test -race -coverprofile=coverage.txt . ./emitter/... ./idutils/...

# Generate html report for unit test code
# coverage to coverage.html
go tool cover -html=coverage.txt -o coverage.html

# Build and execute every utility. If the build or
# execute generates a non-zero status, this script
# will end.
#
# Targeting other operating systems requires setting
# the GOOS and GOARCH environment variables. PowerShell
# and cmd.exe environment variables are set differently
# than this POSIX example.
#
# The use of GGO_ENABLED for the Linux compile is to
# disable the use of cgo so Linux builds can run on
# Alpine Linux which is often used as a base for
# Docker images.
for cmd in hello howdy uuid demo gplcli
do
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "bin/$cmd" "./cmd/$cmd/..."
    GOOS=windows GOARCH=amd64 go build -o "bin/$cmd.exe" "./cmd/$cmd/..."
    "bin/$cmd"
done

echo SUCCESS