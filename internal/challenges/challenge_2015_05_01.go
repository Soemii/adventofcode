package challenges

import (
	"log"
	"strings"
)

type Challenge20150501 struct{}

func (Challenge20150501) GetYear() int {
	return 2015
}

func (Challenge20150501) GetDay() int {
	return 05
}

func (Challenge20150501) GetChallenge() int {
	return 01
}

func (Challenge20150501) Execute(rawFile string) error {
	count := 0
	for _, s := range strings.Split(rawFile, "\r\n") {
		if strings.Contains(s, "ab") || strings.Contains(s, "cd") || strings.Contains(s, "pq") || strings.Contains(s, "xy") {
			continue
		}

		check := false

		c := 0
		for _, i := range s {
			if i == 'a' || i == 'e' || i == 'i' || i == 'o' || i == 'u' {
				c++
			}
		}
		if c < 3 {
			continue
		}

		before := ' '
		for _, r := range s {
			if r == before {
				check = true
			}
			before = r
		}
		if !check {
			continue
		}
		count++
	}
	log.Println(count)
	return nil
}
