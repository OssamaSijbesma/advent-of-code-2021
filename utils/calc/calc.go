package calc

func Max(data []int) (max int) {
	for i, v := range data {
		if i == 0 || v > max {
			max = v
		}
	}

	return
}

func Min(data []int) (min int) {
	for i, v := range data {
		if i == 0 || v < min {
			min = v
		}
	}

	return
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
