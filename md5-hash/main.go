package main

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"io"
)

func main() {

	key := "test"
	cnt := 3

	h := md5.New()
	// hにkeyを書き込んでいる
	if _, err := io.WriteString(h, key); err != nil {
		fmt.Println(err)
	}

	// ハッシュの取得
	digest := h.Sum(nil)

	// 数値に変換
	higherUint64 := binary.BigEndian.Uint64(digest)

	// modを取得
	fmt.Println(higherUint64 % uint64(cnt))
	higherUint64 = binary.LittleEndian.Uint64(digest)
	fmt.Printf("%x", higherUint64)

	data := []byte(key)
	fmt.Printf("%x", md5.Sum(data))
	// higherUint64 = binary.BigEndian.Uint64(md5.Sum(data))
	// fmt.Printf("%x", higherUint64)
}
