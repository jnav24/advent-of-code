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

	fmt.Println(getPartTwoResults(helper.GetExampleFile(current_path)))
	// => 467835

	fmt.Println(getPartTwoResults(helper.GetPuzzleFile(current_path)))
	// => 80703636
}

func getPartTwoResults(fileContent string) int {
	rows := strings.Split(fileContent, "\n")
	total := 0
	gears := map[int]map[int][]int{}

	for rowIndex, row := range rows {
		r := regexp.MustCompile("(\\d+)")

		if !r.MatchString(row) {
			continue
		}

		for _, indexSlice := range r.FindAllStringIndex(row, -1) {
			if isAdjacent, position := isAdjacentToGear(rows, rowIndex, indexSlice[0], indexSlice[0]+(indexSlice[1]-indexSlice[0])-1); isAdjacent {
				number := ""

				for i := indexSlice[0]; i < indexSlice[1]; i++ {
					number += string(rows[rowIndex][i])
				}

				num, _ := strconv.Atoi(number)

				if _, ok := gears[position[0]]; !ok {
					gears[position[0]] = map[int][]int{}
				}

				gears[position[0]][position[1]] = append(gears[position[0]][position[1]], num)
			}
		}

	}

	getTotal(&total, gears)

	return total
}

func getTotal(total *int, gears map[int]map[int][]int) {
	for _, y := range gears {
		for _, x := range y {
			if len(x) == 2 {
				*total += x[0] * x[1]
			}
		}
	}
}

func isAdjacentToGear(rows []string, rowIndex int, offset int, offsetLength int) (bool, []int) {
	steps := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1},
	}

	for col := offset; col <= offsetLength; col++ {
		for _, step := range steps {
			yaxis := rowIndex + step[0]
			xaxis := col + step[1]

			if yaxis >= 0 && yaxis < len(rows) && xaxis >= 0 && xaxis < len(rows[yaxis]) {
				char := string(rows[yaxis][xaxis])

				if char == "*" {
					return true, []int{yaxis, xaxis}
				}
			}
		}
	}

	return false, make([]int, 0)
}
