package main

import "fmt"

func main() {
	f1(false)
	f1(true)
}

func f1(p bool) {
	defer func() {
		fmt.Println("遅延関数開始")
		if err := recover(); err != nil {
			fmt.Println("遅延関数終了")
		}
	}()

	if p {
		panic("パニック発生")
	}
}
