package main

import "fmt"

func main() {
	dst := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	fmt.Println(dst[2:])
	fmt.Println(src)

	count := copy(dst[2:], src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println("count", count)

}
