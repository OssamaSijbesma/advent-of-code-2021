package main

import (
	"advent-of-code-2021/utils/calc"
	"advent-of-code-2021/utils/converter"
	"advent-of-code-2021/utils/files"
	"fmt"
)

func main() {
	input := files.ReadFile(",")
	crabs := converter.ToIntSlice(input)

	fmt.Println("ex1. Least amount of fuel cost for the 🦀🦀🦀:", crabsAlignCost(crabs, false))
	fmt.Println("ex2. Least amount of fuel cost for the 🦀🦀🦀:", crabsAlignCost(crabs, true))
}

func crabsAlignCost(crabs []int, divergentSeries bool) (leastCost int) {
	min, max := calc.Min(crabs), calc.Max(crabs)

	for i := min; i <= max; i++ {
		cost := 0

		for _, v := range crabs {
			n := calc.AbsInt(i - v)
			if divergentSeries {
				cost += (n * (n + 1) / 2)
			} else {
				cost += n
			}
		}

		if i == min || cost < leastCost {
			leastCost = cost
		}
	}

	return
}
