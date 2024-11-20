package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"math"
	"strconv"
	"strings"
)

func NewChallenge201502() adventofcode.Challenge[[]string, []string] {
	return Challenge201502{}
}

type Challenge201502 struct{}

func (c Challenge201502) PrepareFirst(rawfile string) []string {
	return strings.Split(rawfile, "\r\n")
}

func (c Challenge201502) PrepareSecond(rawfile string) []string {
	return c.PrepareFirst(rawfile)
}

func (c Challenge201502) ExecuteFirst(input []string) (string, error) {
	sum := 0
	for _, s := range input {
		split := strings.Split(s, "x")
		l, _ := strconv.Atoi(split[0])
		w, _ := strconv.Atoi(split[1])
		h, _ := strconv.Atoi(split[2])
		size := 2*l*w + 2*w*h + 2*h*l
		min := int(math.Min(float64(l*w), math.Min(float64(w*h), float64(h*l))))
		sum += size + min
	}
	return strconv.Itoa(sum), nil
}

func (c Challenge201502) ExecuteSecond(input []string) (string, error) {
	sum := 0
	for _, s := range input {
		split := strings.Split(s, "x")
		l, _ := strconv.Atoi(split[0])
		w, _ := strconv.Atoi(split[1])
		h, _ := strconv.Atoi(split[2])
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
