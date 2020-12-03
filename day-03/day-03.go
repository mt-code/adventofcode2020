package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	position := 0
	treesEncountered := 0

	for i, line := range lines {

		if i == 0 {
			continue
		}

		// Reset our position to if we are going to exceed the line length, this
		// allows our pattern to repeat infinitely to the right.
		if (position + 3) > len(line)-1 {
			position = 3 - (len(line) - position)
		} else {
			position += 3
		}

		if string(line[position]) == "#" {
			treesEncountered++
		}
	}

	fmt.Println(fmt.Sprintf("A total of %d trees were encountered.", treesEncountered))
}
