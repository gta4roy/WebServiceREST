#!/bin/sh
protoc -I protocol/ protocol/addressservice.proto --go_out=plugins=grpc:protocol
