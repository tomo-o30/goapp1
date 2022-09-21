package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Index("The Go Programming Language", "Go"))

	fmt.Println(strings.Index("The Go Programming Language", "xxx"))
}
