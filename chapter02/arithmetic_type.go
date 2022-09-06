package main

import (
	"fmt"
)

const C1 = 12345
const C2 = 34567

func main() {
	fmt.Println(C1)
	fmt.Println(C2)

	var x int = C1 * C2
	fmt.Println(x)

	var a int32 = 123
	var b int64 = 456

	fmt.Println(a * int32(b))
}
