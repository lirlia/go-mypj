package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// 計算結果格納用
type result struct {
	value map[int]int
}

// フィボナッチ数を呼び出す
func (r result) fib(n int) (int, error) {

	if n < 0 {
		return 0, fmt.Errorf("最小値は0です")
	}

	if n == 0 || n == 1 {
		return n, nil
	}

	// Mapから計算結果を取り出す
	v, ok := r.value[n]

	// 値が存在しない場合のみ計算する
	if !ok {
		v1, _ := r.fib(n - 2)
		v2, _ := r.fib(n - 1)
		v = v1 + v2
		r.value[n] = v
	}

	return v, nil
}

func main() {
	flag.Parse()
	v, err := strconv.Atoi(flag.Args()[0])

	if err != nil {
		fmt.Println(flag.Args()[0], "is invalid argument. arg must by a number")
		os.Exit(1)
	}

	//	r := &result{map[int]int{0: 0, 1: 1}}
	r := &result{
		value: map[int]int{},
	}

	v, err = r.fib(v)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(v)
}
