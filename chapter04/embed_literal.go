package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
	Person
}

func main() {
	var e Employee = Employee{1, Person{"Mike", 28}}
	fmt.Println(e)
}
