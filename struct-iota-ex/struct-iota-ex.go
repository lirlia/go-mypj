package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

type State int

const (
	Running State = iota
	Stopping
	Sleeping
	Deleting
	Creating
)

const (
	value int = iota
)

// Int型にStringメソッドを書いておくと
// print時に自動でこれが使用されるっぽい
func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopping:
		return "Stopping"
	case Deleting:
		return "Deleting"
	case Creating:
		return "Creating"
	default:
		return "unknown"
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

	fmt.Println("Running ->", Running)
	fmt.Println("value ->", value)
}
