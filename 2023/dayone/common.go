package main

import "regexp"

func GetNumberFromString(value string) string {
	r, _ := regexp.Compile("[a-z]")
	result := r.ReplaceAllString(value, "")
	resultLen := len(result)

	if resultLen == 0 {
		return "0"
	}

	if resultLen == 1 {
		return result + result
	}

	return result[0:1] + result[resultLen-1:]
}
