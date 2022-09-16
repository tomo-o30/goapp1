package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}

	fmt.Println(l.Front())
	fmt.Println(l.Back())
	fmt.Println(l.Len())

}
