package main

import (
	helper "aoc/common"
	"fmt"
	"strings"
)

func main() {
	current_path := "2023/daytwo"

	fmt.Println(partTwoResults(helper.GetExampleFile(current_path)))
	// => 2286

	fmt.Println(partTwoResults(helper.GetPuzzleFile(current_path)))
	// => 69629
}

func partTwoResults(fileContent string) int {
	list := strings.Split(fileContent, "\n")
	total := 0

	for _, item := range list {
		cubes := map[string]int{"blue": 0, "red": 0, "green": 0}
		game := strings.Split(item, ":")

		if len(game) == 0 || !strings.Contains(game[0], "Game") {
			continue
		}

		for _, round := range strings.Split(game[1], ";") {
			for _, hand := range strings.Split(round, ",") {
				amount := GetIntFromString(hand)
				color := GetColorFromString(hand)

				if amount > cubes[color] {
					cubes[color] = amount
				}
			}
		}

		total += cubes["blue"] * cubes["red"] * cubes["green"]
	}

	return total
}
