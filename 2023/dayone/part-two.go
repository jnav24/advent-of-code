package main

import (
	helper "aoc/common"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	current_path := "2023/dayone"

	fmt.Println(getPartTwoSolution(strings.Split(helper.GetExampleFile(current_path), "\n")))
	// => 281

	fmt.Println(getPartTwoSolution(strings.Split(helper.GetPuzzleFile(current_path), "\n")))
	// => 54875
}

func getPartTwoSolution(list []string) int {
	total := 0

	for _, item := range list {
		num, _ := strconv.Atoi(GetNumberFromString(replaceNumbersWithDigits(helper.RemoveSpaces(item))))
		total = total + num
	}

	return total
}

func replaceNumbersWithDigits(value string) string {
	// @NOTE
	// created two slices because if I were to loop through a map, Go will randomly shuffle the order, causing random results
	keys := []string{"98", "182", "82", "18", "21", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	values := []string{"nineight", "oneightwo", "eightwo", "oneight", "twone", "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	for k, v := range values {
		value = strings.Replace(value, v, keys[k], -1)
	}

	return value
}
