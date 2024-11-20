package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"math"
	"strconv"
	"strings"
)

func NewChallenge201502() adventofcode.Challenge {
	return Challenge201502{}
}

type Challenge201502 struct{}

func (c Challenge201502) ExecuteFirst(input string) (string, error) {
	sum := 0
	for _, s := range strings.Split(input, "\r\n") {
		split := strings.Split(s, "x")
		l, err := strconv.Atoi(split[0])
		if err != nil {
			return "", err
		}
		w, err := strconv.Atoi(split[1])
		if err != nil {
			return "", err
		}
		h, err := strconv.Atoi(split[2])
		if err != nil {
			return "", err
		}
		size := 2*l*w + 2*w*h + 2*h*l
		min := int(math.Min(float64(l*w), math.Min(float64(w*h), float64(h*l))))
		sum += size + min
	}
	return strconv.Itoa(sum), nil
}

func (c Challenge201502) ExecuteSecond(input string) (string, error) {
	sum := 0
	for _, s := range strings.Split(input, "\r\n") {
		split := strings.Split(s, "x")
		l, err := strconv.Atoi(split[0])
		if err != nil {
			return "", err
		}
		w, err := strconv.Atoi(split[1])
		if err != nil {
			return "", err
		}
		h, err := strconv.Atoi(split[2])
		if err != nil {
			return "", err
		}
		size := l * w * h
		min := int(math.Min(float64(l+w), math.Min(float64(w+h), float64(h+l)))) * 2
		sum += size + min
	}
	return strconv.Itoa(sum), nil
}

func (Challenge201502) GetYear() int {
	return 2015
}

func (Challenge201502) GetDay() int {
	return 02
}
