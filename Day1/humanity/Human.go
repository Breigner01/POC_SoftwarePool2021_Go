package humanity

import (
	"SofwareGoDay1/data"
	"fmt"
	"strconv"
	"strings"
)

type Human struct {
	Name    string
	Age     int
	Country string
}

func NewHumanFromCSV(csv []string) (*Human, error) {
	var human Human
	var err error

	if len(csv) != 3 {
		return nil, fmt.Errorf("invalid argument")
	}
	human.Name = csv[0]
	human.Age, err = strconv.Atoi(csv[1])
	human.Country = csv[2]
	return &human, err
}

func NewHumansFromCsvFile(path string) ([]*Human, error) {
	arr, err := data.ReadFileCSV(path)
	var humans []*Human

	if err != nil {
		return nil, err
	}
	for i := 0; i < len(arr); i++ {
		line, err := data.LineToCSV(arr[i])
		if len(line) != 3 {
			continue
		}
		if err != nil {
			continue
		}
		human, err := NewHumanFromCSV(line)
		if err == nil {
			humans = append(humans, human)
		}
	}
	return humans, nil
}

func NewHumanFromJSON(json []string) (*Human, error) {
	var human Human
	var err error

	if len(json) != 3 {
		return nil, fmt.Errorf("invalid argument")
	}
	human.Name = json[0]
	human.Age, err = strconv.Atoi(json[1])
	human.Country = json[2]
	return &human, err
}

func NewHumansFromJsonFile(path string) ([]*Human, error) {
	var humans []*Human
	var humanContent []string
	var content string
	arr, err := data.ReadFileJSON(path)

	if err != nil {
		return nil, err
	}
	for _, line := range arr {
		if strings.Contains(line, "[") || strings.Contains(line, "{") || strings.Contains(line, "]") ||
			strings.Contains(line, "}") {
			continue
		}
		element, _ := data.LineToJSON(line)
		if strings.Contains(element[0], "Name") || strings.Contains(element[0], "Age") ||
			strings.Contains(element[0], "Country") {
			content = strings.ReplaceAll(element[1], " ", "")
			content = strings.ReplaceAll(content, "\"", "")
			content = strings.ReplaceAll(content, ",", "")
			humanContent = append(humanContent, content)
		}
		if len(content) == 3 {
			human, _ := NewHumanFromJSON(humanContent)
			humans = append(humans, human)
		}
	}
	return humans, nil
}
