package challenges

import (
	"fmt"
	"log"
	"strings"
)

type Challenge20231001 struct{}

func (Challenge20231001) GetYear() int {
	return 2023
}

func (Challenge20231001) GetDay() int {
	return 10
}

func (Challenge20231001) GetChallenge() int {
	return 01
}

func (Challenge20231001) Execute(rawFile string) error {
	type Location struct {
		X int
		Y int
	}
	//locationMap := map[string]Location{
	//	"|": {X: 0, Y: 1},
	//	"-": {X: 1, Y: 0},
	//	"L": {X: -1, Y: -1},
	//	"7": {X: 0, Y: 1},
	//	"J": {X: 0, Y: -1},
	//	"F": {X: -1, Y: 1},
	//}
	split := strings.Split(rawFile, "\r\n")
	getCharAt := func(x, y int) rune {
		if x < 0 || y < 0 {
			return ' '
		}
		if y >= len(split) || x >= len(split[0]) {
			return ' '
		}
		return rune(split[y][x])
	}
	sIndex := strings.IndexRune(strings.ReplaceAll(rawFile, "\r\n", ""), 'S')
	y := sIndex / len(split)
	x := sIndex % len(split)
	//locations := make([]Location, 0)
	//lastX := x
	//lastY := y
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if j == 0 && i == 0 {
				continue
			}
			fmt.Print(string(getCharAt(y+i, x+j)))
		}
		fmt.Println()
	}
	//for split[y][x] != 'S' {
	//
	//}
	log.Println(y, x)
	return nil
}
