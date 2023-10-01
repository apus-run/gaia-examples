package main

import (
	"context"

	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/transport/grpc"
	"github.com/apus-run/sea-kit/log"

	pb "github.com/apus-run/gaia/examples/helloworld/api"

	_ "go.uber.org/automaxprocs"
)

func main() {
	runGrpc()
}

func runGrpc() {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("127.0.0.1:9000"),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	reply, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "gaia"})
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("[grpc] SayHello %+v\n", reply)

	// returns error
	_, err = client.SayHello(context.Background(), &pb.HelloRequest{Name: "error"})
	if err != nil {
		log.Infof("[grpc] SayHello error: %v\n", err)
	}
}
