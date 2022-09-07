package main

import "fmt"

func main() {

	arr := [...]int{0, 1, 2, 3, 4}
	fmt.Println(arr)
	for i := range arr {
		fmt.Println(i)
	}
}
