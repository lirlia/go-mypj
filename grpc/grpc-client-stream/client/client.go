package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"

	pb "github.com/lirlia/go-mypj/grpc/grpc-client-stream/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	defer conn.Close()
	client := pb.NewNotificationClient(conn)
	stream, _ := client.Notification(context.Background())
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		if scanner.Text() == "exit" {
			res := &pb.NotificationRes{}
			res, _ = stream.CloseAndRecv()

			fmt.Println(res.GetNum())
			break
		}
		i, _ := strconv.Atoi(scanner.Text())
		req := &pb.NotificationReq{
			Num: int32(i),
		}
		_ = stream.Send(req)
	}
}
