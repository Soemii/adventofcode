package challenges

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Challenge20230502 struct{}

func (Challenge20230502) GetYear() int {
	return 2023
}

func (Challenge20230502) GetDay() int {
	return 05
}

func (Challenge20230502) GetChallenge() int {
	return 02
}

func (Challenge20230502) Execute(rawFile string) error {
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
			lines := strings.Split(block, "\r\n")
			generatedBlocks[blockI] = make([][]int, len(lines)-1)
			for lineI, line := range lines[1:] {
				generatedBlocks[blockI][lineI] = stringToInt(strings.Split(line, " "))
			}
		}
		return generatedBlocks
	}

	split := strings.Split(rawFile, "\r\n\r\n")
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
		log.Println("Current smallest:", smallest)
	}
	log.Println("Finished", smallest)
	return nil
}
