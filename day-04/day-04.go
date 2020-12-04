package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	validPassports := 0
	passportData := map[string]string{}
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line == "" {
			if checkPassportDataIsValid(passportData) {
				validPassports++
			}

			passportData = map[string]string{}
			continue
		}

		dataSections := strings.Split(line, " ")
		for _, section := range dataSections {
			sectionSplit := strings.Split(section, ":")
			passportData[sectionSplit[0]] = sectionSplit[1]
		}
	}

	fmt.Println(fmt.Sprintf("A total of %d valid passports were found.", validPassports))
}

func checkPassportDataIsValid(passportData map[string]string) bool {

	isValid := true
	requiredKeys := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
	}

	for _, requiredKey := range requiredKeys {
		_, exists := passportData[requiredKey]

		if !exists {
			isValid = false
			break
		}
	}

	return isValid
}
