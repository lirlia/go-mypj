package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

type Human struct {
	age int
	sex string
}

func (h *Human) CountUpAge() {
	h.age++
}

func ChangeSex(h *Human) {
	switch h.sex {
	case "male":
		h.sex = "female"
	case "female":
		h.sex = "male"
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

	h := &Human{10, "male"}
	fmt.Println("--default--")
	fmt.Printf("%+v\n", *h)

	fmt.Println("--change sex--")
	ChangeSex(h)
	fmt.Printf("%+v\n", *h)
	fmt.Println("--change sex--")
	ChangeSex(h)
	fmt.Printf("%+v\n", *h)
	fmt.Println("--change sex--")
	ChangeSex(h)
	fmt.Printf("%+v\n", *h)

	fmt.Println("--count up age--")
	h.CountUpAge()
	fmt.Printf("%+v\n", *h)
	h.CountUpAge()
	fmt.Printf("%+v\n", *h)
	h.CountUpAge()
	fmt.Printf("%+v\n", *h)
	h.CountUpAge()
	fmt.Printf("%+v\n", *h)
	h.CountUpAge()
	fmt.Printf("%+v\n", *h)

}
