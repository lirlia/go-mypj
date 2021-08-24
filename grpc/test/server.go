package main

import (
	"context"
	"log"
	"net"

	pb "github.com/lirlia/go-mypj/grpc/test/myproto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCloudServer
}

func (s *server) GetDate(ctx context.Context, in *pb.DateRequest) (*pb.DateReply, error) {
	locale := in.Locale
	log.Printf("Received: %v", locale)
	return &pb.DateReply{Locale: locale}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCloudServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
