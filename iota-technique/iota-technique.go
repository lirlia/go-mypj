package main

import "fmt"

const (
	name = "name_" + string(iota/10%10+'0') + string(iota/1%10+'0')

	aa
	bb
	cc
	dd
)

type ByteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main() {
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	fmt.Println(dd)

	fmt.Println(KB, MB, GB, TB)

}
