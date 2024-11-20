package challenges

import (
	"github.com/Soemii/AdventOfCode/pkg/adventofcode"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func NewChallenge202307() adventofcode.Challenge {
	return Challenge202307{}
}

type Challenge202307 struct{}

func (Challenge202307) GetYear() int {
	return 2023
}

func (Challenge202307) GetDay() int {
	return 07
}

func (c Challenge202307) ExecuteFirst(input string) (string, error) {
	type Hand struct {
		Cards   []rune
		Bid     int
		Ranking int
	}

	cardOrder := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

	fiveOfAKind := func(m map[rune]int) bool {
		for _, i := range m {
			if i == 5 {
				return true
			}
		}
		return false
	}
	fourOfAKind := func(m map[rune]int) bool {
		for _, i := range m {
			if i == 4 {
				return true
			}
		}
		return false
	}
	fullHouse := func(m map[rune]int) bool {
		three := false
		two := false
		for _, i := range m {
			if i == 3 {
				three = true
			}
			if i == 2 {
				two = true
			}
		}
		return three && two
	}
	threeOfAKind := func(m map[rune]int) bool {
		for _, i := range m {
			if i == 3 {
				return true
			}
		}
		return false
	}
	twoPair := func(m map[rune]int) bool {
		count := 0
		for _, i := range m {
			if i == 2 {
				count++
			}
		}
		return count == 2
	}
	onePair := func(m map[rune]int) bool {
		for _, i := range m {
			if i == 2 {
				return true
			}
		}
		return false
	}
	highCard := func(m map[rune]int) bool {
		return true
	}

	calcRanking := func(cards []rune) int {
		cardmap := make(map[rune]int)
		for _, card := range cards {
			cardmap[card] = cardmap[card] + 1
		}
		checks := []func(map[rune]int) bool{
			fiveOfAKind,
			fourOfAKind,
			fullHouse,
			threeOfAKind,
			twoPair,
			onePair,
			highCard,
		}
		ranking := 0
		for i, check := range checks {
			if check(cardmap) {
				ranking = (len(checks) - i) * 1_00_00_00_00_00
				break
			}
		}
		for i, card := range cards {
			a := 1
			for j := 0; j < len(cards)-i-1; j++ {
				a *= 100
			}
			a *= len(cardOrder) - slices.Index(cardOrder, card)
			ranking += a
		}
		return ranking
	}

	lines := strings.Split(input, "\n")
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		cards := []rune(split[0])
		bid, _ := strconv.Atoi(split[1])
		hands[i] = Hand{
			Cards:   cards,
			Bid:     bid,
			Ranking: calcRanking(cards),
		}
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Ranking < hands[j].Ranking
	})
	count := 0
	for i, hand := range hands {
		count += hand.Bid * (i + 1)
	}

	return strconv.Itoa(count), nil
}

func (c Challenge202307) ExecuteSecond(input string) (string, error) {
	type Hand struct {
		Cards   []rune
		Bid     int
		Ranking int
	}
	cardOrder := []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}
	fiveOfAKind := func(m map[rune]int, joker int) bool {
		if joker == 5 {
			return true
		}
		for _, i := range m {
			if i+joker == 5 {
				return true
			}
		}
		return false
	}
	fourOfAKind := func(m map[rune]int, joker int) bool {
		for _, i := range m {
			if i+joker == 4 {
				return true
			}
		}
		return false
	}
	fullHouse := func(m map[rune]int, joker int) bool {
		three := false
		two := false
		for _, i := range m {
			if i+joker == 3 {
				joker -= 3 - i
				three = true
				continue
			}
			if i+joker == 2 {
				joker -= 2 - i
				two = true
				continue
			}
		}
		return three && two
	}
	threeOfAKind := func(m map[rune]int, joker int) bool {
		for _, i := range m {
			if i+joker == 3 {
				return true
			}
		}
		return false
	}
	twoPair := func(m map[rune]int, joker int) bool {
		count := 0
		for _, i := range m {
			if i+joker == 2 {
				joker -= 2 - i
				count++
			}
		}
		return count == 2
	}
	onePair := func(m map[rune]int, joker int) bool {
		for _, i := range m {
			if i+joker == 2 {
				return true
			}
		}
		return false
	}
	highCard := func(m map[rune]int, joker int) bool {
		return true
	}

	calcRanking := func(cards []rune) int {
		cardmap := make(map[rune]int)
		for _, card := range cards {
			cardmap[card] = cardmap[card] + 1
		}
		checks := []func(map[rune]int, int) bool{
			fiveOfAKind,
			fourOfAKind,
			fullHouse,
			threeOfAKind,
			twoPair,
			onePair,
			highCard,
		}
		ranking := 0
		j := cardmap['J']
		delete(cardmap, 'J')
		for i, check := range checks {
			if check(cardmap, j) {
				ranking = (len(checks) - i) * 1_00_00_00_00_00
				break
			}
		}
		for i, card := range cards {
			a := 1
			for j := 0; j < len(cards)-i-1; j++ {
				a *= 100
			}
			a *= len(cardOrder) - slices.Index(cardOrder, card)
			ranking += a
		}
		return ranking
	}

	lines := strings.Split(input, "\n")
	hands := make([]Hand, len(lines))
	for i, line := range lines {
		split := strings.Split(line, " ")
		cards := []rune(split[0])
		bid, _ := strconv.Atoi(split[1])
		hands[i] = Hand{
			Cards:   cards,
			Bid:     bid,
			Ranking: calcRanking(cards),
		}
	}
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Ranking < hands[j].Ranking
	})
	count := 0
	for i, hand := range hands {
		count += hand.Bid * (i + 1)
	}
	return strconv.Itoa(count), nil
}
