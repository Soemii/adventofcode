package challenges

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

type Challenge20230402 struct{}

func (Challenge20230402) GetYear() int {
	return 2023
}

func (Challenge20230402) GetDay() int {
	return 04
}

func (Challenge20230402) GetChallenge() int {
	return 02
}

func (Challenge20230402) Execute(rawFile string) error {
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
	length := len(strings.Split(rawFile, "\r\n"))
	cardsAmount := make([]int, length)
	for i, _ := range cardsAmount {
		cardsAmount[i] = 1
	}

	for i, s := range strings.Split(rawFile, "\r\n") {
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
	log.Println(count)
	return nil
}
