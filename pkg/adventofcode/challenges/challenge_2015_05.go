package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"strconv"
	"strings"
)

func NewChallenge201505() adventofcode.Challenge {
	return Challenge201505{}
}

type Challenge201505 struct{}

func (Challenge201505) GetYear() int {
	return 2015
}

func (Challenge201505) GetDay() int {
	return 05
}

func (c Challenge201505) ExecuteFirst(input string) (string, error) {
	count := 0
	for _, s := range strings.Split(input, "\n") {
		if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
			continue
		}

		check := false

		c := 0
		for _, i := range s {
			if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' {
				c++
			}
		}
		if c < 3 {
			continue
		}

		before := ' '
		for _, r := range s {
			if r == before {
				check = true
			}
			before = r
		}
		if !check {
			continue
		}
		count++
	}
	return strconv.Itoa(count), nil
}

func (c Challenge201505) ExecuteSecond(input string) (string, error) {
	count := 0
	for _, s := range strings.Split(input, "\n") {
		check := false
		for i := 0; i < len(s)-2; i++ {
			if s[i+2] == s[i] {
				check = true
				break
			}
		}
		if !check {
			continue
		}

		check = false
		for i := 0; i < len(s)-1; i++ {
			current := s[i : i+2]
			newS := s[i+2:]
			if strings.Contains(newS, current) {
				check = true
				break
			}
		}
		if !check {
			continue
		}
		count++
	}
	return strconv.Itoa(count), nil
}
