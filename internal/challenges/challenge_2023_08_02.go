package challenges

import (
	"log"
	"regexp"
	"strings"
)

type Challenge20230802 struct{}

func (Challenge20230802) GetYear() int {
	return 2023
}

func (Challenge20230802) GetDay() int {
	return 8
}

func (Challenge20230802) GetChallenge() int {
	return 02
}

func (Challenge20230802) Execute(rawFile string) error {
	GCD := func(a, b int) int {
		for b != 0 {
			t := b
			b = a % b
			a = t
		}
		return a
	}
	LCM := func(integers ...int) int {
		a := integers[0]
		b := integers[1]
		result := a * b / GCD(a, b)

		for i := 2; i < len(integers); i++ {
			result = result * integers[i] / GCD(result, integers[i])
		}

		return result
	}
	type Values struct {
		Left  string
		Right string
	}
	lines := strings.Split(rawFile, "\r\n")
	order := lines[0]
	lines = lines[2:]
	locationMap := make(map[string]Values)

	compile, err := regexp.Compile("(\\w\\w\\w) = \\((\\w\\w\\w), (\\w\\w\\w)\\)")
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
	startingLocation := make([]string, 0)
	for s, _ := range locationMap {
		if strings.HasSuffix(s, "A") {
			startingLocation = append(startingLocation, s)
		}
	}
	endSteps := make([]int, len(startingLocation))
	for locI, currentLocation := range startingLocation {
		i := 0
		for {
			orderIndex := i % len(order)
			direction := order[orderIndex]
			if direction == 'R' {
				currentLocation = locationMap[currentLocation].Right
			} else {
				currentLocation = locationMap[currentLocation].Left
			}
			i++
			if strings.HasSuffix(currentLocation, "Z") {
				endSteps[locI] = i
				break
			}
		}
	}
	log.Println(LCM(endSteps...))
	return nil
}

/*
	allOnZ := func(s []string) bool {
		for _, s2 := range s {
			if s2[2] != 'Z' {
				return false
			}
		}
		return true
	}
	type Values struct {
		Left  string
		Right string
	}
	lines := strings.Split(rawFile, "\r\n")
	order := lines[0]
	lines = lines[2:]
	locationMap := make(map[string]Values)

	compile, err := regexp.Compile("(\\w\\w\\w) = \\((\\w\\w\\w), (\\w\\w\\w)\\)")
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
	startingLocation := make([]string, 0)
	for s, _ := range locationMap {
		if strings.HasSuffix(s, "A") {
			startingLocation = append(startingLocation, s)
		}
	}
	i := 0
	for {
		orderIndex := i % len(order)
		direction := order[orderIndex]
		for i2, s := range startingLocation {
			if direction == 'R' {
				startingLocation[i2] = locationMap[s].Right
			} else {
				startingLocation[i2] = locationMap[s].Left
			}
		}
		i++
		if allOnZ(startingLocation) {
			break
		}
	}
	log.Println(i)
	return nil
*/
