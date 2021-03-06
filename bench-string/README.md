## source

```go
package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

const (
	times = 50000000
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

func preStart() {
	_, f, _, _ := runtime.Caller(0)
	bytes, _ := ioutil.ReadFile(f)
	fmt.Println("## source")
	fmt.Println()
	fmt.Println("```go")
	fmt.Println(string(bytes))
	fmt.Println("```")
	fmt.Println()
	fmt.Println("## Result")
	fmt.Println()
	fmt.Println("```sh")
}
func preEnd() { fmt.Println("```") }

func main() {

	preStart()
	defer preEnd()
	MeasureTime(BenchmarkCastString)
	MeasureTime(BenchmarkOnlyString)
}

```

## Result

```sh
main.BenchmarkCastString 2.604347584s
main.BenchmarkOnlyString 2.660131375s
```
