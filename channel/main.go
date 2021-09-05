package main

import (
	"fmt"
)

func testFunc(c chan int) {
	for {
		fmt.Println(<-c)
		//time.Sleep(time.Second)
	}
}
func main() {
	c := make(chan int, 5)
	var i int = 0
	go testFunc(c)
	for {
		i = i + 1
		c <- i
	}
}

// go func で go routine
// chan int でint型のチャネル作成
// yy  <- c  / c <- xx
