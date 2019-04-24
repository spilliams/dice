#!/usr/bin/env bash
: ${SPLAT:=./...}
go test -timeout 30s -p 1 $SPLAT -coverpkg=$SPLAT -coverprofile .testCoverage.txt
go tool cover -html=.testCoverage.txt -o .testCoverage.html
