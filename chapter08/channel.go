package main

import "fmt"

func main() {
	c := make(chan int)

	go func(s chan<- int) {
		for i := 0; i < 5; i++ {
			s <- i
		}
		close(s)
	}(c)
	fmt.Println("c:", c)

	for {
		value, ok := <-c
		fmt.Println("ok:", ok)
		if !ok {
			break
		}
		fmt.Println(value)
	}
}
