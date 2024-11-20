package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func NewChallenge202303() adventofcode.Challenge {
	return Challenge202303{}
}

type Challenge202303 struct{}

func (Challenge202303) GetYear() int {
	return 2023
}

func (Challenge202303) GetDay() int {
	return 03
}

func (c Challenge202303) ExecuteFirst(input string) (string, error) {
	characterRegex := regexp.MustCompile("[/*@&%$+=#-]")
	numberRegex := regexp.MustCompile("[0-9]")

	searchNumber := func(x, y int, split []string) int {
		line := split[y]
		if !numberRegex.MatchString(string(line[x])) {
			return 0
		}
		left := x
		for ; left > 0 && numberRegex.MatchString(string(line[left-1])); left-- {
		}
		numberString := ""

		for i := left; i != len(line) && numberRegex.MatchString(string(line[i])); i++ {
			numberString += string(line[i])
		}
		atoi, _ := strconv.Atoi(numberString)
		return atoi
	}

	checkUsedNumbers := func(number int, usedNumbers *[]int) int {
		for _, usedNumber := range *usedNumbers {
			if usedNumber == number {
				return 0
			}
		}
		*usedNumbers = append(*usedNumbers, number)
		return number
	}

	locations := characterRegex.FindAllStringIndex(strings.ReplaceAll(input, "\n", ""), len(input))
	split := strings.Split(input, "\n")
	lineLength := len(split[0])
	count := 0

	for _, locationArray := range locations {
		usedNumbers := make([]int, 0)
		location := locationArray[0]
		x := location % lineLength
		y := int(math.Floor(float64(location) / float64(lineLength)))
		count += checkUsedNumbers(searchNumber(x-1, y-1, split), &usedNumbers)
		count += checkUsedNumbers(searchNumber(x, y-1, split), &usedNumbers)
		count += checkUsedNumbers(searchNumber(x+1, y-1, split), &usedNumbers)
		count += checkUsedNumbers(searchNumber(x+1, y, split), &usedNumbers)
		count += checkUsedNumbers(searchNumber(x+1, y+1, split), &usedNumbers)
		count += checkUsedNumbers(searchNumber(x, y+1, split), &usedNumbers)
		count += checkUsedNumbers(searchNumber(x-1, y+1, split), &usedNumbers)
		count += checkUsedNumbers(searchNumber(x-1, y, split), &usedNumbers)
	}
	return strconv.Itoa(count), nil
}

func (c Challenge202303) ExecuteSecond(input string) (string, error) {
	characterRegex := regexp.MustCompile("[*]")
	numberRegex := regexp.MustCompile("[0-9]")

	searchNumber := func(x, y int, split []string) int {
		line := split[y]
		if !numberRegex.MatchString(string(line[x])) {
			return 1
		}
		left := x
		for ; left > 0 && numberRegex.MatchString(string(line[left-1])); left-- {
		}
		numberString := ""

		for i := left; i != len(line) && numberRegex.MatchString(string(line[i])); i++ {
			numberString += string(line[i])
		}
		atoi, _ := strconv.Atoi(numberString)
		return atoi
	}

	checkUsedNumbers := func(number int, usedNumbers *[]int) int {
		for _, usedNumber := range *usedNumbers {
			if usedNumber == number {
				return 1
			}
		}
		*usedNumbers = append(*usedNumbers, number)
		return number
	}

	locations := characterRegex.FindAllStringIndex(strings.ReplaceAll(input, "\n", ""), len(input))
	split := strings.Split(input, "\n")
	lineLength := len(split[0])
	count := 0

	for _, locationArray := range locations {
		multi := 1
		usedNumbers := make([]int, 0)
		location := locationArray[0]
		x := location % lineLength
		y := int(math.Floor(float64(location) / float64(lineLength)))
		multi *= checkUsedNumbers(searchNumber(x-1, y-1, split), &usedNumbers)
		multi *= checkUsedNumbers(searchNumber(x, y-1, split), &usedNumbers)
		multi *= checkUsedNumbers(searchNumber(x+1, y-1, split), &usedNumbers)
		multi *= checkUsedNumbers(searchNumber(x+1, y, split), &usedNumbers)
		multi *= checkUsedNumbers(searchNumber(x+1, y+1, split), &usedNumbers)
		multi *= checkUsedNumbers(searchNumber(x, y+1, split), &usedNumbers)
		multi *= checkUsedNumbers(searchNumber(x-1, y+1, split), &usedNumbers)
		multi *= checkUsedNumbers(searchNumber(x-1, y, split), &usedNumbers)
		if len(usedNumbers) > 2 {
			count += multi
		}
	}
	return strconv.Itoa(count), nil
}
