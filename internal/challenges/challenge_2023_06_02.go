package challenges

import (
	"log"
	"strconv"
	"strings"
)

type Challenge20230602 struct{}

func (Challenge20230602) GetYear() int {
	return 2023
}

func (Challenge20230602) GetDay() int {
	return 06
}

func (Challenge20230602) GetChallenge() int {
	return 02
}

func (Challenge20230602) Execute(rawFile string) error {
	split := strings.Split(rawFile, "\r\n")
	time, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(split[0], "Time:", ""), " ", ""))
	recordDistance, _ := strconv.Atoi(strings.ReplaceAll(strings.ReplaceAll(split[1], "Distance:", ""), " ", ""))
	count := 0
	for j := 0; j < time; j++ {
		distance := (time - j) * j
		if distance > recordDistance {
			count++
		}
	}
	log.Println(count)
	return nil
}
