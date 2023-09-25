package main

import (
	"grpcservice/protos"
	chat "grpcservice/protos"
	"log"

	"golang.org/x/net/context"
)

func helloTest(c chat.ChatServiceClient) {
	response, err := c.SayHello(context.Background(), &chat.Message{Body: "Hello From Client!"})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}

func retrieveMessages(c chat.ChatServiceClient) *protos.Messages {
	response, err := c.GetMessages(context.Background(), &chat.Empty{})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Println("Response from server: ", response.Messages)

	return response
}

func sendMessage(c chat.ChatServiceClient, body string, user string) {
	response, err := c.SendMessage(context.Background(), &chat.Message{Body: body, User: user})

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Println("Response from server: ", response.Status)
}
