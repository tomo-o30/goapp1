package main

import (
	"fmt"
	"os"
)

func main() {
	write()
	read()
}

func write() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	file.WriteString("test1\n")
	fmt.Fprintf(file, "test2\n")

}

func read() {
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer func() {
		file.Close()
	}()

	var str string
	fmt.Fprintf(file, "%s", &str)

	fmt.Println(str)

	buf := make([]byte, 16)

	size, _ := file.Read(buf)

	fmt.Println(size, buf)
}
