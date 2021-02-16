package main

import (
	"SoftwareGoDay2/server"
	"fmt"
)

func main() {
	r := server.NewServer()
	err := r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
