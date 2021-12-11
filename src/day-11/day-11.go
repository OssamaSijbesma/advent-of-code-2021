package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"image"
)

type Map [][]int

var offsets = [8]image.Point{
	{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1},
}

func main() {
	input := files.ReadFile("\n")
	fmt.Println("ex1. Flashes after 💯 steps:", dumboSim(toDumboMap(input), 100))
	fmt.Println("ex2. Steps until all 🐙 in sync:", dumboSimUntil(toDumboMap(input)))
}

func dumboSim(dumboMap Map, steps int) (flashes int) {
	for i := 0; i < steps; i++ {
		for x, row := range dumboMap {
			for y, _ := range row {
				dumboMap[x][y]++
				flashes += flashDumbo(&dumboMap, image.Pt(x, y))
			}
		}

		for x, row := range dumboMap {
			for y, _ := range row {
				if dumboMap[x][y] > 9 {
					dumboMap[x][y] = 0
				}
			}
		}
	}

	return
}

func dumboSimUntil(dumboMap Map) (step int) {
	amount := 0
	for amount < len(dumboMap)*len(dumboMap[0]) {
		amount = 0
		for x, row := range dumboMap {
			for y, _ := range row {
				dumboMap[x][y]++
				flashDumbo(&dumboMap, image.Pt(x, y))
			}
		}

		for x, row := range dumboMap {
			for y, _ := range row {
				if dumboMap[x][y] > 9 {
					dumboMap[x][y] = 0
					amount++
				}
			}
		}

		step++
	}

	return
}

func flashDumbo(dumboMap *Map, point image.Point) (flashes int) {
	if (*dumboMap)[point.X][point.Y] != 10 {
		return
	} else {
		for _, neigbour := range neighbours(*dumboMap, point) {
			(*dumboMap)[neigbour.X][neigbour.Y]++
			flashes += flashDumbo(dumboMap, neigbour)
		}
		flashes++
	}

	return
}

func toDumboMap(data []string) (dumboMap Map) {
	for _, line := range data {
		points := []int{}
		for _, value := range line {
			// Coverts rune to int
			points = append(points, (int(value) - '0'))
		}

		dumboMap = append(dumboMap, points)
	}

	return
}

func neighbours(m Map, point image.Point) (neighbours []image.Point) {
	for _, n_ := range offsets {
		n := point.Add(n_)
		if validPoint(m, n) {
			neighbours = append(neighbours, n)
		}
	}
	return
}

func validPoint(m Map, point image.Point) bool {
	goodX := 0 <= point.X && point.X < len(m)
	goodY := 0 <= point.Y && point.Y < len(m[0])
	return goodX && goodY
}
