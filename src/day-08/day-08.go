package main

import (
	"advent-of-code-2021/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile("\n")
	splitInput := splitContent(input)

	fmt.Println("ex1. The digits 1, 4, 7 or 8 appear:", checkDigitOccurance(splitInput), "times")
	fmt.Println("ex2. The sum of all output values is:", sumNumbers(splitInput))
}

func sumNumbers(data [][][]string) (sum int) {
	for _, lines := range data {
		numbers, fLN, sLN := make([]string, 10), []string{}, []string{}

		for _, value := range lines[0] {
			switch len(value) {
			case 2:
				numbers[1] = value
			case 3:
				numbers[7] = value
			case 4:
				numbers[4] = value
			case 7:
				numbers[8] = value
			case 5:
				fLN = append(fLN, value)
			case 6:
				sLN = append(sLN, value)
			}
		}

		for _, v := range fLN {
			if matchCount(v, numbers[1]) == 2 {
				numbers[3] = v
			} else if matchCount(v, numbers[4]) == 3 {
				numbers[5] = v
			} else {
				numbers[2] = v
			}
		}

		for _, v := range sLN {
			if matchCount(v, numbers[4]) == 4 {
				numbers[9] = v
			} else if matchCount(v, numbers[4]) == 3 && matchCount(v, numbers[1]) == 1 {
				numbers[6] = v
			} else {
				numbers[0] = v
			}
		}

		strNum := ""

		for _, value := range lines[1] {
			for number, matcher := range numbers {
				if len(value) == len(matcher) && matchCount(value, matcher) == len(matcher) {
					strNum = strNum + strconv.Itoa(number)
				}
			}
		}

		num, _ := strconv.Atoi(strNum)
		sum += num
	}

	return
}

func matchCount(original string, matcher string) (count int) {
	for _, v1 := range original {
		for _, v2 := range matcher {
			if v1 == v2 {
				count++
			}
		}
	}

	return
}

func checkDigitOccurance(data [][][]string) (count int) {
	for _, lines := range data {
		for _, value := range lines[1] {
			switch len(value) {
			case
				2, 3, 4, 7:
				count++
			}
		}
	}

	return
}

func splitContent(data []string) (splittedContent [][][]string) {
	for _, line := range data {
		values := strings.Split(line, "|")
		splittedContent = append(splittedContent, [][]string{
			strings.Split(strings.TrimSpace(values[0]), " "),
			strings.Split(strings.TrimSpace(values[1]), " "),
		})
	}

	return
}
