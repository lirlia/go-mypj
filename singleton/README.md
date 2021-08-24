## source

```go
package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"strconv"
)

type Counter struct {
	count int
}

func (c *Counter) String() string {
	return strconv.Itoa(c.count)
}

var Instance = &Counter{}

func SetInstance(c int) {
	Instance.count = Instance.count + c
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

	SetInstance(4)
	fmt.Println(Instance.String())
	SetInstance(6)
	fmt.Println(Instance.String())
	SetInstance(5)
	fmt.Println(Instance.String())

	Instance.count = 100
	fmt.Println(Instance.String())
}

```

## Result

```sh
4
10
15
100
```
