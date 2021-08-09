package main

import "fmt"

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
func main() {

	var i I = &Human{"tanaka"}
	i.showName()

	describe(i)

}
