package proto

//go:generate protoc -I . -I ../../../third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./user.proto

//go:generate protoc -I . -I ../../../third_party --go_out=paths=source_relative:. --go-gin_out=paths=source_relative:. ./user.proto

//go:generate protoc-go-inject-tag -input=*.pb.go

//go:generate protoc -I . -I ../../../third_party --go_out=paths=source_relative:. --openapi_out=paths=source_relative:. ./user.proto
