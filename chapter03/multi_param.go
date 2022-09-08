package main

import "fmt"

func main() {
	f2(fi())
}

func fi() (int, string, float32) {
	return 0, "xyz", 3.14
}

func f2(i int, s string, f float32) {
	fmt.Println(i, s, f)
}
