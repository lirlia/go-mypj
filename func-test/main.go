package main

import "fmt"

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

func main() {
	a, b := 100, 100

	var t TestFunc
	t = sum
	fmt.Println(t(a, b))

	t = diff
	fmt.Println(t(a, b))

	v := &ValuePool{a, b}
	v.Run(sum)
	v.Run(diff)

	v.Run(TestFunc(sum))
}
