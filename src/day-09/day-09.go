package main

import (
	"advent-of-code-2021/utils/calc"
	"advent-of-code-2021/utils/files"
	"advent-of-code-2021/utils/vector"
	"fmt"
	"sort"
)

func main() {
	input := files.ReadFile("\n")
	heightmap := toHeightmap(input)
	lowPoints := getLowPoints(heightmap)
	riskLevels := calcRiskLevel(heightmap, lowPoints)
	fmt.Println("ex1. Sum risk levels:", calc.Sum(riskLevels))
	fmt.Println("ex2. Sum largest basin size:", largestBasins(heightmap, lowPoints))
}

func largestBasins(heightmap [][]int, lowPoints []vector.Vector) (value int) {
	sizes := []int{}

	for _, lowPoint := range lowPoints {
		sizes = append(sizes, getBasinSize(heightmap, lowPoint))
	}

	sort.Ints(sizes)

	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

func calcRiskLevel(heightmap [][]int, lowPoints []vector.Vector) (riskLevels []int) {
	for _, lowPoint := range lowPoints {
		x, y := int(lowPoint.X), int(lowPoint.Y)
		riskLevels = append(riskLevels, heightmap[x][y]+1)
	}

	return
}

func getLowPoints(heightmap [][]int) (lowPoints []vector.Vector) {
	for v, line := range heightmap {
		for h, value := range line {
			if (v == 0 || value < heightmap[v-1][h]) && (v == len(heightmap)-1 || value < heightmap[v+1][h]) &&
				(h == 0 || value < heightmap[v][h-1]) && (h == len(heightmap[0])-1 || value < heightmap[v][h+1]) {
				lowPoints = append(lowPoints, vector.New(v, h))
			}
		}
	}

	return
}

func getBasinSize(heightmap [][]int, root vector.Vector) (size int) {
	current := root
	seenPositions := []vector.Vector{current}
	queue := []vector.Vector{current}
	size++

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		for _, neighbour := range getValidNeighbours(heightmap, current) {
			valid := true

			for _, position := range seenPositions {
				if neighbour == position {
					valid = false
				}
			}

			if valid {
				seenPositions = append(seenPositions, neighbour)
				queue = append(queue, neighbour)
				size++
			}
		}
	}

	return size
}

func getValidNeighbours(heightmap [][]int, root vector.Vector) (neighbours []vector.Vector) {
	x, y := int(root.X), int(root.Y)
	value := heightmap[x][y]

	if value+1 == 9 {
		return
	}

	if x != 0 && value+1 == heightmap[x-1][y] {
		neighbours = append(neighbours, vector.New(x-1, y))
	}
	if x != len(heightmap)-1 && value+1 == heightmap[x+1][y] {
		neighbours = append(neighbours, vector.New(x+1, y))
	}
	if y != 0 && value+1 == heightmap[x][y-1] {
		neighbours = append(neighbours, vector.New(x, y-1))
	}
	if y != len(heightmap[0])-1 && value+1 == heightmap[x][y+1] {
		neighbours = append(neighbours, vector.New(x, y+1))
	}

	return
}

func toHeightmap(data []string) (heightmap [][]int) {
	for _, line := range data {
		points := []int{}
		for _, value := range line {
			// Coverts rune to int
			points = append(points, (int(value) - '0'))
		}

		heightmap = append(heightmap, points)
	}

	return
}
