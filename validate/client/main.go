package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	pb "github.com/apus-run/gaia/examples/validate/api"
	"github.com/apus-run/gaia/middleware/recovery"
	"github.com/apus-run/gaia/transport/grpc"
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

	client := pb.NewExampleServiceClient(conn)

	reply, err := client.TestValidate(context.Background(), &pb.Request{
		Id:      10001,
		Age:     18,
		Code:    1,
		Score:   0.55,
		State:   true,
		Path:    "/hello",
		Phone:   "13484903846",
		Explain: "sss",
		Name:    "ok",
		Card:    "111",
		Info:    &pb.Info{Address: "鼓楼"},
	})

	e, ok := status.FromError(err)
	if ok && e.Code() == codes.InvalidArgument {
		log.Printf("参数错误: %+v\n", e.Message())
	}

	log.Printf("[grpc] TestValidate %+v\n", reply)
}
