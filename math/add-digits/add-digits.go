package main

import "strconv"

func addDigits(num int) int {
	if num < 10 {
		return num
	}

	numStr := strconv.Itoa(num)
	current := 0

	for _, v := range numStr {
		numInt, _ := strconv.Atoi(string(v))

		current += numInt
	}

	if current < 10 {
		return current
	}

	return addDigits(current)
}

func main() {
	num := 38

	result := addDigits(num)
	println(result)
}
