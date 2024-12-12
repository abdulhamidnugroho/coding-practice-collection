package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		digits[i]++

		if digits[i] < 10 {
			return digits
		}

		// set it to 0 and continue with the carry over
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}

func main() {
	digits := []int{4, 3, 2, 1}

	result := plusOne(digits)
	fmt.Println(result) // Expected output: [4 3 2 2]
}
