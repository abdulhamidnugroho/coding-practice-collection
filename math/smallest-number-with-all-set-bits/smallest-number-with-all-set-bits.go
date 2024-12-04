package main

import (
	"fmt"
)

// Function to find the smallest number >= n with only set bits in its binary representation
func smallestNumber(n int) int {
	for {
		if n&(n+1) == 0 {
			return n
		}
		n++
	}
}

func main() {
	// Test cases
	testCases := []int{5, 10, 3, 67, 93}

	for _, n := range testCases {
		result := smallestNumber(n)
		fmt.Printf("Input: %d, Output: %d\n", n, result)
	}
}
