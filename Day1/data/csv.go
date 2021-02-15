package data

import (
	"io/ioutil"
	"strings"
)

func ReadFile(filename string) ([]string, error) {
	text, err := ioutil.ReadFile(filename)
	arr := strings.Split(string(text), "\n")

	return arr, err
}
