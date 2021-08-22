package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/lirlia/go-mypj/grpc/test/myproto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedCloudServer
}

func (s *server) GetData(ctx context.Context, in *pb.DateRequest) (*pb.DateResponse, error) {
	locale := in.getLocale()
	log.Printf("Received: %v", locale)
	now := time.Now()
	nowUTC := now.UTC()

	switch locale {
	case "jst":
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		return &pb.DateResponse{date: nowUTC.In(jst)}, nil
	default:
		return &pb.DateResponse{date: nowUTC}, nil
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
