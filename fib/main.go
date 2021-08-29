package main

import (
	"flag"
	"fmt"
	"strconv"
)

// 計算結果格納用
type result struct {
	value map[int]int
}

// フィボナッチ数を呼び出す
func (r result) fib(n int) int {

	// Mapから計算結果を取り出す
	v, ok := r.value[n]

	// 値が存在しない場合のみ計算する
	if !ok {
		v = r.fib(n-2) + r.fib(n-1)
		r.value[n] = v
	}

	return v
}

func main() {
	flag.Parse()
	v, err := strconv.Atoi(flag.Args()[0])

	if err != nil {
		panic(err)
	}

	if v < 0 {
		err = fmt.Errorf("最小値は0")
		panic(err)
	}

	r := &result{map[int]int{0: 0, 1: 1}}

	fmt.Println(r.fib(v))
}
