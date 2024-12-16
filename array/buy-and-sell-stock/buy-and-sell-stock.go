package main

import "fmt"

func buyStock(prices []int) int {
	smallest := prices[0]
	max := 0

	for _, v := range prices {
		if v < smallest {
			smallest = v
		}

		now := v - smallest
		if now > max {
			max = now
		}
	}

	return max
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(buyStock(prices)) // 5

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(buyStock(prices)) // 0
}
