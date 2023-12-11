package challenges

import (
	"fmt"
	"log"
)

type Challenge20150302 struct{}

func (Challenge20150302) GetYear() int {
	return 2015
}

func (Challenge20150302) GetDay() int {
	return 03
}

func (Challenge20150302) GetChallenge() int {
	return 02
}

func (Challenge20150302) Execute(rawFile string) error {
	visited := make(map[string]int)
	currentX := 0
	currentRX := 0
	currentY := 0
	currentRY := 0
	visited["0:0"] = 1
	for i, v := range rawFile {
		switch v {
		case '^':
			if i%2 == 0 {
				currentY++
			} else {
				currentRY++
			}
		case 'v':
			if i%2 == 0 {
				currentY--
			} else {
				currentRY--
			}
		case '>':
			if i%2 == 0 {
				currentX++
			} else {
				currentRX++
			}
		case '<':
			if i%2 == 0 {
				currentX--
			} else {
				currentRX--
			}
		}
		if i%2 == 0 {
			c, ok := visited[fmt.Sprintf("%v:%v", currentX, currentY)]
			if ok {
				visited[fmt.Sprintf("%v:%v", currentX, currentY)] = c + 1
			} else {
				visited[fmt.Sprintf("%v:%v", currentX, currentY)] = 1
			}
		} else {
			c, ok := visited[fmt.Sprintf("%v:%v", currentRX, currentRY)]
			if ok {
				visited[fmt.Sprintf("%v:%v", currentRX, currentRY)] = c + 1
			} else {
				visited[fmt.Sprintf("%v:%v", currentRX, currentRY)] = 1
			}
		}

	}
	log.Println(len(visited))
	return nil
}
