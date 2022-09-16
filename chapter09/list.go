package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()

	l.PushBack("a")
	l.PushBack("b")
	l.PushBack("c")

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Printf("%s\n", e.Value)
	}
}
