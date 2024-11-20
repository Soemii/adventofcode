package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"strconv"
	"strings"
)

func NewChallenge202306() adventofcode.Challenge {
	return Challenge202306{}
}

type Challenge202306 struct{}

func (Challenge202306) GetYear() int {
	return 2023
}

func (Challenge202306) GetDay() int {
	return 06
}

func (c Challenge202306) ExecuteFirst(input string) (string, error) {
	stringToInt := func(s []string) (ints []int) {
		ints = make([]int, 0)
		for _, s2 := range s {
			if strings.Trim(s2, " ") == "" {
				continue
			}
			atoi, _ := strconv.Atoi(strings.Trim(s2, " "))
			ints = append(ints, atoi)
		}
		return
	}
	split := strings.Split(input, "\n")
	times := stringToInt(strings.Split(strings.ReplaceAll(split[0], "Time:", ""), " "))
	distances := stringToInt(strings.Split(strings.ReplaceAll(split[1], "Distance:", ""), " "))
	count := 1
	for i := 0; i < len(times); i++ {
		countWays := 0
		for j := 0; j < times[i]; j++ {
			distance := (times[i] - j) * j
			if distance > distances[i] {
				countWays++
			}
		}
		count *= countWays
	}
	return strconv.Itoa(count), nil
}

func (c Challenge202306) ExecuteSecond(input string) (string, error) {
	split := strings.Split(input, "\n")
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(split[0], "Time:", ""), " ", ""))
	recordDistance, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(split[1], "Distance:", ""), " ", ""))
	count := 0
	for j := 0; j < time; j++ {
		distance := (time - j) * j
		if distance > recordDistance {
			count++
		}
	}
	return strconv.Itoa(count), nil
}
