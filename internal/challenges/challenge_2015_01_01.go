package challenges

import "log"

type Challenge20150101 struct{}

func (Challenge20150101) GetYear() int {
	return 2015
}

func (Challenge20150101) GetDay() int {
	return 01
}

func (Challenge20150101) GetChallenge() int {
	return 01
}

func (Challenge20150101) Execute(rawFile string) error {
	i := 0
	for _, v := range rawFile {
		if v == '(' {
			i++
		} else {
			i--
		}
	}
	log.Println(i)
	return nil
}
