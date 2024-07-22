@echo off

goctl rpc protoc voice.proto --go_out=. --go-grpc_out=. --zrpc_out=.