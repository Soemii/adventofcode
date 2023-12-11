package challenges

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Challenge20150202 struct{}

func (Challenge20150202) GetYear() int {
	return 2015
}

func (Challenge20150202) GetDay() int {
	return 02
}

func (Challenge20150202) GetChallenge() int {
	return 02
}

func (Challenge20150202) Execute(rawFile string) error {
	sum := 0
	for _, s := range strings.Split(rawFile, "\r\n") {
		split := strings.Split(s, "x")
		l, _ := strconv.Atoi(split[0])
		w, _ := strconv.Atoi(split[1])
		h, _ := strconv.Atoi(split[2])
		size := l * w * h
		min := int(math.Min(float64(l+w), math.Min(float64(w+h), float64(h+l)))) * 2
		sum += size + min
	}
	log.Println(sum)
	return nil
}
