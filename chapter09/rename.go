package main

import (
	"fmt"
	"os"
)

func main() {
	os.MkdirAll("testdir", 0777)

	err := os.Rename("testdir", "newdir")

	// エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := os.Create("testfile")
	file.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
