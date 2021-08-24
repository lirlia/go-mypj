package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/lirlia/go-mypj/grpc/grpc-server-stream/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type ServerServerSide struct {
	pb.UnimplementedNotificationServer
}

func (s *ServerServerSide) Notification(req *pb.NotificationReq, stream pb.Notification_NotificationServer) error {
	fmt.Println("receive request")

	for i := int32(0); i < req.GetNum(); i++ {
		message := fmt.Sprintf("%d", i)
		if err := stream.Send(&pb.NotificationRes{
			Message: message,
		}); err != nil {
			return err
		}
		time.Sleep(time.Second * 1)
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
