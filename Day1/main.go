package main

import (
	"SofwareGoDay1/data"
	"fmt"
)

func main() {
	content, err := data.ReadFile("go.mod")
	fmt.Println(err)
	fmt.Println(content)
}
