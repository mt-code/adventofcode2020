package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(data), "\n")
	multiplySum := 1

	slopeOffsets := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for offsetCount, offsets := range slopeOffsets {
		rightIncrease := offsets[0]
		downIncrease := offsets[1]
		position := 0
		treesEncountered := 0

		for i := 0; i < len(lines); {

			if i == 0 {
				i += downIncrease
				continue
			}

			line := lines[i]

			// Reset our position to if we are going to exceed the line length, this
			// allows our pattern to repeat infinitely to the right.
			if (position + rightIncrease) > len(line)-1 {
				position = rightIncrease - (len(line) - position)
			} else {
				position += rightIncrease
			}

			if string(line[position]) == "#" {
				treesEncountered++
			}

			i += downIncrease
		}

		multiplySum *= treesEncountered
		fmt.Println(fmt.Sprintf("A total of %d trees were encountered for offset #%d.", treesEncountered, offsetCount))
	}

	fmt.Println(fmt.Sprintf("The total number of trees encountered multiplied together is %d.", multiplySum))
}
