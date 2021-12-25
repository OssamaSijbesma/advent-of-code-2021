package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"image"
	"regexp"
	"strconv"
	"strings"
)

type Map [][]int

func main() {
	input := files.ReadFile("\n")
	dots, instructions := decodeInput(input)
	var transparent Map
	transparent = createMap(dots, instructions)
	transparent = foldMap(transparent, instructions[0])
	fmt.Println("ex1. Amount of dots after the first fold:", countDots(transparent))

	for _, instruction := range instructions {
		transparent = foldMap(transparent, instruction)
	}

	fmt.Println("ex2. Infrared thermal imaging camera activation code:")

	for _, row := range transparent {
		for _, item := range row {
			if item > 0 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
	}
}

func decodeInput(data []string) (dots []image.Point, instructions []image.Point) {
	re := regexp.MustCompile("[0-9]+")

	for _, line := range data {
		if len(line) > 0 {
			if line[0] >= '0' && line[0] <= '9' {
				dotPos := strings.Split(line, ",")
				x, _ := strconv.Atoi(dotPos[0])
				y, _ := strconv.Atoi(dotPos[1])
				dots = append(dots, image.Point{x, y})
			} else if line[11] == 'x' {
				x, _ := strconv.Atoi(re.FindAllString(line, 1)[0])
				instructions = append(instructions, image.Point{x, 0})
			} else if line[11] == 'y' {
				y, _ := strconv.Atoi(re.FindAllString(line, 1)[0])
				instructions = append(instructions, image.Point{0, y})
			}
		}
	}

	return
}

func createMap(dots []image.Point, instructions []image.Point) (transparent Map) {
	xLength := (instructions[0].X*2 + 1)
	yLength := (instructions[1].Y*2 + 1)

	transparent = make(Map, yLength)
	for i := range transparent {
		transparent[i] = make([]int, xLength)
	}

	for _, dot := range dots {
		transparent[dot.Y][dot.X] = 1
	}

	return transparent
}

func foldMap(initialMap Map, instruction image.Point) (foldedMap Map) {
	var xLength, yLength int
	var xFold bool

	if instruction.X != 0 {
		xLength = instruction.X
		yLength = len(initialMap)
		xFold = true
	} else {
		xLength = len(initialMap[0])
		yLength = instruction.Y
		xFold = false
	}

	foldedMap = make(Map, yLength)
	for row := range foldedMap {
		foldedMap[row] = make([]int, xLength)
		for i := range foldedMap[row] {
			if xFold {
				foldedMap[row][i] = initialMap[row][i] + initialMap[row][len(initialMap[i])-i-1]
			} else {
				foldedMap[row][i] = initialMap[row][i] + initialMap[len(initialMap)-row-1][i]
			}
		}
	}

	return
}

func countDots(transparant Map) (count int) {
	for _, v := range transparant {
		for _, item := range v {
			if item > 0 {
				count++
			}
		}
	}
	return
}
