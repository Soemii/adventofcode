package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"math"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge202304() adventofcode.Challenge {
	return Challenge202304{}
}

type Challenge202304 struct{}

func (Challenge202304) GetYear() int {
	return 2023
}

func (Challenge202304) GetDay() int {
	return 04
}

func (c Challenge202304) ExecuteFirst(input string) (string, error) {
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

	count := 0

	for _, s := range strings.Split(input, "\n") {
		s = s[strings.Index(s, ": ")+1:]
		split := strings.Split(s, "|")
		myNumbers := stringToInt(strings.Split(split[0], " "))
		winningNumbers := stringToInt(strings.Split(split[1], " "))
		points := 0
		for _, number := range winningNumbers {
			if slices.Contains(myNumbers, number) {
				points += int(math.Max(float64(points), 1))
			}
		}
		count += points
	}
	return strconv.Itoa(count), nil
}

func (c Challenge202304) ExecuteSecond(input string) (string, error) {
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
	length := len(strings.Split(input, "\n"))
	cardsAmount := make([]int, length)
	for i := range cardsAmount {
		cardsAmount[i] = 1
	}

	for i, s := range strings.Split(input, "\n") {
		s = s[strings.Index(s, ": ")+1:]
		split := strings.Split(s, "|")
		myNumbers := stringToInt(strings.Split(split[0], " "))
		winningNumbers := stringToInt(strings.Split(split[1], " "))
		wins := 0
		for _, number := range winningNumbers {
			if slices.Contains(myNumbers, number) {
				wins++
			}
		}
		for j := i + 1; j < i+wins+1 && j < length; j++ {
			cardsAmount[j] += cardsAmount[i]
		}
	}
	count := 0
	for _, i := range cardsAmount {
		count += i
	}
	return strconv.Itoa(count), nil
}
