package main

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/lirlia/go-mypj/grpc/test/myproto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	var name string = "japan"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := pb.NewCloudClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetDate(ctx, &pb.DateRequest{Locale: name})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.GetLocale())

}
