#!/bin/bash

BRANCH=$(git symbolic-ref --short -q HEAD)
COMMIT=$(git rev-parse --verify HEAD)
NOW=$(date '+%FT%T%z')
VERSION="0.1.0"
RUN_USER=$(whoami)

go build -ldflags "-X go-base/cmd.AppName=go-base\
 -X go-base/cmd.Branch=${BRANCH}\
 -X go-base/cmd.Commit=${COMMIT}\
 -X go-base/cmd.Author=${RUN_USER}\
 -X go-base/cmd.Date=${NOW}\
 -X go-base/cmd.Version=${VERSION}"