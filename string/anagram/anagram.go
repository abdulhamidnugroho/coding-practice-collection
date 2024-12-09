package main

import "fmt"

// makeAnagram determines the minimum number of deletions required to make two strings anagrams.
func makeAnagram(a, b string) int {
	freq := make([]int, 26)

	// Count frequency differences
	for _, char := range a {
		freq[char-'a']++
	}
	for _, char := range b {
		freq[char-'a']--
	}

	// Sum the absolute differences
	deletions := 0
	for _, count := range freq {
		if count < 0 {
			deletions -= count
		} else {
			deletions += count
		}
	}

	return deletions
}

func main() {
	// Example input
	a := "cde"
	b := "dcf"

	result := makeAnagram(a, b)

	// Output the result
	fmt.Printf("minimum deletions to make anagrams: %d\n", result)
}
