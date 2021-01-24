#!/bin/sh

protoc --descriptor_set_out=./proto/api_descriptor.pb --go_out=. --go-grpc_out=. --go_opt=module=github.com/Urotea/cloud-endpoints-sample --go-grpc_opt=module=github.com/Urotea/cloud-endpoints-sample ./proto/*.proto
