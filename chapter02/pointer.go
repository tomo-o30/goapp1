package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("ptrのアドレス", ptr)
	var i int = 12345
	ptr = &i

	fmt.Println("iのアドレス", &i)
	fmt.Println("ptrのアドレス", ptr)

	fmt.Println("iの値", i)
	fmt.Println("ポインタ経由のiの値", *ptr)

	*ptr = 999
	fmt.Println("ポインタ経由で変更したiの値", i)

}
