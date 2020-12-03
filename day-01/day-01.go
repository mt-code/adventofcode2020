package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("File read error: ", err)
	}

	lines := strings.Split(string(data), "\n")

	for i := 0; i < len(lines); i++ {
		firstNum, _ := strconv.Atoi(lines[i])

		for k := i + 1; k < len(lines); k++ { // Start from i+1 as anything before i has already been checked
			secondNum, _ := strconv.Atoi(lines[k])

			if (firstNum + secondNum) == 2020 {
				fmt.Println("#1: ", firstNum)
				fmt.Println("#2: ", secondNum)
				fmt.Println("#1 * #2 = ", firstNum*secondNum)
				return
			}
		}
	}
}
