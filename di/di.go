package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

type Car struct {
	Suzuki
	Shell
}

type CarDI struct {
	carInterface
}

type CarTest1 struct {
	Honda
	Panasonic
}
type CarTest2 struct {
	Suzuki
	Panasonic
}
type CarTest3 struct {
	Suzuki
	Shell
}

type carInterface interface {
	Wheel()
	Gas()
}

type Suzuki struct{}
type Honda struct{}
type Shell struct{}
type Panasonic struct{}

func (s *Suzuki) Wheel() { fmt.Println("Wheel : Suzuki") }
func (s *Honda) Wheel()  { fmt.Println("Wheel : Honda") }

func (s *Shell) Gas()     { fmt.Println("Gas : Shell") }
func (s *Panasonic) Gas() { fmt.Println("Gas : Panasonic") }

func NewCarDI(i carInterface) *CarDI {
	return &CarDI{i}
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

	fmt.Println("--Carにパーツが溶接されている状態--")
	car := &Car{}
	car.Wheel()
	car.Gas()

	fmt.Println("\n-------DIしてみる-------")
	fmt.Println("\n--Honda / Panasonic--")
	carTest1 := &CarTest1{}
	car1 := &CarDI{carTest1}
	car1.Wheel()
	car1.Gas()

	fmt.Println("\n--Suzuki / Panasonic--")
	carTest2 := &CarTest2{}
	car2 := &CarDI{carTest2}
	car2.Wheel()
	car2.Gas()

	fmt.Println("\n--Suzuki / Shell--")
	carTest3 := &CarTest3{}
	car3 := &CarDI{carTest3}
	car3.Wheel()
	car3.Gas()

}
