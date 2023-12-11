package challenges

import (
	"errors"
	"log"
)

type Challenge20150102 struct{}

func (Challenge20150102) GetYear() int {
	return 2015
}

func (Challenge20150102) GetDay() int {
	return 01
}

func (Challenge20150102) GetChallenge() int {
	return 02
}

func (Challenge20150102) Execute(rawFile string) error {
	i := 0
	for j, v := range rawFile {
		if v == '(' {
			i++
		} else {
			i--
		}
		if i < 0 {
			log.Println(j + 1)
			return nil
		}
	}
	return errors.New("never in the basement")
}
