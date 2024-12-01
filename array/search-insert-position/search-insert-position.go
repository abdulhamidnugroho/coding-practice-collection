package main

import (
	"fmt"
)

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// If not found, `left` will be the insertion position
	return left
}

func main() {
	// Test cases
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5)) // Output: 2
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2)) // Output: 1
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7)) // Output: 4
}
