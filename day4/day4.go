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
	for _, card := range cards {
		stack.PushBack(card)
	}
	for stack.Len() > 0 {
		element := stack.Front()
		stack.Remove(element)
		card, acc := element.Value.(Card), element.Value.(Card).Number
		for _, num := range card.Numbers {
			if slices.Contains(card.WiningNumbers, num) {
				acc++
				stack.PushFront(cards[acc-1])
			}
		}
		result += 1
	}
	return result
}

// Assumption: Input is given in order
func ParsePuzzle(input string) (res PuzzleInput) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	res = make([]Card, len(lines))
	for idx, line := range lines {
		mainParts := strings.Split(line, ":")
		card := Card{
			Number: idx + 1,
		}
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
		res[idx] = card
	}
	return res
}

type Card struct {
	Number        int
	WiningNumbers []int
	Numbers       []int
}

type PuzzleInput = []Card

func main() {
	myPuzzleInput, _ := commons.LoadFile(`input.txt`)
	fmt.Println("Part 1 Solution:", Part1(myPuzzleInput))
	fmt.Println("Part 2 Solution:", Part2(myPuzzleInput))
}
