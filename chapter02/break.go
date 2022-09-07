package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 2 {
			break
		}
	}
}
