package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func NewChallenge202301() adventofcode.Challenge {
	return Challenge202301{}
}

type Challenge202301 struct{}

func (Challenge202301) GetYear() int {
	return 2023
}

func (Challenge202301) GetDay() int {
	return 01
}

func (c Challenge202301) ExecuteFirst(input string) (string, error) {
	count := 0
	compile := regexp.MustCompile("[a-zA-Z]")
	for _, s := range strings.Split(input, "\n") {
		allString := compile.ReplaceAllString(s, "")
		first, err := strconv.Atoi(string(allString[0]))
		if err != nil {
			return "", err
		}
		last, err := strconv.Atoi(string(allString[len(allString)-1]))
		if err != nil {
			return "", err
		}
		number := first*10 + last
		count += number
	}
	return strconv.Itoa(count), nil
}

func (c Challenge202301) ExecuteSecond(input string) (string, error) {
	words := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	reverseString := func(string2 string) string {
		b := []byte(string2)
		slices.Reverse(b)
		return string(b)
	}

	replaceStrings := func(string2 string) int {
		i, ok := words[string2]
		if ok {
			return i
		}
		atoi, _ := strconv.Atoi(string2)
		return atoi
	}

	regex := "[1-9]|(four)|(six)|(eight)|(two)|(three)|(five)|(seven)|(nine)|(one)"
	reverseRegex := "[1-9]|(enin)|(eno)|(eerht)|(evif)|(neves)|(owt)|(ruof)|(xis)|(thgie)"
	count := 0
	compile := regexp.MustCompile(regex)
	reverseCompile := regexp.MustCompile(reverseRegex)
	for _, s := range strings.Split(input, "\n") {
		first := replaceStrings(compile.FindString(s))
		last := replaceStrings(reverseString(reverseCompile.FindString(reverseString(s))))

		count += first*10 + last
	}
	return strconv.Itoa(count), nil
}
