package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	pb "github.com/lirlia/go-mypj/grpc/mygrpc/proto"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func stringToInt64(i []string) []int64 {
	f := make([]int64, len(i))
	for n := range i {
		num, _ := strconv.Atoi(i[n])
		f[n] = int64(num)
	}
	return f
}

func main() {
	var way string
	var nums []int64
	if len(os.Args) > 1 {
		way = os.Args[1]
		nums = stringToInt64(os.Args[2:])
	}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		panic(err)
	}

	defer conn.Close()
	c := pb.NewCalcClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var r *pb.CalcRes
	switch way {
	case "add":
		r, _ = c.Add(ctx, &pb.CalcReq{Params: nums})
	case "minus":
		r, _ = c.Minus(ctx, &pb.CalcReq{Params: nums})
	case "product":
		r, _ = c.Product(ctx, &pb.CalcReq{Params: nums})
	case "div":
		r, _ = c.Div(ctx, &pb.CalcReq{Params: nums})
	default:
		fmt.Println(way, "is not supported.")
		panic(err)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println(r.GetParams())

}
