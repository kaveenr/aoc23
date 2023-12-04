package main

import (
	"container/list"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/kaveenr/aoc23/commons"
)

func Part1(input string) (result int) {
	puzzles := ParsePuzzle(input)
	for _, puz := range puzzles {
		var cardVal int
		for _, num := range puz.Numbers {
			if slices.Contains(puz.WiningNumbers, num) {
				if cardVal == 0 {
					cardVal = 1
					continue
				}
				cardVal *= 2
			}
		}
		result += cardVal
	}
	return result
}

func Part2(input string) (result int) {
	cards, stack := ParsePuzzle(input), list.New()
	for idx := 0; idx < len(cards); idx++ {
		stack.PushBack(cards[idx+1])
	}
	for stack.Len() > 0 {
		element := stack.Front()
		stack.Remove(element)
		card, acc := element.Value.(Card), element.Value.(Card).Number
		for _, num := range card.Numbers {
			if slices.Contains(card.WiningNumbers, num) {
				acc += 1
				if val, ok := cards[acc]; ok {
					stack.PushFront(val)
				}
			}
		}
		result += 1
	}
	return result
}

func ParsePuzzle(input string) (res PuzzleInput) {
	res = make(map[int]Card)
	for idx, line := range strings.Split(strings.TrimSpace(input), "\n") {
		cardnum := idx + 1
		card := Card{
			Number: cardnum,
		}
		mainParts := strings.Split(line, ":")
		numberParts := strings.Split(mainParts[1], "|")
		for _, num := range strings.Split(strings.TrimSpace(numberParts[0]), " ") {
			parsedNum, err := strconv.Atoi(num)
			if err == nil {
				card.WiningNumbers = append(card.WiningNumbers, parsedNum)
			}
		}
		for _, num := range strings.Split(strings.TrimSpace(numberParts[1]), " ") {
			parsedNum, err := strconv.Atoi(num)
			if err == nil {
				card.Numbers = append(card.Numbers, parsedNum)
			}
		}
		res[cardnum] = card
	}
	return res
}

type Card struct {
	Number        int
	WiningNumbers []int
	Numbers       []int
}

type PuzzleInput = map[int]Card

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
