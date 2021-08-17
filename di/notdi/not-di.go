package main

import (
	"fmt"
)

type Car struct {
	Suzuki
	Shell
}

type Suzuki struct{}
type Honda struct{}
type Shell struct{}
type Panasonic struct{}

func (s *Suzuki) Wheel() { fmt.Println("Wheel : Suzuki") }
func (s *Shell) Gas()    { fmt.Println("Gas : Shell") }

func main() {
	fmt.Println("--Carにパーツが溶接されている状態--")
	car := &Car{}
	car.Wheel()
	car.Gas()
}
