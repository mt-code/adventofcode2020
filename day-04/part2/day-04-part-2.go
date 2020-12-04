package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	validPassports := 0
	passportData := map[string]string{}
	content, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if line == "" {
			if checkPassportIsValid(passportData) {
				validPassports++
			}

			passportData = map[string]string{}
			continue
		}

		dataSections := strings.Split(line, " ")
		for _, section := range dataSections {
			sectionSplit := strings.Split(section, ":")
			key, value := sectionSplit[0], sectionSplit[1]

			if checkDataIsValid(key, value) {
				passportData[key] = value
			}
		}
	}

	fmt.Println(fmt.Sprintf("A total of %d valid passports were found.", validPassports))
}

func checkPassportIsValid(passportData map[string]string) bool {
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

func checkDataIsValid(key, value string) bool {
	validationRules := map[string]string{
		"byr": "^\\d{4}$",                      // four digits
		"iyr": "^\\d{4}$",                      // four digits
		"eyr": "^\\d{4}$",                      // four digits
		"hgt": "^(\\d+)(cm|in)$",               // a number followed by either cm or in
		"hcl": "^#[0-9a-f]{6}$",                // a # followed by exactly six characters 0-9 or a-f.
		"ecl": "^amb|blu|brn|gry|grn|hzl|oth$", // exactly one of: amb blu brn gry grn hzl oth.
		"pid": "^\\d{9}$",                      // a nine-digit number, including leading zeroes.
	}

	regex := regexp.MustCompile(validationRules[key])
	if !regex.MatchString(value) {
		return false
	}

	switch key {
	case "byr":
		num, _ := strconv.Atoi(value)
		return num >= 1920 && num <= 2002

	case "iyr":
		num, _ := strconv.Atoi(value)
		return num >= 2010 && num <= 2020

	case "eyr":
		num, _ := strconv.Atoi(value)
		return num >= 2020 && num <= 2030

	case "hgt":
		matches := regex.FindAllStringSubmatch(value, -1)
		height, _ := strconv.Atoi(matches[0][1])
		unit := matches[0][2]

		if unit == "cm" {
			return height >= 150 && height <= 193
		} else {
			return height >= 59 && height <= 76
		}

	default:
		return true
	}
}
