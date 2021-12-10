package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"sort"
)

func main() {
	input := files.ReadFile("\n")
	syntaxErrorScore, autoCompleteScore := checkSyntax(input)
	fmt.Println("ex1. Syntax error score:", syntaxErrorScore)
	fmt.Println("ex2. Autocomplete middle score:", autoCompleteScore)
}

func checkSyntax(data []string) (syntaxErrorScore int, autocompleteScore int) {
	scoreMap := map[rune]int{'(': 1, ')': 3, '[': 2, ']': 57, '{': 3, '}': 1197, '<': 4, '>': 25137}
	pairMap := map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}
	autocompleteScores := []int{}

	for _, line := range data {
		stack := []rune{}
		corruped := false
		for _, v := range line {
			switch v {
			case ')', ']', '}', '>':
				char := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if pairMap[char] != v {
					syntaxErrorScore += scoreMap[v]
					corruped = true
					break
				}
			default:
				stack = append(stack, v)
			}
		}

		if !corruped {
			totalScore := 0
			for len(stack) > 0 {
				char := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				totalScore = totalScore*5 + scoreMap[char]
			}

			autocompleteScores = append(autocompleteScores, totalScore)
		}
	}

	sort.Ints(autocompleteScores)
	autocompleteScore = autocompleteScores[len(autocompleteScores)/2]

	return
}
