package main

import (
	"github.com/lillrurre/adventofcode-2023/util"
	"slices"
	"strings"
)

type game struct {
	cards string
	bid   int
	hand  hand
}

type hand int

const (
	HighCard hand = iota + 1
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func main() {
	input := util.FileAsStringArr(7, "\n")

	gamesPart1, gamesPart2 := make([]*game, 0), make([]*game, 0)
	for _, line := range input {
		fields := strings.Fields(line)
		cards, bid := fields[0], util.Atoi(fields[1])

		cardMap, jokers := getCardMap(cards, false)
		h := getBestHand(cardMap, jokers)
		gamesPart1 = append(gamesPart1, &game{cards: cards, bid: bid, hand: h})

		cardMap, jokers = getCardMap(cards, true)
		h = getBestHand(cardMap, jokers)
		gamesPart2 = append(gamesPart2, &game{cards: cards, bid: bid, hand: h})
	}

	solve := func(games []*game, valueMap map[rune]int) (sum int) {
		slices.SortFunc(games, func(a, b *game) int {
			if a.hand != b.hand {
				return int(a.hand - b.hand)
			}
			for i, card := range a.cards {
				if card != rune(b.cards[i]) {
					return valueMap[card] - valueMap[rune(b.cards[i])]
				}
			}
			return 1
		})
		for rank, g := range games {
			sum += (rank + 1) * g.bid
		}
		return sum
	}

	var cardValueMap = map[rune]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}

	util.Run(1, func() any {
		return solve(gamesPart1, cardValueMap)
	})

	cardValueMap['J'] = 1
	util.Run(2, func() any {
		return solve(gamesPart2, cardValueMap)
	})

}

func getCardMap(cards string, part2 bool) (cardMap map[int]int, jokers int) {
	cardMap = make(map[int]int)
	for _, card := range cards {
		if part2 && card == 'J' {
			jokers++
			continue
		}
		cardMap[int(card)]++
	}
	return cardMap, jokers
}

func getBestHand(cardMap map[int]int, jokers int) hand {
	for card, amount := range cardMap {
		switch {
		case amount == 5:
			return FiveOfAKind
		case amount == 4 && jokers == 1:
			return FiveOfAKind
		case amount == 4:
			return FourOfAKind
		case amount == 3 && jokers == 2:
			return FiveOfAKind
		case amount == 3 && jokers == 1:
			return FourOfAKind
		case amount == 2 && jokers == 3:
			return FiveOfAKind
		case amount == 2 && jokers == 2:
			return FourOfAKind

		case amount == 3:
			for otherCard, otherAmount := range cardMap {
				switch {
				case card == otherCard:
					continue
				case otherAmount == 2:
					return FullHouse
				}
			}
			return ThreeOfAKind

		case amount == 2:
			for otherCard, otherAmount := range cardMap {
				switch {
				case card == otherCard:
					continue
				case otherAmount == 2:
					if jokers == 1 {
						return FullHouse
					}
					return TwoPairs
				case otherAmount == 3:
					return FullHouse
				}
			}
			if jokers == 1 {
				return ThreeOfAKind
			}
			return OnePair
		}
	}

	switch jokers {
	case 4, 5:
		return FiveOfAKind
	case 3:
		return FourOfAKind
	case 2:
		return ThreeOfAKind
	case 1:
		return OnePair
	default:
		return HighCard
	}
}
