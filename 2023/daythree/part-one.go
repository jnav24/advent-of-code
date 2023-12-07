package main

import (
	helper "aoc/common"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	current_path := "2023/daythree"

	fmt.Println(getPartOneResults(helper.GetExampleFile(current_path)))
	// => 4361

	fmt.Println(getPartOneResults(helper.GetPuzzleFile(current_path)))
	// => 539590
}

func getPartOneResults(fileContent string) int {
	rows := strings.Split(fileContent, "\n")
	total := 0

	for rowIndex, row := range rows {
		r := regexp.MustCompile("(\\d+)")

		if !r.MatchString(row) {
			continue
		}

		for _, indexSlice := range r.FindAllStringIndex(row, -1) {
			if isAdjacentToSymbol(rows, rowIndex, indexSlice[0], indexSlice[0]+(indexSlice[1]-indexSlice[0])-1) {
				number := ""

				for i := indexSlice[0]; i < indexSlice[1]; i++ {
					number += string(rows[rowIndex][i])
				}

				num, _ := strconv.Atoi(number)
				total += num
			}
		}

	}

	return total
}

func isAdjacentToSymbol(rows []string, rowIndex int, offset int, offsetLength int) bool {
	steps := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for col := offset; col <= offsetLength; col++ {
		for _, step := range steps {
			yaxis := rowIndex + step[0]
			xaxis := col + step[1]

			if yaxis >= 0 && yaxis < len(rows) && xaxis >= 0 && xaxis < len(rows[yaxis]) {
				char := string(rows[yaxis][xaxis])

				if !isNumber(char) && char != "." {
					return true
				}
			}
		}
	}

	return false
}

func isNumber(value string) bool {
	r, _ := regexp.Compile("^[0-9]+$")
	return r.MatchString(value)
}
