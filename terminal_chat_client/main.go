package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"

	pb "github.com/40grivenprog/terminal_chat/pb"
	"github.com/40grivenprog/terminal_chat/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTerminalChatClient(conn)

	reader := bufio.NewReader(os.Stdin)
	log.Println("Simple Shell")
	log.Println("---------------------")
	username := util.RandomOwner()

	for {
		log.Printf("%s-->", username)
		text, _ := reader.ReadString('\n')

		r, err := c.ProcessMessage(context.Background(), &pb.MessageRequest{Username: username, Message: text})
		if err != nil {
			log.Fatalf("could not process message: %v", err)
		}

		log.Printf("SERVER REPLY: %s", r.GetMessage())
	}
}
