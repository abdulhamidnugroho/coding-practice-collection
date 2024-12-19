package main

import "fmt"

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	nums2 := []int{2, 2, 1}

	fmt.Println(singleNumber(nums))  // output: 4
	fmt.Println(singleNumber(nums2)) // output: 1
}
