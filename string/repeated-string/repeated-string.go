package main

import "fmt"

func repeatedString(s string, n int64) int64 {
	var count int64
	length := int64(len(s))

	for _, c := range s {
		if c == 'a' {
			count++
		}
	}

	fullReps := n / length
	count *= fullReps

	for i := int64(0); i < n%length; i++ {
		if s[i] == 'a' {
			count++
		}
	}

	return count
}

func main() {
	// Define input
	s := "abcac"
	n := int64(10)

	// Calculate the result
	result := repeatedString(s, n)

	// Output the result
	fmt.Printf("String: %s\n", s)
	fmt.Printf("Length to consider: %d\n", n)
	fmt.Printf("Number of 'a's: %d\n", result)
}
