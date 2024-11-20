package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"regexp"
	"strconv"
	"strings"
)

func NewChallenge202308() adventofcode.Challenge {
	return Challenge202308{}
}

type Challenge202308 struct{}

func (Challenge202308) GetYear() int {
	return 2023
}

func (Challenge202308) GetDay() int {
	return 8
}

func (c Challenge202308) ExecuteFirst(input string) (string, error) {
	type Values struct {
		Left  string
		Right string
	}
	lines := strings.Split(input, "\n")
	order := lines[0]
	lines = lines[2:]
	locationMap := make(map[string]Values)

	compile, err := regexp.Compile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)
	if err != nil {
		return "", err
	}
	for _, line := range lines {
		submatch := compile.FindStringSubmatch(line)
		locationMap[submatch[1]] = Values{
			Left:  submatch[2],
			Right: submatch[3],
		}
	}
	i := 0
	currentLocation := "AAA"
	for {
		orderIndex := i % len(order)
		direction := order[orderIndex]
		if direction == 'R' {
			currentLocation = locationMap[currentLocation].Right
		} else {
			currentLocation = locationMap[currentLocation].Left
		}
		i++
		if currentLocation == "ZZZ" {
			break
		}
	}
	return strconv.Itoa(i), err
}

func (c Challenge202308) ExecuteSecond(input string) (string, error) {
	GCD := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	LCM := func(integers ...int) int {
		a := integers[0]
		b := integers[1]
		result := a * b / GCD(a, b)

		for i := 2; i < len(integers); i++ {
			result = result * integers[i] / GCD(result, integers[i])
		}

		return result
	}
	type Values struct {
		Left  string
		Right string
	}
	lines := strings.Split(input, "\n")
	order := lines[0]
	lines = lines[2:]
	locationMap := make(map[string]Values)

	compile, err := regexp.Compile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)
	if err != nil {
		return "", err
	}
	for _, line := range lines {
		submatch := compile.FindStringSubmatch(line)
		locationMap[submatch[1]] = Values{
			Left:  submatch[2],
			Right: submatch[3],
		}
	}
	startingLocation := make([]string, 0)
	for s := range locationMap {
		if strings.HasSuffix(s, "A") {
			startingLocation = append(startingLocation, s)
		}
	}
	endSteps := make([]int, len(startingLocation))
	for locI, currentLocation := range startingLocation {
		i := 0
		for {
			orderIndex := i % len(order)
			direction := order[orderIndex]
			if direction == 'R' {
				currentLocation = locationMap[currentLocation].Right
			} else {
				currentLocation = locationMap[currentLocation].Left
			}
			i++
			if strings.HasSuffix(currentLocation, "Z") {
				endSteps[locI] = i
				break
			}
		}
	}
	return strconv.Itoa(LCM(endSteps...)), nil
}
