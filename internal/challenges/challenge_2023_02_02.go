package challenges

import (
	"log"
	"strconv"
	"strings"
)

type Challenge20230202 struct{}

func (Challenge20230202) GetYear() int {
	return 2023
}

func (Challenge20230202) GetDay() int {
	return 02
}

func (Challenge20230202) GetChallenge() int {
	return 02
}

func (Challenge20230202) Execute(rawFile string) error {

	compare := func(s string, indexFrom string, min int) int {
		split := strings.Split(strings.Trim(s, " "), " ")
		if split[1] == indexFrom {
			number, _ := strconv.Atoi(split[0])
			if number > min {
				return number
			}
		}
		return min
	}

	count := 0
	rawFile = strings.ReplaceAll(rawFile, ";", ",")
	for _, s := range strings.Split(rawFile, "\r\n") {
		line := s[strings.Index(s, ": ")+1:]
		red := 0
		blue := 0
		green := 0
		for _, v := range strings.Split(line, ",") {
			red = compare(v, "red", red)
			green = compare(v, "green", green)
			blue = compare(v, "blue", blue)
		}
		count += red * green * blue
	}
	log.Println(count)
	return nil
}
