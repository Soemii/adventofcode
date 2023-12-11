package challenges

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

type Challenge20230501 struct{}

func (Challenge20230501) GetYear() int {
	return 2023
}

func (Challenge20230501) GetDay() int {
	return 05
}

func (Challenge20230501) GetChallenge() int {
	return 01
}

func (Challenge20230501) Execute(rawFile string) error {
	stringToInt := func(s []string) (ints []int) {
		ints = make([]int, len(s))
		for i, s2 := range s {
			atoi, _ := strconv.Atoi(strings.Trim(s2, " "))
			ints[i] = atoi
		}
		return
	}

	mask := func(s []string, seeds []int) map[int]int {
		m := make(map[int]int)
		for _, s2 := range s {
			ints := stringToInt(strings.Split(s2, " "))
			for _, seed := range seeds {
				if seed > ints[1] && seed < ints[1]+ints[2] {
					m[seed] = ints[0] + seed - ints[1]
				}
			}
		}
		return m
	}

	split := strings.Split(rawFile, "\r\n\r\n")
	seeds := stringToInt(strings.Split(split[0][7:], " "))
	for _, v := range split[1:] {
		m := mask(strings.Split(v, "\r\n")[1:], seeds)
		for i, seed := range seeds {
			if vs, ok := m[seed]; ok {
				seeds[i] = vs
			}
		}
	}
	sort.Ints(seeds)
	log.Println(seeds[0])
	return nil
}
