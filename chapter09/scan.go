package main

import "fmt"

func main() {
	var fname, name string
	fmt.Println("あなたの名前は（姓・名）？")

	fmt.Scanln(&fname)
	fmt.Scanln(&name)

	fmt.Println("あなたの年齢は？")

	var age int
	fmt.Scanf("%d", &age)

	fmt.Printf("名前: %s %s 年齢:%d \n", fname, name, age)

}
