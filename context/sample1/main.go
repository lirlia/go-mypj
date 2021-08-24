package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	go parentProcess(ctx)
	time.Sleep(time.Second * 1000)
}

func parentProcess(ctx context.Context) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	go childProcess(ctx, "with timeout")

	childCtx, _ := context.WithCancel(ctx)
	childCtx2, _ := context.WithCancel(childCtx)

	go childProcess(childCtx, "child1")
	go childProcess(childCtx2, "child2")

	time.Sleep(time.Second * 5)
	cancel()
}

func childProcess(ctx context.Context, s string) {
	for i := 1; i <= 1000; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("%s: cancelled \n", s)
			fmt.Printf("reason: %v \n", ctx.Err().Error())
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("%s:%d sec..\n", s, i)
		}
	}
}
