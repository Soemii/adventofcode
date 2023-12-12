package challenges

import (
	"log"
	"regexp"
	"strings"
)

type Challenge20230801 struct{}

func (Challenge20230801) GetYear() int {
	return 2023
}

func (Challenge20230801) GetDay() int {
	return 8
}

func (Challenge20230801) GetChallenge() int {
	return 01
}

func (Challenge20230801) Execute(rawFile string) error {
	type Values struct {
		Left  string
		Right string
	}
	lines := strings.Split(rawFile, "\r\n")
	order := lines[0]
	lines = lines[2:]
	locationMap := make(map[string]Values)

	compile, err := regexp.Compile(`(\w\w\w) = \((\w\w\w), (\w\w\w)\)`)
	if err != nil {
		return err
	}
	for _, line := range lines {
		submatch := compile.FindStringSubmatch(line)
		locationMap[submatch[1]] = Values{
			Left:  submatch[2],
			Right: submatch[3],
		}
	}
	i := 0
	currentLocation := "AAA"
	for {
		orderIndex := i % len(order)
		direction := order[orderIndex]
		if direction == 'R' {
			currentLocation = locationMap[currentLocation].Right
		} else {
			currentLocation = locationMap[currentLocation].Left
		}
		i++
		if currentLocation == "ZZZ" {
			break
		}
	}
	log.Println(i)
	return nil
}
