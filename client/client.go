package main

import (
	"bufio"
	"fmt"
	chat "grpcservice/protos"
	"os"
	"regexp"
	"strings"
	"time"
)

func input(prompt string) string {
	fmt.Print(prompt)

	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)

	is_alphanumeric := regexp.MustCompile(`^[a-zA-Z0-9 ]*$`).MatchString(input)

	if !is_alphanumeric || len(input) < 1 {
		panic("Wrong input")
	}

	return input
}

func inputLoop(client chat.ChatServiceClient, name string) {
	messages := retrieveMessages(client)
	fmt.Printf("\x1b[2J") // safe alternative for terminal
	fmt.Println(" ----- gRPC Chat Client -----")
	fmt.Println("Rules: Only English, without special characters.")
	fmt.Println("-> Latest messages: ")

	for _, m := range messages.GetMessages() {
		fmt.Printf("%v: %v\n", m.User, m.Body)
	}

	body := input("\nEnter Message: ")

	sendMessage(client, body, name)
}

func main() {
	client, conn := initGrpc()
	defer conn.Close()

	var name string

	if !isProfileExists() {
		name = input("\nYour Name: ")
		initProfile(Profile{Name: name})
	} else {
		name = getProfile().Name
		fmt.Printf("Welcome back, %v!\n\n", name)
		time.Sleep(3 * time.Second)
	}

	for {
		inputLoop(client, name)
	}

	// helloTest(client)
}
