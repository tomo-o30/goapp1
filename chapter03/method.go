package main

import "fmt"

type myType int

func (value myType) printLn() {
	fmt.Println(value)
}

func main() {
	var z myType = 1234
	fmt.Println(z)
	z.printLn()
}
