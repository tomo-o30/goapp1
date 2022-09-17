package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()
	for i := 0; i < 5; i++ {
		list.PushBack(i)
	}

	list.Init()

	fmt.Println("要素数:", list.Len())
}
