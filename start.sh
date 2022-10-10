#!/bin/sh

rm start-server
go build cmd/start-server.go
./start-server
