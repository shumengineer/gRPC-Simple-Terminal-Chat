package protos

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/context"
)

type Server struct {
	UnimplementedChatServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) SendMessage(ctx context.Context, in *Message) (*Status, error) {
	log.Printf("[SendMessage] Receive: %s", in.Body)
	m, err := NewMessage()

	if err != nil {
		return nil, err
	}

	id, err := m.Insert(in)

	if err != nil {
		log.Println("ERROR [SendMessage]: ", err.Error())
		return nil, err
	}

	log.Println("+ Message. id: ", id)
	return &Status{Status: 1}, nil
}

func (s *Server) GetMessages(ctx context.Context, in *Empty) (*Messages, error) {
	log.Printf("Getting messages")
	m, _ := NewMessage()

	res, _ := m.List(200)

	return &Messages{Messages: res}, nil
}
