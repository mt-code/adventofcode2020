package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type BagContents struct {
	name  string
	count int
}

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	containsMap := make(map[string][]BagContents)

	for _, line := range lines { // First pass reads all lines into our map
		lineSplit := strings.Split(line, " bags contain ")
		bag := lineSplit[0]
		contents := lineSplit[1]

		regex := regexp.MustCompile("(\\d) ([\\w ]+) bag")
		matches := regex.FindAllStringSubmatch(contents, -1)

		for i := 0; i < len(matches); i++ {
			bagCount, _ := strconv.Atoi(matches[i][1])
			bagName := matches[i][2]
			containsMap[bag] = append(containsMap[bag], BagContents{bagName, bagCount})
		}
	}

	sum := countBagsRecursive("shiny gold", containsMap)
	fmt.Println(fmt.Sprintf("Total count: %d", sum))
}

func countBagsRecursive(bagName string, containsMap map[string][]BagContents) (sum int) {
	for _, containedBag := range containsMap[bagName] {
		sum += containedBag.count
		sum += containedBag.count * (countBagsRecursive(containedBag.name, containsMap))
	}
	return sum
}
