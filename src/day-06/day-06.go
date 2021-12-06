package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"strconv"
)

func main() {
	input := files.ReadFile(",")
	var laternfish []int
	for _, fish := range input {
		fishLife, _ := strconv.Atoi(fish)
		laternfish = append(laternfish, fishLife)
	}

	fmt.Println("ex1. Amount of 🐟 after 80 days:", lanternfishSim(laternfish, 80))
	fmt.Println("ex2. Amount of 🐟 after 256 days:", efficientLaternfishSim(laternfish, 256))
}

func lanternfishSim(initial []int, days int) int {
	var laternfish []int
	laternfish = append(laternfish, initial...)
	newFish := 0
	for i := 0; i < days; i++ {
		for i, v := range laternfish {
			if v == 0 {
				newFish++
				laternfish[i] = 6
			} else {
				laternfish[i]--
			}
		}

		for i := 0; i < newFish; i++ {
			laternfish = append(laternfish, 8)
		}
		newFish = 0
	}

	return len(laternfish)
}

func efficientLaternfishSim(initial []int, days int) int {
	laternfish := make([]int, 9)

	for _, v := range initial {
		laternfish[v]++
	}

	for i := 0; i < days; i++ {
		nextLaternfish := make([]int, 9)
		nextLaternfish[6] = 0 + laternfish[0]
		nextLaternfish[8] = 0 + laternfish[0]

		for i := 1; i < 9; i++ {
			nextLaternfish[i-1] = nextLaternfish[i-1] + laternfish[i]
		}

		laternfish = nextLaternfish
	}

	sum := 0
	for _, v := range laternfish {
		sum = sum + v
	}

	return sum
}
