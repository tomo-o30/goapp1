package main

import "fmt"

type myType int

func (value *myType) add(increment myType) myType {
	*value += increment
	return *value
}

func main() {
	var i myType

	fmt.Println(i)
	fmt.Println(i.add(3))
	mv := i.add
	fmt.Println(mv(3))
	fmt.Println(i)
}
