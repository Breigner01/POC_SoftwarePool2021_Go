package main

import (
	"SofwareGoDay1/humanity"
	"fmt"
)

func main() {
	files := []string{"resources/big_csv.csv", "resources/example.csv", "resources/medium_csv.csv"}

	for _, file := range files {
		humans, _ := humanity.NewHumansFromCsvFile(file)
		for _, human := range humans {
			fmt.Println(human.Name, human.Age, human.Country)
		}
	}

	for _, file := range files {
		humans, _ := humanity.NewHumansFromJsonFile(file)
		for _, human := range humans {
			fmt.Println(human.Name, human.Age, human.Country)
		}
	}
}
