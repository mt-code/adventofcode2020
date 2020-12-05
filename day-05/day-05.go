package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	maxSeatId := 0
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		minRow, maxRow := 0, 128
		minCol, maxCol := 0, 8

		for _, char := range line {
			rowAdjustment := (maxRow - minRow) / 2
			colAdjustment := (maxCol - minCol) / 2

			switch string(char) {
			case "F":
				maxRow -= rowAdjustment
			case "B":
				minRow += rowAdjustment
			case "R":
				minCol += colAdjustment
			case "L":
				maxCol -= colAdjustment
			}
		}

		row := maxRow - 1
		col := maxCol - 1
		seatId := (row * 8) + col

		if seatId > maxSeatId {
			maxSeatId = seatId
		}
	}

	fmt.Println(fmt.Sprintf("The max seat ID that we found was %d.", maxSeatId))
}
