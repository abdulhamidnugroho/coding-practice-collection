package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Two pointers, i and k, to iterate through the array and keep track of unique elements
	k := 0

	for i := 1; i < len(nums); i++ {
		// If a new unique element is found
		if nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	// Return the count of unique elements
	return k + 1
}

func main() {
	// Example 1
	nums1 := []int{1, 1, 2}
	k1 := removeDuplicates(nums1)
	fmt.Printf("Output: %d, nums = %v\n", k1, nums1[:k1])

	// Example 2
	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k2 := removeDuplicates(nums2)
	fmt.Printf("Output: %d, nums = %v\n", k2, nums2[:k2])
}
