package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"strconv"
	"strings"
)

func NewChallenge202302() adventofcode.Challenge {
	return Challenge202302{}
}

type Challenge202302 struct{}

func (Challenge202302) GetYear() int {
	return 2023
}

func (Challenge202302) GetDay() int {
	return 02
}

func (c Challenge202302) ExecuteFirst(input string) (string, error) {
	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	compare := func(s string, indexFrom string, max int) bool {
		split := strings.Split(strings.Trim(s, " "), " ")
		if split[1] == indexFrom {
			number, _ := strconv.Atoi(split[0])
			if number > max {
				return true
			}
		}
		return false
	}

	check := func(line string) bool {
		for _, v := range strings.Split(line, ",") {
			if compare(v, "red", redCubes) {
				return false
			}
			if compare(v, "green", greenCubes) {
				return false
			}
			if compare(v, "blue", blueCubes) {
				return false
			}
		}
		return true
	}

	count := 0
	input = strings.ReplaceAll(input, ";", ",")
	for i, s := range strings.Split(input, "\n") {
		s = s[strings.Index(s, ": ")+1:]
		if check(s) {
			count += i + 1
		}
	}
	return strconv.Itoa(count), nil
}

func (c Challenge202302) ExecuteSecond(input string) (string, error) {
	compare := func(s string, indexFrom string, min int) int {
		split := strings.Split(strings.Trim(s, " "), " ")
		if split[1] == indexFrom {
			number, _ := strconv.Atoi(split[0])
			if number > min {
				return number
			}
		}
		return min
	}

	count := 0
	input = strings.ReplaceAll(input, ";", ",")
	for _, s := range strings.Split(input, "\n") {
		line := s[strings.Index(s, ": ")+1:]
		red := 0
		blue := 0
		green := 0
		for _, v := range strings.Split(line, ",") {
			red = compare(v, "red", red)
			green = compare(v, "green", green)
			blue = compare(v, "blue", blue)
		}
		count += red * green * blue
	}
	return strconv.Itoa(count), nil
}
