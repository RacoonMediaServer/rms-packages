#!/bin/bash

PATH="$PATH:$GOBIN"
protoc --proto_path=../../api/:.  --micro_out=paths=source_relative:. --go_out=paths=source_relative:. events.proto