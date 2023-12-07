package main

import (
	helper "aoc/common"
	"fmt"
	"strings"
)

func main() {
	current_path := "2023/daytwo"

	fmt.Println(partOneResults(helper.GetExampleFile(current_path)))
	// => 8

	fmt.Println(partOneResults(helper.GetPuzzleFile(current_path)))
	//=> 2632
}

func partOneResults(fileContent string) int {
	list := strings.Split(fileContent, "\n")
	total := 0

	for _, item := range list {
		isPossible := true
		cubes := map[string]int{"blue": 14, "red": 12, "green": 13}

		game := strings.Split(item, ":")

		if len(game) == 0 || !strings.Contains(game[0], "Game") {
			continue
		}

		id := GetIntFromString(game[0])

		for _, round := range strings.Split(game[1], ";") {
			for _, hand := range strings.Split(round, ",") {
				amount := GetIntFromString(hand)
				color := GetColorFromString(hand)
				if isPossible {
					isPossible = amount <= cubes[color]
				}
			}
		}

		if isPossible {
			total += id
		}
	}

	return total
}
