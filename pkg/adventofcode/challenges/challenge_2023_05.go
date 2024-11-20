package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"math"
	"sort"
	"strconv"
	"strings"
)

func NewChallenge202305() adventofcode.Challenge {
	return Challenge202305{}
}

type Challenge202305 struct{}

func (Challenge202305) GetYear() int {
	return 2023
}

func (Challenge202305) GetDay() int {
	return 05
}

func (c Challenge202305) ExecuteFirst(input string) (string, error) {
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

	split := strings.Split(input, "\n\n")
	seeds := stringToInt(strings.Split(split[0][7:], " "))
	for _, v := range split[1:] {
		m := mask(strings.Split(v, "\n")[1:], seeds)
		for i, seed := range seeds {
			if vs, ok := m[seed]; ok {
				seeds[i] = vs
			}
		}
	}
	sort.Ints(seeds)
	return strconv.Itoa(seeds[0]), nil
}

func (c Challenge202305) ExecuteSecond(input string) (string, error) {
	stringToInt := func(s []string) (ints []int) {
		ints = make([]int, len(s))
		for i, s2 := range s {
			atoi, _ := strconv.Atoi(strings.Trim(s2, " "))
			ints[i] = atoi
		}
		return
	}

	calc := func(s int, blocks [][][]int) int {
		seed := s
		for _, block := range blocks {
			for _, ranges := range block {
				if seed >= ranges[1] && seed < ranges[1]+ranges[2] {
					seed = ranges[0] + seed - ranges[1]
					break
				}
			}
		}
		return seed
	}

	generateBlocks := func(blocks []string) [][][]int {
		generatedBlocks := make([][][]int, len(blocks))
		for blockI, block := range blocks {
			lines := strings.Split(block, "\n")
			generatedBlocks[blockI] = make([][]int, len(lines)-1)
			for lineI, line := range lines[1:] {
				generatedBlocks[blockI][lineI] = stringToInt(strings.Split(line, " "))
			}
		}
		return generatedBlocks
	}

	split := strings.Split(input, "\n\n")
	seeds := stringToInt(strings.Split(split[0][7:], " "))

	blocks := generateBlocks(split[1:])
	smallestChannel := make(chan int)
	for i := 0; i < len(seeds); i += 2 {
		go func(channel chan int, start, length int) {
			mySmallest := math.MaxInt
			for j := 0; j < length; j++ {
				mySmallest = int(math.Min(float64(mySmallest), float64(calc(j+start, blocks))))
			}
			smallestChannel <- mySmallest
		}(smallestChannel, seeds[i], seeds[i+1])
	}
	smallest := math.MaxInt
	for i := 0; i < len(seeds)/2; i++ {
		sm := <-smallestChannel
		smallest = int(math.Min(float64(smallest), float64(sm)))
	}
	return strconv.Itoa(smallest), nil
}
