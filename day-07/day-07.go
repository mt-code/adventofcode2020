package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	fitsInMap := make(map[string][]string)

	for _, line := range lines { // First pass reads all lines into our map
		lineSplit := strings.Split(line, " bags contain ")
		bag := lineSplit[0]
		contents := lineSplit[1]

		regex := regexp.MustCompile("\\d ([\\w ]+) bag")
		matches := regex.FindAllStringSubmatch(contents, -1)

		for i := 0; i < len(matches); i++ {
			bagName := matches[i][1]
			fitsInMap[bagName] = append(fitsInMap[bagName], bag)
		}
	}

	alreadyChecked := make(map[string]bool)
	sum := countBagsRecursive("shiny gold", fitsInMap, alreadyChecked)
	fmt.Println(fmt.Sprintf("Total count: %d", sum))
}

func countBagsRecursive(bagName string, fitsInMap map[string][]string, alreadyChecked map[string]bool) (count int) {
	if alreadyChecked[bagName] {
		count--
	}

	alreadyChecked[bagName] = true
	for _, bag := range fitsInMap[bagName] {
		count += countBagsRecursive(bag, fitsInMap, alreadyChecked)
		count++
	}

	return count
}
