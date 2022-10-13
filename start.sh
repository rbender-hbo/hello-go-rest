#!/bin/sh

#TODO Replace with Makefile

rm start-server
go build cmd/start-server.go
./start-server
