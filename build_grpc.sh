#!/bin/sh

protoc --go_out=. --go-grpc_out=. --go_opt=module=github.com/Urotea/cloud-endpoints-sample --go-grpc_opt=module=github.com/Urotea/cloud-endpoints-sample ./proto/*.proto
