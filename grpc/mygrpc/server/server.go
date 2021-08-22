package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/lirlia/go-mypj/grpc/mygrpc/proto"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type calcServer struct {
	pb.UnimplementedCalcServer
}

func (s *calcServer) Add(ctx context.Context, in *pb.CalcReq) (*pb.CalcRes, error) {
	nums := in.GetParams()
	var total int64
	for _, n := range nums {
		total = n + total
	}
	fmt.Println("add  %v %n", nums, total)
	return &pb.CalcRes{Params: total}, nil
}

func (s *calcServer) Minus(ctx context.Context, in *pb.CalcReq) (*pb.CalcRes, error) {
	nums := in.GetParams()
	var total int64 = nums[0] * 2
	for _, n := range nums {
		total = total - n
	}
	fmt.Println("add  %v %n", nums, total)
	return &pb.CalcRes{Params: total}, nil
}

func (s *calcServer) Product(ctx context.Context, in *pb.CalcReq) (*pb.CalcRes, error) {
	nums := in.GetParams()
	var total int64 = 1
	for _, n := range nums {
		total = total * n
	}
	return &pb.CalcRes{Params: total}, nil
}

func (s *calcServer) Div(ctx context.Context, in *pb.CalcReq) (*pb.CalcRes, error) {
	nums := in.GetParams()
	var total int64 = nums[0] * nums[0]
	for _, n := range nums {
		total = total / n
	}
	return &pb.CalcRes{Params: total}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalcServer(s, &calcServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
