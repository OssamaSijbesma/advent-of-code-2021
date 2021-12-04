package converter

import (
	"strconv"
	"strings"
)

func ToIntSlice(data []string) []int {
	var intSlice []int

	for _, value := range data {
		// Convert string to int
		result, error := strconv.Atoi(value)

		if error != nil {
			panic(error)
		}

		intSlice = append(intSlice, result)
	}

	return intSlice
}

func StringToIntSlice(data string, delimiter string) []int {
	var intSlice []int

	slicedContent := strings.Split(data, delimiter)

	for _, value := range slicedContent {

		// Convert string to int
		result, error := strconv.Atoi(value)

		if error != nil {
			panic(error)
		}

		intSlice = append(intSlice, result)
	}

	return intSlice
}
