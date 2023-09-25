package main

import (
	chat "grpcservice/protos"
	"log"

	"google.golang.org/grpc"
)

func initGrpc() (chat.ChatServiceClient, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("[gRPC] Not connected.: %s", err)
	}

	c := chat.NewChatServiceClient(conn)
	return c, conn
}
