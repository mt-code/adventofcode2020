package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")
	acc := 0
	var processedInstructions []int

	for i := 0; i < len(lines); {
		split := strings.Split(lines[i], " ")
		value, _ := strconv.Atoi(split[1])

		if sliceContains(processedInstructions, i) {
			fmt.Println(fmt.Sprintf("Accumulator value = %d", acc))
			break
		}

		processedInstructions = append(processedInstructions, i)

		switch split[0] {
		case "acc":
			acc += value
			i++
		case "jmp":
			i += value
		default:
			i++
		}
	}
}

func sliceContains(slice []int, value int) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
