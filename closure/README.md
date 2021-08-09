## source

```go
package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func addr() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func createCounter() func() int {
	n := 0
	return func() int {
		n += 1
		return n
	}
}

var lang = "japanese"

func hanteiHello() func() string {

	var msg string
	switch lang {
	case "english":
		msg = "hello"
	case "japanese":
		msg = "こんにちは"
	default:
		msg = "no such lang"
	}

	return func() string {
		return msg
	}
}

func makeFibGen() func() int {
	f1 := 0
	f2 := 1
	return func() int {
		f2, f1 = (f1 + f2), f2
		return f1
	}
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
	a, b := addr(), addr()
	counter := createCounter()
	hello := hanteiHello()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i), b(10*i))
	}

	fmt.Println("カウンター")
	fmt.Println(counter())
	fmt.Println(counter())

	fmt.Println("挨拶")
	fmt.Println(hello())

	fmt.Println("フィボナッチ数列")
	gen := makeFibGen()
	for i := 0; i < 10; i++ {
		fmt.Println(gen())
	}
}

```

## Result

```sh
0 0
1 10
3 30
6 60
10 100
15 150
21 210
28 280
36 360
45 450
カウンター
1
2
挨拶
こんにちは
フィボナッチ数列
1
1
2
3
5
8
13
21
34
55
```
