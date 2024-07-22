@echo off

goctl rpc protoc live.proto --go_out=. --go-grpc_out=. --zrpc_out=.