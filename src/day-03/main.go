package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"strconv"
)

func main() {
	input := files.ReadFile("\n")
	gammaRate, epsilonRate := calcPowerConsumption(input)
	oxygenRating, co2Rating := calcLifeSupport(input)

	fmt.Println("ex1. power consumption", gammaRate*epsilonRate)
	fmt.Println("ex2. life support rating", oxygenRating*co2Rating)
}

func calcPowerConsumption(data []string) (int, int) {
	gamma := ""
	epsilon := ""

	for i := 0; i < len(data[0]); i++ {
		count := 0

		for _, v := range data {
			if v[i] == '1' {
				count++
			}
		}

		if count > len(data)/2 {
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		} else {
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		}
	}

	gammaRate, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilon, 2, 64)

	return int(gammaRate), int(epsilonRate)
}

func calcLifeSupport(data []string) (int, int) {
	baseChunks := chunkSlice(data, 0)
	oxygen := maxChunk(baseChunks, 0)
	co2 := minChunk(baseChunks, 0)

	for i := 1; i < len(data[0]); i++ {
		oxygen = maxChunk(chunkSlice(oxygen, i), i)
		co2 = minChunk(chunkSlice(co2, i), i)
	}

	oxygenRating, _ := strconv.ParseInt(oxygen[0], 2, 64)
	co2Rating, _ := strconv.ParseInt(co2[0], 2, 64)

	return int(oxygenRating), int(co2Rating)
}

func chunkSlice(slice []string, position int) [2][]string {
	var chunks [2][]string

	for _, v := range slice {
		if v[position] == '0' {
			chunks[0] = append(chunks[0], v)
		} else {
			chunks[1] = append(chunks[1], v)
		}
	}

	return chunks
}

func maxChunk(chunks [2][]string, pos int) []string {
	if len(chunks[0]) == len(chunks[1]) {
		if chunks[0][0][pos] == '1' {
			return chunks[0]
		} else {
			return chunks[1]
		}
	} else if len(chunks[0]) > len(chunks[1]) {
		return chunks[0]
	}
	return chunks[1]
}

func minChunk(chunks [2][]string, pos int) []string {

	if len(chunks[0]) == 0 {
		return chunks[1]
	} else if len(chunks[1]) == 0 {
		return chunks[0]
	}

	if len(chunks[0]) == len(chunks[1]) {
		if chunks[0][0][pos] == '0' {
			return chunks[0]
		} else {
			return chunks[1]
		}
	} else if len(chunks[0]) < len(chunks[1]) {
		return chunks[0]
	}
	return chunks[1]
}
