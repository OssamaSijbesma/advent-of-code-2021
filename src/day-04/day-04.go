package main

import (
	"advent-of-code-2021/utils/converter"
	"advent-of-code-2021/utils/files"
	"fmt"
	"strings"
)

func main() {
	input := files.ReadFile("\n")
	numberPool := converter.StringToIntSlice(input[0], ",")
	bingoCards := chunkBingoCards(input[1:])

	playBingo(numberPool, bingoCards)
}

func chunkBingoCards(slice []string) [][][]int {
	chunks := make([][][]int, (len(slice) / 6))

	for i, j := 0, -1; i < len(slice); i++ {
		if i%6 == 0 {
			j++
		} else {
			inputString := strings.ReplaceAll(strings.TrimSpace(slice[i]), "  ", " ")
			values := converter.StringToIntSlice(inputString, " ")
			chunks[j] = append(chunks[j], values)
		}
	}

	return chunks
}

func playBingo(numberPool []int, bingoCards [][][]int) map[int]int {
	winningCards := make(map[int]int)

	for _, drawnNumber := range numberPool {
		for c, card := range bingoCards {
			for r, row := range card {
				for i, value := range row {
					if value == drawnNumber {
						bingoCards[c][r][i] = -1
						if sumRow(bingoCards[c][r]) == -5 || sumColumn(bingoCards[c], i) == -5 {
							_, ok := winningCards[c]
							if !ok {
								winningCards[c] = (sumMatrix(bingoCards[c])) * value

								if len(winningCards) == 1 {
									fmt.Println("ex1. THE FIRST WINNER WITH A SCORE OF", winningCards[c])
								}

								if len(winningCards) == len(bingoCards) {
									fmt.Println("ex2. THE LAST LOSER? WITH A SCORE OF", winningCards[c])
									return winningCards
								}
							}
						}
					}
				}
			}
		}
	}

	return winningCards
}

func sumRow(row []int) int {
	sum := 0

	for _, v := range row {
		sum = sum + v
	}

	return sum
}

func sumColumn(matrix [][]int, pos int) int {
	sum := 0

	for i := 0; i < len(matrix); i++ {
		sum = sum + matrix[i][pos]
	}

	return sum
}

func sumMatrix(matrix [][]int) int {
	sum := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] != -1 {
				sum = sum + matrix[i][j]
			}
		}
	}

	return sum
}
