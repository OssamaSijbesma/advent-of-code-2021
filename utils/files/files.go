package files

import (
	"io/ioutil"
	"strings"
)

func ReadFile(delimiter string) []string {
	file, err := ioutil.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, delimiter)

	if delimiter == "\n" {
		// fetch utils always adds a new line at the end of a file, which could lead to some problems when parsing it
		// this is why the last row is just removed if the delimiter is a newline
		return slicedContent[:len(slicedContent)-1]
	} else {
		// if the delimiter is not a newline and we split on eg. a comma, the newline will be appended to the last
		// element in the slice which then cannot be converted to an int.
		// this is the reason the last element in the slice is modified (the last char is removed
		// [which is the extra newline]) so it can be worked with
		lastElement := slicedContent[len(slicedContent)-1]
		slicedContent[len(slicedContent)-1] = lastElement[:len(lastElement)-1]

		return slicedContent
	}
}
