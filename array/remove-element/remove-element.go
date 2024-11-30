package main

import (
	"fmt"
)

// should modify the slice as well
func removeElement(nums []int, val int) int {
	k := 0

	for _, num := range nums {
		// If the current number is not equal to val, keep it
		if num != val {
			nums[k] = num
			k++

			fmt.Println(nums)
		}
	}

	// Return the count of elements not equal to val
	return k
}

func main() {
	// Example
	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	k2 := removeElement(nums2, val2)
	fmt.Printf("Output: %d, nums = %v\n", k2, nums2[:k2])
}
