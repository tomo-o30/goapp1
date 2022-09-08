package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func main() {
	var a, b int = 1, 2
	c := plus(a, b)
	fmt.Printf("%d\n", c)
}
