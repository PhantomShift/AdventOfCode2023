package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

const DaySevenTestInput = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`

var CardStrengths map[rune]int = map[rune]int{'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12}

var CardStrengthsRevised map[rune]int = map[rune]int{'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'Q': 10, 'K': 11, 'A': 12}

type camelHand struct {
	strengths []int
	bid       int
	category  int
}

func getCategory(hand string, joker bool) int {
	cards := map[rune]int{}
	if joker {
		max, maxCard := 0, 'J'
		for _, char := range hand {
			cards[char]++
			if char == 'J' {
				continue
			}
			if cards[char] > max || (cards[char] == max && CardStrengths[char] >= CardStrengths[maxCard]) {
				max = cards[char]
				maxCard = char
			}
		}
		if maxCard != 'J' {
			cards[maxCard] += cards['J']
			delete(cards, 'J')
		}
	} else {
		for _, char := range hand {
			cards[char]++
		}
	}
	counts := map[int]int{}
	for _, count := range cards {
		counts[count]++
	}
	if counts[5] == 1 {
		return 6
	} else if counts[4] == 1 {
		return 5
	} else if counts[3] == 1 && counts[2] == 1 {
		return 4
	} else if counts[3] == 1 && counts[1] == 2 {
		return 3
	} else if counts[2] == 2 {
		return 2
	} else if counts[2] == 1 {
		return 1
	}
	return 0
}

func compareHands(a camelHand, b camelHand) int {
	cmp := a.category - b.category
	if cmp != 0 {
		return cmp
	}

	for i, strengthA := range a.strengths {
		strengthB := b.strengths[i]
		res := strengthA - strengthB
		if res != 0 {
			return res
		}
	}

	return 0
}

func day7(part int, testing bool) {
	input := DaySevenTestInput
	if !testing {
		input = readInput(7, false)
	}
	joker := part == 2
	strengthRule := CardStrengths
	if joker {
		strengthRule = CardStrengthsRevised
	}

	hands := make([]camelHand, 0)
	for _, line := range strings.Split(input, "\n") {
		if len(line) == 0 {
			continue
		}
		hand, bidString, _ := strings.Cut(line, " ")
		strengths := make([]int, len(hand))
		for i, char := range hand {
			strengths[i] = strengthRule[char]
		}
		bid, _ := strconv.Atoi(bidString)
		hands = append(hands, camelHand{
			strengths: strengths,
			bid:       bid,
			category:  getCategory(hand, joker),
		})
	}
	slices.SortFunc(hands, compareHands)
	sum := 0
	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	fmt.Printf("The sum is %d\n", sum)
}
