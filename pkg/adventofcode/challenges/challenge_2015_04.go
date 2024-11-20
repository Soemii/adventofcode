package challenges

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"strconv"
	"strings"
)

func NewChallenge201504() adventofcode.Challenge {
	return Challenge201504{}
}

type Challenge201504 struct{}

func (Challenge201504) GetYear() int {
	return 2015
}

func (Challenge201504) GetDay() int {
	return 04
}

func (c Challenge201504) ExecuteFirst(input string) (string, error) {
	i := 0
	for {
		sum := md5.Sum([]byte(fmt.Sprintf("%v%v", input, i)))
		by := sum[:5]
		if strings.HasPrefix(hex.EncodeToString(by), "00000") {
			return strconv.Itoa(i), nil
		}
		i++
	}
}

func (c Challenge201504) ExecuteSecond(input string) (string, error) {
	i := 0
	for {
		sum := md5.Sum([]byte(fmt.Sprintf("%v%v", input, i)))
		by := sum[:5]
		if strings.HasPrefix(hex.EncodeToString(by), "000000") {
			return strconv.Itoa(i), nil
		}
		i++
	}
}
