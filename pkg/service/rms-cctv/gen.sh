#!/bin/bash

PATH="$PATH:$GOBIN"
protoc --proto_path=../../../api/:. --micro_out=. --go_out=. rms-cctv.proto