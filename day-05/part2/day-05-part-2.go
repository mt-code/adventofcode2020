package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	var seatIds []int
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
		seatIds = append(seatIds, seatId)
	}

	sort.Ints(seatIds)
	minSeatId, maxSeatId := 84, 866

	for i := 0; i < maxSeatId-minSeatId; i++ {
		if seatIds[i] != i+minSeatId {
			fmt.Println(fmt.Sprintf("Missing seat ID found: %d", i+minSeatId))
			break
		}
	}
}
