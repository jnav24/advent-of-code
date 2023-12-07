package helper

import (
	"os"
	"regexp"
)

func RemoveSpaces(value string) string {
	r, _ := regexp.Compile("\t")
	return r.ReplaceAllString(value, "")
}

func getFile(filename string) string {
	dir, _ := os.Getwd()
	b, err := os.ReadFile(dir + "/" + filename)

	if err != nil {
		panic("Could not find the file " + filename)
	}

	return string(b)
}

func GetExampleFile(path string) string {
	return getFile(path + "/example.txt")
}

func GetPuzzleFile(path string) string {
	return getFile(path + "/puzzle.txt")
}
