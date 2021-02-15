package data

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFileCSV(path string) ([]string, error) {
	text, err := ioutil.ReadFile(path)
	arr := strings.Split(string(text), "\n")

	return arr, err
}

func LineToCSV(line string) ([]string, error) {
	arr := strings.Split(line, ",")

	if len(arr) == 0 {
		return nil, fmt.Errorf("empty string")
	}
	return arr, nil
}

func ReadFileJSON(path string) ([]string, error) {
	text, err := ioutil.ReadFile(path)
	arr := strings.Split(string(text), "\n")

	return arr, err
}

func LineToJSON(line string) ([]string, error) {
	arr := strings.Split(line, ":")
	return arr, nil
}
