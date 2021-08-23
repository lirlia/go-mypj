package main

import (
	"context"
	"fmt"
)

func addValue(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func readValue(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}

func main() {
	ctx := context.Background()

	ctx = addValue(ctx, "a", "123")
	ctx = addValue(ctx, "b", "456")
	fmt.Println(readValue(ctx, "a").(string))
	fmt.Println(readValue(ctx, "b").(string))
}
