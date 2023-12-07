package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetIntFromString(value string) int {
	r, err := regexp.Compile("(?i)[a-z]")

	if err != nil {
		fmt.Println(err)
		return 0
	}

	i, err := strconv.Atoi(strings.TrimSpace(r.ReplaceAllString(value, "")))

	if err != nil {
		fmt.Println(err)
		return 0
	}

	return i
}

func GetColorFromString(value string) string {
	r, _ := regexp.Compile("[0-9]")
	return strings.TrimSpace(r.ReplaceAllString(value, ""))
}
