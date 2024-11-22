package main

import "fmt"

func twoSum(nums []int, target int) []int {
	res := make(map[int]int)

	for i, num := range nums {
		sub := target - num
		if index, ok := res[sub]; ok {
			return []int{index, i}
		}

		res[num] = i
	}

	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // Output: [0 1]
}
