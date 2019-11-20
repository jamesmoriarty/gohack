#!/bin/sh

set -u
set +x
set +e

VERSION=$(git rev-parse --short HEAD)
DATE=$(date -u +.%Y%m%d.%H%M%S)

go build -v -ldflags "-X github.com/jamesmoriarty/gohack/config.Version=$VERSION -X github.com/jamesmoriarty/gohack/config.Date=$DATE"