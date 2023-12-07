package main

import (
	helper "aoc/common"
	"fmt"
	"regexp"
	"slices"
	"strings"
)

func main() {
	current_path := "2023/dayfour"

	fmt.Println(getPartOneSolution(helper.GetExampleFile(current_path)))
	// => 13

	fmt.Println(getPartOneSolution(helper.GetPuzzleFile(current_path)))
	// => 23028
}

func getPartOneSolution(fileContent string) int {
	items := strings.Split(helper.RemoveSpaces(fileContent), "\n")
	total := 0

	for _, item := range items {
		if !strings.Contains(item, ":") {
			continue
		}

		cardTotal := 0
		list := strings.Split(item, ":")

		r, _ := regexp.Compile("\\s+")
		trimmed := r.ReplaceAllString(list[1], " ")

		cards := strings.Split(trimmed, "|")
		winningCards := strings.Split(strings.TrimSpace(cards[0]), " ")
		myCards := strings.Split(cards[1], " ")

		for _, card := range myCards {
			if slices.Contains(winningCards, card) {
				if cardTotal == 0 {
					cardTotal += 1
					continue
				}

				cardTotal = cardTotal * 2
			}
		}

		total += cardTotal
	}

	return total
}
