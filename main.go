package main

import (
	"fmt"
	"log"
	"net"

	protos "grpcservice/protos"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("○ Go gRPC Search Service ○")
	fmt.Println("// Starting server at ", tcpPort)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", tcpPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := protos.Server{}
	grpcServer := grpc.NewServer()

	protos.RegisterChatServiceServer(grpcServer, &server)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
