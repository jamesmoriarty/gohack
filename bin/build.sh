#!/bin/sh

set -u
set +x
set +e

VERSION=$(git rev-parse --short HEAD)
DATE=$(date -u +.%Y%m%d.%H%M%S)

go build -v -ldflags "-X main.version=$VERSION -X main.date=$DATE"
