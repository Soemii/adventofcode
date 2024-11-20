package challenges

import (
	"fmt"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"strconv"
)

func NewChallenge201503() adventofcode.Challenge {
	return Challenge201503{}
}

type Challenge201503 struct{}

func (c Challenge201503) ExecuteFirst(input string) (string, error) {
	visited := make(map[string]int)
	currentX := 0
	currentY := 0
	visited["0:0"] = 1
	for _, v := range input {
		switch v {
		case '^':
			currentY++
		case 'v':
			currentY--
		case '>':
			currentX++
		case '<':
			currentX--
		}
		i, ok := visited[fmt.Sprintf("%v:%v", currentX, currentY)]
		if ok {
			visited[fmt.Sprintf("%v:%v", currentX, currentY)] = i + 1
		} else {
			visited[fmt.Sprintf("%v:%v", currentX, currentY)] = 1
		}
	}
	return strconv.Itoa(len(visited)), nil
}

func (c Challenge201503) ExecuteSecond(input string) (string, error) {
	visited := make(map[string]int)
	currentX := 0
	currentRX := 0
	currentY := 0
	currentRY := 0
	visited["0:0"] = 1
	for i, v := range input {
		switch v {
		case '^':
			if i%2 == 0 {
				currentY++
			} else {
				currentRY++
			}
		case 'v':
			if i%2 == 0 {
				currentY--
			} else {
				currentRY--
			}
		case '>':
			if i%2 == 0 {
				currentX++
			} else {
				currentRX++
			}
		case '<':
			if i%2 == 0 {
				currentX--
			} else {
				currentRX--
			}
		}
		if i%2 == 0 {
			c, ok := visited[fmt.Sprintf("%v:%v", currentX, currentY)]
			if ok {
				visited[fmt.Sprintf("%v:%v", currentX, currentY)] = c + 1
			} else {
				visited[fmt.Sprintf("%v:%v", currentX, currentY)] = 1
			}
		} else {
			c, ok := visited[fmt.Sprintf("%v:%v", currentRX, currentRY)]
			if ok {
				visited[fmt.Sprintf("%v:%v", currentRX, currentRY)] = c + 1
			} else {
				visited[fmt.Sprintf("%v:%v", currentRX, currentRY)] = 1
			}
		}

	}
	return strconv.Itoa(len(visited)), nil
}

func (Challenge201503) GetYear() int {
	return 2015
}

func (Challenge201503) GetDay() int {
	return 03
}
