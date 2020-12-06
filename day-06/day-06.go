package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	characters := ""
	sum := 0

	for _, line := range lines {
		if line == "" {
			sum += len(characters)
			characters = ""
			continue
		}

		for _, char := range line {
			if !strings.Contains(characters, string(char)) {
				characters += string(char)
			}
		}
	}

	fmt.Println(fmt.Sprintf("The sum of answered questions is: %d", sum))
}
