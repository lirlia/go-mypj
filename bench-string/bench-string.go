package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

const (
	times = 500000000
)

func BenchmarkOnlyString() {
	base := []string{}
	for i := 0; i < times; i++ {
		x := string(i + '0')
		base = append(base, x)
	}
}

func BenchmarkCastString() {
	base := []string{}
	for i := 0; i < times; i++ {
		x := strconv.Itoa(i)
		base = append(base, x)
	}
}

// 引数の関数にかかった時間を計測します
func MeasureTime(f func()) {
	start := time.Now()
	f()
	end := time.Now()
	fmt.Println(
		runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(),
		end.Sub(start))
}

func main() {
	MeasureTime(BenchmarkCastString)
	MeasureTime(BenchmarkOnlyString)
}
