package converter

import (
	"strconv"
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
