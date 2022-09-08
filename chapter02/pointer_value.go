package main

import "fmt"

func main() {
	a, b := 1, 1
	fmt.Println(&b)
	double(a, &b)

	fmt.Println("値渡し：", a)
	fmt.Println("ポインタ渡し：", b)
}

func double(x int, y *int) {
	x = x * 2
	fmt.Println(*y)
	*y = *y * 2
}
