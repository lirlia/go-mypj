package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

type I interface {
	showName()
}

type Human struct {
	name string
}

func (h *Human) showName() {
	fmt.Println(h.name)
}

func describe(i I) {
	fmt.Println(i)
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

	var i I = &Human{"tanaka"}
	i.showName()

	describe(i)

}
