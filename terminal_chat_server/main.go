package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/40grivenprog/terminal_chat/pb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedTerminalChatServer
}

func (s *server) ProcessMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	message := fmt.Sprintf("Received: %v From: %v", in.GetMessage(), in.GetUsername())
	log.Printf(message)

	return &pb.MessageReply{Message: message}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTerminalChatServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
