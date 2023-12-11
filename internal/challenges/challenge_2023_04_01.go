package challenges

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Challenge20230401 struct{}

func (Challenge20230401) GetYear() int {
	return 2023
}

func (Challenge20230401) GetDay() int {
	return 04
}

func (Challenge20230401) GetChallenge() int {
	return 01
}

func (Challenge20230401) Execute(rawFile string) error {
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

	for _, s := range strings.Split(rawFile, "\r\n") {
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
	log.Println(count)
	return nil
}
