package main

import (
	"advent-of-code-2021/utils/converter"
	"advent-of-code-2021/utils/files"
	"advent-of-code-2021/utils/vector"
	"fmt"
	"strings"
)

func main() {
	input := files.ReadFile("\n")
	hydrothermalVentsMap1 := scanHydrothermalVents(filterAngleDroit(toVectorChunks(input)))
	hydrothermalVentsMap2 := scanHydrothermalVents(toVectorChunks(input))
	fmt.Println("ex1. Angle droit number of dangerous hydrothermal vents:", countHydrothermalVents(hydrothermalVentsMap1))
	fmt.Println("ex2. Number of dangerous hydrothermal vents:", countHydrothermalVents(hydrothermalVentsMap2))
}

func toVectorChunks(data []string) (result [][]vector.Vector) {
	for _, line := range data {
		values := converter.StringToIntSlice(strings.ReplaceAll(line, " -> ", ","), ",")
		result = append(result, []vector.Vector{vector.New(values[0], values[1]), vector.New(values[2], values[3])})
	}

	return
}

func filterAngleDroit(data [][]vector.Vector) (result [][]vector.Vector) {
	for _, v := range data {
		if v[0].X == v[1].X || v[0].Y == v[1].Y {
			result = append(result, v)
		}
	}

	return
}

func scanHydrothermalVents(ventData [][]vector.Vector) map[vector.Vector]int {
	hydrothermalVentsMap := make(map[vector.Vector]int)

	for _, v := range ventData {
		hydrothermalVentsMap[v[0]]++

		for v[0].X != v[1].X || v[0].Y != v[1].Y {
			if v[0].X < v[1].X {
				v[0].X++
			} else if v[0].X > v[1].X {
				v[0].X--
			}

			if v[0].Y < v[1].Y {
				v[0].Y++
			} else if v[0].Y > v[1].Y {
				v[0].Y--
			}

			hydrothermalVentsMap[v[0]]++
		}
	}

	return hydrothermalVentsMap
}

func countHydrothermalVents(hydrothermalVentsMap map[vector.Vector]int) (count int) {
	for _, v := range hydrothermalVentsMap {
		if v > 1 {
			count++
		}
	}

	return
}
