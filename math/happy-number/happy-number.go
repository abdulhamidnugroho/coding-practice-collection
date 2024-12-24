package main

import "fmt"

func isHappy(n int) bool {
	exists := make(map[int]bool)

	for n != 1 {
		if exists[n] {
			return false
		}
		exists[n] = true
		n = perProcess(n)
	}

	return true
}

func perProcess(num int) int {
	sum := 0
	for num > 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19)) // true
	fmt.Println(isHappy(2))  // false
}
