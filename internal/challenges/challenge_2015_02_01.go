package challenges

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Challenge20150201 struct{}

func (Challenge20150201) GetYear() int {
	return 2015
}

func (Challenge20150201) GetDay() int {
	return 02
}

func (Challenge20150201) GetChallenge() int {
	return 01
}

func (Challenge20150201) Execute(rawFile string) error {
	sum := 0
	for _, s := range strings.Split(rawFile, "\r\n") {
		split := strings.Split(s, "x")
		l, _ := strconv.Atoi(split[0])
		w, _ := strconv.Atoi(split[1])
		h, _ := strconv.Atoi(split[2])
		size := 2*l*w + 2*w*h + 2*h*l
		min := int(math.Min(float64(l*w), math.Min(float64(w*h), float64(h*l))))
		sum += size + min
	}
	log.Println(sum)
	return nil
}
