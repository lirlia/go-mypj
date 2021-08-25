package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/lirlia/go-mypj/grpc/grpc-client-stream/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type ServerServerSide struct {
	pb.UnimplementedNotificationServer
}

func (s *ServerServerSide) Notification(stream pb.Notification_NotificationServer) error {
	fmt.Println("receive request")
	req := &pb.NotificationReq{}
	var err error
	var num int32
	for {
		req, err = stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.NotificationRes{
				Num: num,
			})
		}

		if err != nil {
			return err
		}
		num = num + req.GetNum()
		fmt.Println(num)
	}

	return nil
}

func set() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return errors.Wrap(err, "fail to listen port")
	}

	s := grpc.NewServer()
	var server ServerServerSide
	pb.RegisterNotificationServer(s, &server)
	if err := s.Serve(lis); err != nil {
		return errors.Wrap(err, "fail to launch server")
	}
	return nil
}

func main() {
	fmt.Println("launch")
	if err := set(); err != nil {
		log.Fatalf("%v", err)
	}
}
