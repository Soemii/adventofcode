package challenges

import (
	"fmt"
	"log"
)

type Challenge20150301 struct{}

func (Challenge20150301) GetYear() int {
	return 2015
}

func (Challenge20150301) GetDay() int {
	return 03
}

func (Challenge20150301) GetChallenge() int {
	return 01
}

func (Challenge20150301) Execute(rawFile string) error {
	visited := make(map[string]int)
	currentX := 0
	currentY := 0
	visited["0:0"] = 1
	for _, v := range rawFile {
		switch v {
		case '^':
			currentY++
		case 'v':
			currentY--
		case '>':
			currentX++
		case '<':
			currentX--
		}
		i, ok := visited[fmt.Sprintf("%v:%v", currentX, currentY)]
		if ok {
			visited[fmt.Sprintf("%v:%v", currentX, currentY)] = i + 1
		} else {
			visited[fmt.Sprintf("%v:%v", currentX, currentY)] = 1
		}
	}
	log.Println(len(visited))
	return nil
}
