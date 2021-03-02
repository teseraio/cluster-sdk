SHELL := /bin/bash

protoc:
	protoc --go_out=plugins=grpc:. ./cluster/proto/*.proto
