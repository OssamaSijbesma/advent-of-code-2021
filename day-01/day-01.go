package main

import (
	"advent-of-code-2021/utils/converter"
	"advent-of-code-2021/utils/files"
	"fmt"
)

func main() {
	input := files.ReadFile("\n")
	values := converter.ToIntSlice(input)
	threeValues := toThreeMeasurement(values)
	fmt.Println(measurementAnalyzer(values))
	fmt.Println(measurementAnalyzer(threeValues))
}

func toThreeMeasurement(measurements []int) []int {
	var intSlice []int

	for i := 0; i < len(measurements)-2; i++ {
		intSlice = append(intSlice, (measurements[i] + measurements[i+1] + measurements[i+2]))
	}

	return intSlice
}

func measurementAnalyzer(measurements []int) int {
	prev, count := measurements[0], 0

	for _, measurement := range measurements[1:] {
		if measurement > prev {
			count++
		}

		prev = measurement
	}

	return count
}
