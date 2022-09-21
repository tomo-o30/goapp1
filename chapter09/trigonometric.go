package main

import (
	"fmt"
	"math"
)

func main() {

	// 45度
	var degree float64 = 45

	// ラジアンに変換
	radian := degree * math.Pi / 180

	// ラジアンを出力
	fmt.Println("角度(ラジアン):", radian)

	// 正弦(サイン)を求める
	fmt.Println("正弦:", math.Sin(radian))

	// 余弦(コサイン)を求める
	fmt.Println("余弦:", math.Cos(radian))

	// 正接(タンジェント)を求める
	fmt.Println("正接:", math.Tan(radian))
}
