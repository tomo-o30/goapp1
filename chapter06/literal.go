package main

import "fmt"

func main() {
	array1 := [5]float32{}
	array2 := [6]int{1, 2, 3, 4}
	array3 := [...]string{"one", "two", "three"}
	array4 := []float32{}

	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)

}
