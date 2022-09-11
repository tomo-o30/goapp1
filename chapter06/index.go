package main

import "fmt"

func main() {
	var date [7]string

	date[0] = "月曜日"
	date[1] = "火曜日"
	date[2] = "水曜日"
	date[3] = "木曜日"
	date[4] = "金曜日"
	date[5] = "土曜日"
	date[6] = "日曜日"

	for i := 0; i < len(date); i++ {
		fmt.Println(date[i], " ")
	}
	fmt.Println()

	for _, value := range date {
		fmt.Println(value)
	}
}
