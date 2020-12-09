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
	target := 26134589

	for i := 0; i < len(lines); i++ {
		total, _ := strconv.Atoi(lines[i])
		min, max := total, total

		for k := i + 1; k < len(lines); k++ {
			num, _ := strconv.Atoi(lines[k])
			total += num

			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
			if total == target {
				fmt.Println(min + max)
			}
		}
	}
}
