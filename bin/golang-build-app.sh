#!/bin/sh

#
#if [ "$1" == 'linux' ]
#then
#  env GOOS=linux GARCH=amd64 CGO_ENABLED=0 go install -v -a -installsuffix cgo
#fi

export GOOS=linux
export GARCH=amd64
export CGO_ENABLED=0
go install -v -a -installsuffix cgo
go build -o ../docker/main github.com/scarrionv/simple-api/app/src/
