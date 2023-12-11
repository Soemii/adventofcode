package challenges

import (
	"log"
	"strings"
)

type Challenge20150502 struct{}

func (Challenge20150502) GetYear() int {
	return 2015
}

func (Challenge20150502) GetDay() int {
	return 05
}

func (Challenge20150502) GetChallenge() int {
	return 02
}

func (Challenge20150502) Execute(rawFile string) error {
	count := 0
	for _, s := range strings.Split(rawFile, "\r\n") {
		check := false
		for i := 0; i < len(s)-2; i++ {
			if s[i+2] == s[i] {
				check = true
				break
			}
		}
		if !check { continue }

		check = false
		for i := 0; i < len(s)-1; i++ {
			current := s[i : i+2]
			newS := s[i+2:]
			if strings.Contains(newS, current) {
				check = true
				break
			}
		}
		if !check { continue }
		count++
	}
	log.Println(count)
	return nil
}
