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
	totalAcc := 0

	for i := 0; i < len(lines); {
		if returnAcc := ContinueExecution(i, lines); returnAcc != -1 {
			fmt.Println(returnAcc + totalAcc)
			break
		}
		acc, pos, _ := ProcessInstruction(i, lines[i], false)
		i, totalAcc = pos, totalAcc+acc
	}
}

func ContinueExecution(currentPosition int, lines []string) (acc int) {
	jumpsEncountered := make(map[int]bool)

	for i := currentPosition; i < len(lines); {
		if jumpsEncountered[i] {
			return -1
		}

		returnAcc, returnPos, isJump := ProcessInstruction(i, lines[i], i == currentPosition)
		jumpsEncountered[i] = isJump
		acc += returnAcc
		i = returnPos
	}
	return acc
}

func ProcessInstruction(position int, line string, shouldSwitch bool) (int, int, bool) {
	split := strings.Split(line, " ")
	instruction := split[0]
	value, _ := strconv.Atoi(split[1])

	if shouldSwitch {
		switch instruction {
		case "jmp":
			instruction = "nop"
		case "nop":
			instruction = "jmp"
		}
	}

	switch instruction {
	case "acc":
		return value, position + 1, false
	case "jmp":
		return 0, position + value, true
	default:
		return 0, position + 1, false
	}
}
