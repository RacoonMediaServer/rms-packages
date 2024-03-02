#!/bin/bash

PATH="$PATH:$GOBIN"
protoc --proto_path=../../api/:.  --go_out=paths=source_relative:. media.proto