package main

import (
	"fmt"
	"time"
)

func main() {
	for day := time.Sunday; day <= time.Saturday; day++ {
		switch day {
		case time.Sunday:
			fallthrough
		case time.Saturday:
			fmt.Println("休日")
		default:
			fmt.Println(day, "多分平日")

		}
	}
}
