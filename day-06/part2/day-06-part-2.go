package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	characterMap := map[string]int{}
	groupSize, sum := 0, 0

	for _, line := range lines {
		if line == "" {
			for _, value := range characterMap {
				if value == groupSize {
					sum++
				}
			}

			groupSize = 0
			characterMap = map[string]int{}
			continue
		}

		for _, char := range line {
			characterMap[string(char)] += 1
		}
		groupSize++
	}

	fmt.Println(fmt.Sprintf("The sum of answered questions is: %d", sum))
}
