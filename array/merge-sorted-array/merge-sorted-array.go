package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Pointers initialized at the end of the arrays
	point1, point2, p := m-1, n-1, m+n-1

	// Merge nums1 and nums2 starting from the end
	for point1 >= 0 && point2 >= 0 {
		if nums1[point1] > nums2[point2] {
			nums1[p] = nums1[point1]
			point1--
		} else {
			nums1[p] = nums2[point2]
			point2--
		}
		p--
	}

	// If there are remaining elements in nums2, copy them
	for point2 >= 0 {
		nums1[p] = nums2[point2]
		point2--
		p--
	}
}

func main() {
	// Test Case 1
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	fmt.Println("Before merge:")
	fmt.Println("nums1:", nums1)
	fmt.Println("nums2:", nums2)

	// Merge nums2 into nums1
	merge(nums1, m, nums2, n)

	fmt.Println("\nAfter merge:")
	fmt.Println("nums1:", nums1) // Expected output: [1, 2, 2, 3, 5, 6]

	// Test Case 2
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0

	fmt.Println("\nBefore merge:")
	fmt.Println("nums1:", nums1)
	fmt.Println("nums2:", nums2)

	// Merge nums2 into nums1
	merge(nums1, m, nums2, n)

	fmt.Println("\nAfter merge:")
	fmt.Println("nums1:", nums1) // Expected output: [1]
}
