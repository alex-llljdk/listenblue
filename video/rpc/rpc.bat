@echo off

goctl rpc protoc video.proto --go_out=. --go-grpc_out=. --zrpc_out=.