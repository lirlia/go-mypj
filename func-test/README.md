## source

```go
package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

type TestFunc func(int, int) int
type ValuePool struct {
	a int
	b int
}

func sum(a int, b int) int {
	return a + b
}

func diff(a int, b int) int {
	return a - b
}

func (t ValuePool) Run(f TestFunc) {
	fmt.Println(f(t.a, t.b))
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
	a, b := 100, 100

	var t TestFunc
	t = sum
	fmt.Println(t(a, b))

	t = diff
	fmt.Println(t(a, b))

	v := &ValuePool{a, b}
	v.Run(sum)
	v.Run(diff)

}

```

## Result

```sh
200
0
200
0
```
