package main

func biggest(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func max(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	// for i := 0; i <= len(numbers)-1; i++ {
	// 	if numbers[i] > max {
	// 		max = numbers[i]
	// 	}
	// }
	return max
}

func exSwitch(day string) {
	switch day {
	case "sunday":
		// cases don't "fall through" by default!
		fallthrough

	case "saturday":
		// rest()

	default:
		// work()
	}
}
