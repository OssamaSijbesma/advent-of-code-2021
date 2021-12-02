package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile("\n")
	data := splitValuesIntoMap(input, " ")
	dept := data["down"] - data["up"]

	fmt.Println("ex1. horizontal x dept =", data["forward"]*dept)
	fmt.Println("ex2. horizontal x dept =", data["horizontal"]*data["dept"])
}

func splitValuesIntoMap(stringSlice []string, delimiter string) map[string]int {
	mapResult := make(map[string]int)
	for _, value := range stringSlice {
		parts := strings.Split(value, delimiter)
		parsedValue, error := strconv.Atoi(parts[1])

		if error != nil {
			panic(error)
		}

		mapResult[parts[0]] += parsedValue

		if parts[0] == "forward" {
			aim := mapResult["down"] - mapResult["up"]
			mapResult["horizontal"] += parsedValue
			mapResult["dept"] += aim * parsedValue
		}
	}
	return mapResult
}
