package challenges

import (
	"log"
	"strconv"
	"strings"
)

type Challenge20230601 struct{}

func (Challenge20230601) GetYear() int {
	return 2023
}

func (Challenge20230601) GetDay() int {
	return 06
}

func (Challenge20230601) GetChallenge() int {
	return 01
}

func (Challenge20230601) Execute(rawFile string) error {
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
	split := strings.Split(rawFile, "\r\n")
	times := stringToInt(strings.Split(strings.ReplaceAll(split[0], "Time:", ""), " "))
	distances := stringToInt(strings.Split(strings.ReplaceAll(split[1], "Distance:", ""), " "))
	count := 1
	for i := 0; i < len(times); i++ {
		countWays := 0
		for j := 0; j < times[i]; j++ {
			distance := (times[i] - j) * j
			if distance > distances[i] {
				countWays++
			}
		}
		count *= countWays
	}
	log.Println(count)
	return nil
}
