package main

import "fmt"

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

func main() {

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
