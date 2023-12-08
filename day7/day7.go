package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kaveenr/aoc23/commons"
	"golang.org/x/exp/maps"
)

func Part1(input string) (result int) {
	games := parseInput(input)
	games.Sort(false)
	return games.Score()
}

func Part2(input string) (result int) {
	games := parseInput(input)
	games.Sort(true)
	return games.Score()
}

func parseInput(input string) (res Games) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(strings.TrimSpace(line), " ")
		if len(parts) != 2 {
			continue
		}
		bid, _ := strconv.Atoi(parts[1])
		res = append(res, Game{
			Hands: []rune(parts[0]),
			Bid:   bid,
		})
	}
	return res
}

type Games []Game

func (g Games) Sort(joker bool) {
	var lookup map[rune]int
	if joker {
		lookup = cardValuesAlt
	} else {
		lookup = cardValues
	}
	slices.SortFunc(g, func(j, k Game) int {
		j_r, k_r := j.Rank(joker), k.Rank(joker)
		if j_r == k_r {
			for idx := 0; idx < 5; idx++ {
				if j.Hands[idx] == k.Hands[idx] {
					continue
				}
				if lookup[j.Hands[idx]] > lookup[k.Hands[idx]] {
					return +1
				} else {
					return -1
				}
			}
		} else if j_r > k_r {
			return +1
		} else {
			return -1
		}
		return 0
	})
}

func (g Games) Score() (result int) {
	for rank, game := range g {
		result += game.Bid * (rank + 1)
	}
	return
}

type Game struct {
	Hands []rune
	Bid   int
}

func (g Game) Rank(joker bool) (rank int) {
	counts := make(map[rune]int)
	for _, card := range g.Hands {
		counts[card] += 1
	}
	values := maps.Values(counts)
	slices.Sort(values)
	// Five of a kind
	if slices.Equal(values, []int{5}) {
		rank = 7
		// Four of a kind
	} else if slices.Equal(values, []int{1, 4}) {
		rank = 6
		if joker && counts['J'] >= 1 {
			rank = 7
		}
		// Full house
	} else if slices.Equal(values, []int{2, 3}) {
		rank = 5
		if joker && counts['J'] >= 1 {
			rank = 7
		}
		// Three of a kind
	} else if slices.Equal(values, []int{1, 1, 3}) {
		rank = 4
		if joker && counts['J'] >= 1 {
			rank = 6
		}
		// Two pair
	} else if slices.Equal(values, []int{1, 2, 2}) {
		rank = 3
		if joker && counts['J'] == 1 {
			rank = 5
		} else if joker && counts['J'] == 2 {
			rank = 6
		}
		// One pair
	} else if slices.Equal(values, []int{1, 1, 1, 2}) {
		rank = 2
		if joker && counts['J'] >= 1 {
			rank = 4
		}
		// High Card
	} else if slices.Equal(values, []int{1, 1, 1, 1, 1}) {
		rank = 1
		if joker && counts['J'] == 1 {
			rank = 2
		}
	}
	return rank
}

var (
	cardValues = map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	cardValuesAlt = map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}
)

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
