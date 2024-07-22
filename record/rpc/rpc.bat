@echo off

goctl rpc protoc record.proto --go_out=. --go-grpc_out=. --zrpc_out=.