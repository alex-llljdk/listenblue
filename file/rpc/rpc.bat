@echo off

goctl rpc protoc file.proto --go_out=. --go-grpc_out=. --zrpc_out=.