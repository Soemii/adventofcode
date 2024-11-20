package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"log"
	"strconv"
	"strings"
)

func NewChallenge202310() adventofcode.Challenge {
	return Challenge202310{}
}

type Challenge202310 struct{}

func (Challenge202310) GetYear() int {
	return 2023
}

func (Challenge202310) GetDay() int {
	return 10
}

func (c Challenge202310) ExecuteFirst(input string) (string, error) {
	type Location struct {
		X int
		Y int
	}
	locationMap := map[rune][]Location{
		'|': {{X: 0, Y: 1}, {X: 0, Y: -1}},
		'-': {{X: 1, Y: 0}, {X: -1, Y: 0}},
		'L': {{X: 0, Y: -1}, {X: 1, Y: 0}},
		'7': {{X: 0, Y: 1}, {X: -1, Y: 0}},
		'J': {{X: 0, Y: -1}, {X: -1, Y: 0}},
		'F': {{X: 1, Y: 0}, {X: 0, Y: 1}},
	}
	split := strings.Split(rawFile, "\n")
	getCharAt := func(location Location) rune {
		if location.X < 0 || location.Y < 0 {
			return ' '
		}
		if location.Y >= len(split) || location.X >= len(split[0]) {
			return ' '
		}
		return rune(split[location.Y][location.X])
	}
	getConnectedLocations := func(char rune, location Location) (Location, Location) {
		l, ok := locationMap[char]
		if !ok {
			return Location{}, Location{}
		}
		firstLocation := Location{
			X: location.X + l[0].X,
			Y: location.Y + l[0].Y,
		}
		secondLocation := Location{
			X: location.X + l[1].X,
			Y: location.Y + l[1].Y,
		}
		return firstLocation, secondLocation
	}
	sIndex := strings.IndexRune(strings.ReplaceAll(rawFile, "\n", ""), 'S')
	y := sIndex / len(split[0])
	x := sIndex % len(split[0])
	sLocation := Location{x, y}
	var startingLocation Location
	for _, v := range []Location{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		currentLocation := Location{x + v.X, y + v.Y}
		char := getCharAt(currentLocation)
		if char != ' ' && char != '.' {
			l, l2 := getConnectedLocations(char, currentLocation)
			if l == sLocation {
				startingLocation = currentLocation
				break
			} else if l2 == sLocation {
				startingLocation = currentLocation
				break
			}
		}
	}
	locations := make([]Location, 0)
	locations = append(locations, startingLocation)
	lastLocation := sLocation
	for getCharAt(startingLocation) != 'S' {
		r := getCharAt(startingLocation)
		l, l2 := getConnectedLocations(r, startingLocation)
		var newLocation Location
		if l != lastLocation {
			newLocation = l
		} else if l2 != lastLocation {
			newLocation = l2
		}
		lastLocation = startingLocation
		startingLocation = newLocation
		locations = append(locations, newLocation)
	}
	return strconv.Itoa(len(locations) / 2), nil
}

func (c Challenge202310) ExecuteSecond(input string) (string, error) {
	return "Not Implemented", nil
}
