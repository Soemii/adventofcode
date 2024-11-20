package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge202309() adventofcode.Challenge {
	return Challenge202309{}
}

type Challenge202309 struct{}

func (Challenge202309) GetYear() int {
	return 2023
}

func (Challenge202309) GetDay() int {
	return 9
}

func (c Challenge202309) ExecuteFirst(input string) (string, error) {
	stringToInt := func(s []string) (ints []int) {
		ints = make([]int, 0)
		for _, s2 := range s {
			if strings.Trim(s2, " ") == "" {
				continue
			}
			atoi, _ := strconv.Atoi(strings.Trim(s2, " "))
			ints = append(ints, atoi)
		}
		return
	}
	allZero := func(ints []int) bool {
		for _, i := range ints {
			if i != 0 {
				return false
			}
		}
		return true
	}
	getForLine := func(line string) int {
		order := stringToInt(strings.Split(line, " "))
		orders := make([][]int, 0)
		orders = append(orders, order)
		for !allZero(orders[len(orders)-1]) {
			lastOrder := orders[len(orders)-1]
			newOrder := make([]int, 0)
			before := lastOrder[0]
			for i := 1; i < len(lastOrder); i++ {
				current := lastOrder[i]
				diff := current - before
				before = current
				newOrder = append(newOrder, diff)
			}
			orders = append(orders, newOrder)
		}
		slices.Reverse(orders)
		last := 0
		for _, ints := range orders[1:] {
			i := ints[len(ints)-1]
			last = i + last
		}
		return last
	}
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		sum += getForLine(s)
	}
	return strconv.Itoa(sum), nil
}

func (c Challenge202309) ExecuteSecond(input string) (string, error) {
	stringToInt := func(s []string) (ints []int) {
		ints = make([]int, 0)
		for _, s2 := range s {
			if strings.Trim(s2, " ") == "" {
				continue
			}
			atoi, _ := strconv.Atoi(strings.Trim(s2, " "))
			ints = append(ints, atoi)
		}
		return
	}
	allZero := func(ints []int) bool {
		for _, i := range ints {
			if i != 0 {
				return false
			}
		}
		return true
	}
	getForLine := func(line string) int {
		order := stringToInt(strings.Split(line, " "))
		orders := make([][]int, 0)
		orders = append(orders, order)
		for !allZero(orders[len(orders)-1]) {
			lastOrder := orders[len(orders)-1]
			newOrder := make([]int, 0)
			before := lastOrder[0]
			for i := 1; i < len(lastOrder); i++ {
				current := lastOrder[i]
				diff := current - before
				before = current
				newOrder = append(newOrder, diff)
			}
			orders = append(orders, newOrder)
		}
		slices.Reverse(orders)
		last := 0
		for _, ints := range orders[1:] {
			i := ints[0]
			last = i - last
		}
		return last
	}
	sum := 0
	for _, s := range strings.Split(input, "\n") {
		sum += getForLine(s)
	}
	return strconv.Itoa(sum), nil
}
