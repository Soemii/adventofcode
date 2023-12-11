package challenges

import (
	"log"
	"strconv"
	"strings"
)

type Challenge20230201 struct{}

func (Challenge20230201) GetYear() int {
	return 2023
}

func (Challenge20230201) GetDay() int {
	return 02
}

func (Challenge20230201) GetChallenge() int {
	return 01
}

func (Challenge20230201) Execute(rawFile string) error {
	redCubes := 12
	greenCubes := 13
	blueCubes := 14

	compare := func(s string, indexFrom string, max int) bool {
		split := strings.Split(strings.Trim(s, " "), " ")
		if split[1] == indexFrom {
			number, _ := strconv.Atoi(split[0])
			if number > max {
				return true
			}
		}
		return false
	}

	check := func(line string) bool {
		for _, v := range strings.Split(line, ",") {
			if compare(v, "red", redCubes) {
				return false
			}
			if compare(v, "green", greenCubes) {
				return false
			}
			if compare(v, "blue", blueCubes) {
				return false
			}
		}
		return true
	}

	count := 0
	rawFile = strings.ReplaceAll(rawFile, ";", ",")
	for i, s := range strings.Split(rawFile, "\r\n") {
		s = s[strings.Index(s, ": ")+1:]
		if check(s) {
			count += i + 1
		}
	}
	log.Println(count)
	return nil
}
