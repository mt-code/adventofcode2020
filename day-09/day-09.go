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

	for i := 25; i < len(lines); i++ {
		isFound := false
		total, _ := strconv.Atoi(lines[i])

		for k := i; k >= i-25; k-- {
			num1, _ := strconv.Atoi(lines[k])
			for j := i; j >= i-25; j-- {
				num2, _ := strconv.Atoi(lines[j])
				if num1+num2 == total {
					isFound = true
				}
			}
		}

		if !isFound {
			fmt.Println(lines[i])
			break
		}
	}
}
