#!/bin/bash

PATH="$PATH:$GOBIN"
protoc --proto_path=../../api/:.  --go_out=. communication.proto