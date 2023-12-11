package challenges

import (
	"log"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Challenge20230702 struct{}

func (Challenge20230702) GetYear() int {
	return 2023
}

func (Challenge20230702) GetDay() int {
	return 07
}

func (Challenge20230702) GetChallenge() int {
	return 02
}

func (Challenge20230702) Execute(rawFile string) error {
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

	lines := strings.Split(rawFile, "\r\n")
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
	log.Println(count)
	return nil
}
