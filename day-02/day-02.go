package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")

	if err != nil {
		log.Fatal("File read error: ", err)
	}

	validPasswords := 0
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		regex := regexp.MustCompile("(\\d+)-(\\d+) (\\w): (\\w+)")
		matches := regex.FindAllStringSubmatch(line, -1)

		count := 0
		minCount, _ := strconv.Atoi(matches[0][1])
		maxCount, _ := strconv.Atoi(matches[0][2])
		char := matches[0][3]
		sample := matches[0][4]

		sampleRunes := []rune(sample)
		for _, charRune := range sampleRunes {
			if string(charRune) == char {
				count++
			}
		}

		if count >= minCount && count <= maxCount {
			validPasswords++
		}
	}

	fmt.Println(fmt.Sprintf("A total of %d valid passwords were found.", validPasswords))
}
