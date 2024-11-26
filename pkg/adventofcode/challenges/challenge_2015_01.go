package challenges

import (
	"errors"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"log"
	"strconv"
)

func NewChallenge201501() adventofcode.Challenge {
	return Challenge20150101{}
}

type Challenge20150101 struct{}

func (Challenge20150101) GetYear() int {
	return 2015
}

func (Challenge20150101) GetDay() int {
	return 01
}

func (c Challenge20150101) ExecuteFirst(input string) (string, error) {
	i := 0
	for _, r := range input {
		if r == '(' {
			i++
		} else {
			i--
		}
	}
	return strconv.Itoa(i), nil
}

func (c Challenge20150101) ExecuteSecond(input string) (string, error) {
	i := 0
	for j, r := range input {
		if r == '(' {
			i++
		} else {
			i--
		}
		if i < 0 {
			log.Println(j + 1)
			return strconv.Itoa(j + 1), nil
		}
	}
	return "", errors.New("never in the basement")
}
