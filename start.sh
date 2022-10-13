#!/bin/sh

#TODO Replace with Makefile

rm start-server
go test ./...
go build cmd/start-server.go
./start-server
