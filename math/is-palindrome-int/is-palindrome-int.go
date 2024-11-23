package main

import "fmt"

// Not a mathematical way
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	s := fmt.Sprint(x)

	left, right := 0, len(s)-1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func main() {
	x := 12321
	fmt.Println(isPalindrome(x))
}
