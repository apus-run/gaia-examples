package api

//go:generate protoc -I . -I ../../third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./user.proto
