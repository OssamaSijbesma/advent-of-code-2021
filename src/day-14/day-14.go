package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := files.ReadFile("\n")
	polymerTemplate, pairInsertion := decodeInput(input)

	fmt.Println("ex1. Polymerization quantity after 10 steps:", bruteForce(polymerTemplate, pairInsertion))
	fmt.Println("ex2. Polymerization quantity after 40 steps:", memoization(polymerTemplate, pairInsertion))
}

func decodeInput(data []string) (polymerTemplate []string, pairInsertion map[string]string) {
	pairInsertion = make(map[string]string)
	polymerTemplate = strings.Split(data[0], "")

	for _, line := range data[2:] {
		elements := strings.Split(line, " -> ")
		pairInsertion[elements[0]] = elements[1]
	}

	return
}

func polymerGrowStep(polymerTemplate []string, pairInsertion map[string]string) (newPolymerTemplate []string) {

	for i := range polymerTemplate[:len(polymerTemplate)-1] {
		key := polymerTemplate[i] + polymerTemplate[i+1]
		newPolymerTemplate = append(newPolymerTemplate, polymerTemplate[i], pairInsertion[key])
	}
	newPolymerTemplate = append(newPolymerTemplate, polymerTemplate[len(polymerTemplate)-1])
	return
}

func memoizationPolymerGrowStep(mem map[string]map[string]int, polymerTemplate []string, pairInsertion map[string]string, step int) (elementCount map[string]int) {
	elementCount = map[string]int{}
	if step == 0 {
		return elementCount
	}

	key := fmt.Sprint(polymerTemplate, step)

	if result, ok := mem[key]; ok {
		return result
	}

	for i := range polymerTemplate[:len(polymerTemplate)-1] {
		pair := polymerTemplate[i] + polymerTemplate[i+1]
		elementCount[pairInsertion[pair]]++

		elementCountResult := memoizationPolymerGrowStep(mem, []string{polymerTemplate[i], pairInsertion[pair], polymerTemplate[i+1]}, pairInsertion, step-1)
		for k, count := range elementCountResult {
			elementCount[k] += count
		}
	}

	mem[key] = elementCount
	return
}

func countElements(polymerTemplate []string) (leastCommon int, mostCommon int) {
	elements := make(map[string]int)

	for _, v := range polymerTemplate {
		elements[v] = elements[v] + 1
	}

	leastCommon, mostCommon = elements[polymerTemplate[0]], elements[polymerTemplate[0]]

	for _, elementCount := range elements {
		if elementCount > mostCommon {
			mostCommon = elementCount
		} else if elementCount < leastCommon {
			leastCommon = elementCount
		}
	}

	return
}

func bruteForce(polymerTemplate []string, pairInsertion map[string]string) (total int) {
	for i := 0; i < 10; i++ {
		polymerTemplate = polymerGrowStep(polymerTemplate, pairInsertion)
	}
	leastCommon, mostCommon := countElements(polymerTemplate)
	return mostCommon - leastCommon
}

func memoization(polymerTemplate []string, pairInsertion map[string]string) (total int) {
	elementCount := memoizationPolymerGrowStep(map[string]map[string]int{}, polymerTemplate, pairInsertion, 40)

	for _, v := range polymerTemplate {
		elementCount[v]++
	}

	mostCommon, leastCommon := 0, math.MaxInt64
	for _, count := range elementCount {
		if count > mostCommon {
			mostCommon = count
		} else if count < leastCommon {
			leastCommon = count
		}
	}

	return mostCommon - leastCommon
}
