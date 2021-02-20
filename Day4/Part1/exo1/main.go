package main

import (
	"SoftwareGoDay2/server"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	r := server.NewServer()
	err := godotenv.Load()

	if err != nil {
		fmt.Println("No .env file found")
	}
	err = r.Run()
	if err != nil {
		fmt.Println(err)
	}
}
