## source

```go
package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

type I interface {
	M() error
}

type A struct {
	str string
}

func (a *A) M() error {

	if a == nil {
		return fmt.Errorf("receiver is null")
	}
	if a.str == "" {
		return fmt.Errorf("str is blank")
	}
	fmt.Println(a.str)
	return nil
}

func (a *A) setStr(s string) {
	a.str = s
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

	fmt.Println("---a---")
	// interface Iに構造体Aのポインタを代入
	var i I = &A{"aaa"}

	if err := i.M(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("---b1---")
	fmt.Println("bはAのポインタ型であると定義されたが、中身は入っていないのでエラーとなる")
	// bが構造体Aのポインタ型であることを宣言
	var b *A

	// 中身に何も入っていないのでエラー
	if err := b.M(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("---b2---")
	fmt.Println("bにAインスタンスを突っ込んだ")
	// bに構造体Aのポインタを入れる
	b = &A{"bbb"}
	if err := b.M(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("---c1---")
	fmt.Println("cは値を空で定義した構造体Aのインスタンス")
	// cが構造体Aのポインタ型であることを宣言
	var c *A = &A{}
	if err := c.M(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("---c2---")
	fmt.Println("cは値を空で定義した構造体Aのインスタンス、メソッドで値セット")
	c.setStr("ccc")

	if err := c.M(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("---c3---")
	fmt.Println("ポインタに直接値を代入")
	c.str = "ccc2"
	if err := c.M(); err != nil {
		fmt.Println(err)
	}
	// 以下はメモリのアクセス違反エラーがでる
	// fmt.Println("---d---")
	// // dが構造体Aのポインタ型であることを宣言
	// var d *A
	// d.setStr("ddd")
	// if err := d.M(); err != nil {
	// 	fmt.Println(err)
	// }
}

```

## Result

```sh
---a---
aaa
---b1---
bはAのポインタ型であると定義されたが、中身は入っていないのでエラーとなる
receiver is null
---b2---
bにAインスタンスを突っ込んだ
bbb
---c1---
cは値を空で定義した構造体Aのインスタンス
str is blank
---c2---
cは値を空で定義した構造体Aのインスタンス、メソッドで値セット
ccc
---c3---
ポインタに直接値を代入
ccc2
```
