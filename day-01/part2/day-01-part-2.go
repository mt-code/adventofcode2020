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

		for k := i + 1; k < len(lines); k++ {
			secondNum, _ := strconv.Atoi(lines[k])
			currentTotal := firstNum + secondNum

			if currentTotal < 2020 {
				for _, line := range lines {
					thirdNum, _ := strconv.Atoi(line)

					if (currentTotal + thirdNum) == 2020 {
						fmt.Println("#1", firstNum)
						fmt.Println("#2", secondNum)
						fmt.Println("#3", thirdNum)
						fmt.Println("#1 * #2 * #3 = ", firstNum*secondNum*thirdNum)
						return
					}
				}
			}
		}
	}
}
