package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"regexp"
	"strconv"
	"strings"
)

func NewChallenge201506() adventofcode.Challenge {
	return Challenge201506{}
}

type Challenge201506 struct{}

func (Challenge201506) GetYear() int {
	return 2015
}

func (Challenge201506) GetDay() int {
	return 06
}

func (c Challenge201506) ExecuteFirst(input string) (string, error) {
	lights := make([][]bool, 9999)
	for i, _ := range lights {
		lights[i] = make([]bool, 9999)
	}

	compile, err := regexp.Compile(`(.*) (\d{0,3}),(\d{0,3}) through (\d{0,3}),(\d{0,3})`)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(input, "\n") {
		submatch := compile.FindStringSubmatch(line)
		var get func(x int, y int) bool
		switch submatch[1] {
		case "turn on":
			get = func(x int, y int) bool {
				return true
			}
		case "turn off":
			get = func(x int, y int) bool {
				return false
			}
		case "toggle":
			get = func(x int, y int) bool {
				return !lights[y][x]
			}
		}
		startX, err := strconv.Atoi(submatch[2])
		if err != nil {
			return "", err
		}
		startY, err := strconv.Atoi(submatch[3])
		if err != nil {
			return "", err
		}
		endX, err := strconv.Atoi(submatch[4])
		if err != nil {
			return "", err
		}
		endY, err := strconv.Atoi(submatch[5])
		if err != nil {
			return "", err
		}
		for i := startY; i <= endY; i++ {
			for j := startX; j <= endX; j++ {
				lights[i][j] = get(j, i)
			}
		}
	}
	count := 0
	for _, lightRow := range lights {
		for _, light := range lightRow {
			if light {
				count++
			}
		}
	}
	return strconv.Itoa(count), nil
}

func (c Challenge201506) ExecuteSecond(input string) (string, error) {
	lights := make([][]int, 9999)
	for i, _ := range lights {
		lights[i] = make([]int, 9999)
	}

	compile, err := regexp.Compile(`(.*) (\d{0,3}),(\d{0,3}) through (\d{0,3}),(\d{0,3})`)
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(input, "\n") {
		submatch := compile.FindStringSubmatch(line)
		var get func(x int, y int) int
		switch submatch[1] {
		case "turn on":
			get = func(x int, y int) int {
				return 1
			}
		case "turn off":
			get = func(x int, y int) int {
				if lights[y][x] == 0 {
					return 0
				}
				return -1
			}
		case "toggle":
			get = func(x int, y int) int {
				return 2
			}
		}
		startX, err := strconv.Atoi(submatch[2])
		if err != nil {
			return "", err
		}
		startY, err := strconv.Atoi(submatch[3])
		if err != nil {
			return "", err
		}
		endX, err := strconv.Atoi(submatch[4])
		if err != nil {
			return "", err
		}
		endY, err := strconv.Atoi(submatch[5])
		if err != nil {
			return "", err
		}
		for i := startY; i <= endY; i++ {
			for j := startX; j <= endX; j++ {
				lights[i][j] += get(j, i)
			}
		}
	}
	count := 0
	for _, lightRow := range lights {
		for _, light := range lightRow {
			count += light
		}
	}
	return strconv.Itoa(count), nil
}
