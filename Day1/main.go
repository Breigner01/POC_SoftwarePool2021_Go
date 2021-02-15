package main

import (
	"SofwareGoDay1/data"
	"fmt"
)

func main() {
	files := []string{"resources/big_csv.csv", "resources/example.csv", "resources/medium_csv.csv"}

	for i := 0; i < len(files); i++ {
		content, err := data.ReadFile(files[i])
		for j := 0; j < len(content); j++ {
			arr, err2 := data.LineToCSV(content[j])
			fmt.Println(arr, err2)
		}
		fmt.Println(err)
		fmt.Println(content)
	}
}
