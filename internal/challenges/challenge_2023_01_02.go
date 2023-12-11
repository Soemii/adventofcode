package challenges

import (
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Challenge20230102 struct{}

func (Challenge20230102) GetYear() int {
	return 2023
}

func (Challenge20230102) GetDay() int {
	return 01
}

func (Challenge20230102) GetChallenge() int {
	return 02
}

func (Challenge20230102) Execute(rawFile string) error {
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
	for _, s := range strings.Split(rawFile, "\r\n") {
		first := replaceStrings(compile.FindString(s))
		last := replaceStrings(reverseString(reverseCompile.FindString(reverseString(s))))

		count += first*10 + last
	}
	log.Println(count)
	return nil
}
