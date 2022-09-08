package main

import "fmt"

func culc(a int, b int) (int, int, int, float32) {
	return a + b, a - b, a * b, float32(a) / float32(b)
}

func main() {
	a, b, c, d := culc(1, 2)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
