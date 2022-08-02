package api

//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./example.proto
//go:generate protoc --proto_path=. --proto_path=../../third_party --go_out=paths=source_relative:. --validate_out=paths=source_relative,lang=go:. ./example.proto
