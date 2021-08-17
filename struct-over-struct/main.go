package main

import "fmt"

type Name struct {
	Name string
}

func (n *Name) String() {
	fmt.Println(n.Name, "だよ〜")
}

type Person struct {
	Name
	Age int
}

type Person2 struct {
	*Name
	Age int
}

func (p *Person) String() {
	fmt.Println(p.Name.Name, "じゃないよ〜")
}

func main() {

	n := &Name{"taro"}
	n.String()

	p := &Person{*n, 20}

	// メソッドはオーバーライドされる
	p.String()

	// Persopn.Name.Nameを変えても n には影響しない（ポインタじゃないから）
	p.Name.Name = "hanako"
	p.String()
	n.String()

	pp := &Person2{n, 20}
	pp.String()
	p.Name.Name = "jiro"
	p.String()
	n.String()

}
