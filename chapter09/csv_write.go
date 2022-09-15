package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer func() {
		file.Close()
	}()

	csvWriter := csv.NewWriter(file)

	csvWriter.Write([]string{"No.", "商品", "価格", "備考"})
	csvWriter.Write([]string{"1", "キャベツ", "150", "とれたて新鮮"})
	csvWriter.Write([]string{"1", "キャベツ", "150", "とれたて新鮮"})
	csvWriter.Write([]string{"1", "キャベツ", "150", "とれたて新鮮"})

	csvWriter.Flush()
}
