package main

import (
	"fmt"
	"time"
)

func main() {

	// 現在時刻の取得
	time := time.Now()

	// 時刻を文字列にフォーマットして出力
	fmt.Printf("%04d/%02d/%02d %02d:%02d:%02d\n",
		time.Year(),
		time.Month(),
		time.Day(),
		time.Hour(),
		time.Minute(),
		time.Second())
}
