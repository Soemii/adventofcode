package challenges

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Challenge20230101 struct{}

func (Challenge20230101) GetYear() int {
	return 2023
}

func (Challenge20230101) GetDay() int {
	return 01
}

func (Challenge20230101) GetChallenge() int {
	return 01
}

func (Challenge20230101) Execute(rawFile string) error {
	count := 0
	compile := regexp.MustCompile("[a-zA-Z]")
	for _, s := range strings.Split(rawFile, "\r\n") {
		allString := compile.ReplaceAllString(s, "")
		first, _ := strconv.Atoi(string(allString[0]))
		last, _ := strconv.Atoi(string(allString[len(allString)-1]))
		number := first*10 + last
		count += number
	}
	log.Println(count)
	return nil
}
