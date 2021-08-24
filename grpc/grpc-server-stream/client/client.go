package main

import (
	"context"
	"fmt"
	"io"

	pb "github.com/lirlia/go-mypj/grpc/grpc-server-stream/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func request(client pb.NotificationClient, num int32) error {
	req := &pb.NotificationReq{
		Num: num,
	}

	stream, err := client.Notification(context.Background(), req)
	if err != nil {
		return errors.Wrap(err, "stream error")
	}

	for {
		reply, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		fmt.Println("これ:", reply.GetMessage())
	}
	return nil
}

func exec(num int32) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return errors.Wrap(err, "connection err")
	}

	defer conn.Close()
	client := pb.NewNotificationClient(conn)
	return request(client, num)
}
func main() {
	num := int32(5)
	if err := exec(num); err != nil {
		fmt.Println(err)
	}
}
