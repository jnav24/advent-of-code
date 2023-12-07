package main

import (
	helper "aoc/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	current_path := "2023/dayone"

	fmt.Println(getPartOneSolution(helper.GetExampleFile(current_path)))
	// => 142

	fmt.Println(getPartOneSolution(helper.GetPuzzleFile(current_path)))
	// => 55538
}

func getPartOneSolution(fileContent string) int {
	list := strings.Split(fileContent, "\n")
	total := 0

	for _, item := range list {
		num, _ := strconv.Atoi(GetNumberFromString(helper.RemoveSpaces(item)))
		total = total + num
	}

	return total
}
