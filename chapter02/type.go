package main

import (
	"fmt"
)

func main() {
	type myInteger int
	var i myInteger = 123

	i += 1
	fmt.Println(i)

	type myStruct struct {
		a int
		b int
	}

	var sample myStruct
	sample.a = 1
	sample.b = 2
	fmt.Println(sample)
	fmt.Println(sample.a)
	fmt.Println(sample.b)
}
