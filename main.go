package main

func Solution(A []int) int {
	seen := make([]bool, len(A)+2)

	for _, value := range A {
		if value > 0 && value < len(seen) {
			seen[value] = true
		}
	}

	for candidate := 1; candidate < len(seen); candidate++ {
		if !seen[candidate] {
			return candidate
		}
	}

	return len(seen)
}
