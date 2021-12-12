package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := files.ReadFile("\n")
	exits := setCaves(input)
	fmt.Println("ex1. Paths through this cave system that visit small caves at most once:", allPossiblePathsVisitOnce(exits, "start", map[string]int{}))
	fmt.Println("ex2. Amount of paths through this cave system:", allPossiblePathsOneDoubleVisit(exits, "start", map[string]int{}, false))
}

func setCaves(data []string) map[string][]string {
	exits := map[string][]string{}
	for _, line := range data {
		caves := strings.Split(line, "-")
		exits[caves[0]] = append(exits[caves[0]], caves[1])
		exits[caves[1]] = append(exits[caves[1]], caves[0])
	}
	return exits
}

func allPossiblePathsVisitOnce(exits map[string][]string, currentCave string, visits map[string]int) (sum int) {
	visits[currentCave] += 1
	for _, cave := range exits[currentCave] {
		if cave == "end" {
			sum++
			continue
		}
		if !unicode.IsLower(rune(cave[0])) || visits[cave] == 0 {
			nextVisits := visits
			if unicode.IsLower(rune(cave[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsVisitOnce(exits, cave, nextVisits)
		}
	}
	return
}

func allPossiblePathsOneDoubleVisit(exits map[string][]string, currentCave string, visits map[string]int, hasVisitedTwice bool) (sum int) {
	if visits[currentCave] > 0 && unicode.IsLower(rune(currentCave[0])) {
		hasVisitedTwice = true
	}
	visits[currentCave] += 1
	for _, cave := range exits[currentCave] {
		if cave == "end" {
			sum++
			continue
		}
		if cave != "start" && (!unicode.IsLower(rune(cave[0])) || visits[cave] < 1 || !hasVisitedTwice) {
			nextVisits := visits
			if unicode.IsLower(rune(cave[0])) {
				nextVisits = copyMap(visits)
			}
			sum += allPossiblePathsOneDoubleVisit(exits, cave, nextVisits, hasVisitedTwice)
		}
	}
	return
}

func copyMap(visits map[string]int) map[string]int {
	res := make(map[string]int, len(visits))
	for k, v := range visits {
		res[k] = v
	}
	return res
}
