package main

import (
	"fmt"
)

func main() {

	// 文字列を準備
	s := "abcdefghijklmn"

	// スライス式を使って部分文字列を取り出す
	fmt.Println(s[:10])
	fmt.Println(s[3:10])
	fmt.Println(s[3:])
	fmt.Println(s[:])
}
