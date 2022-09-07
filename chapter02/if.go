package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " は", "偶数 ")
		} else {
			fmt.Println(i, " は奇数")
		}
	}
}
