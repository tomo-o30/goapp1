package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
)

func main() {
	goroot := runtime.GOROOT()
	fileinfos, _ := ioutil.ReadDir(goroot)

	for _, finfo := range fileinfos {
		if finfo.IsDir() {
			fmt.Println(finfo.Name())
		}
	}
}
