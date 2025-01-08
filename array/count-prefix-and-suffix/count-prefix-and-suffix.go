package main

import (
	"fmt"
	"strings"
)

func countPrefixSuffixPairs(words []string) int {
	isPrefixAndSuffix := func(str1, str2 string) bool {
		return strings.HasPrefix(str2, str1) && strings.HasSuffix(str2, str1)
	}

	counter := 0

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				counter++
			}
		}
	}

	return counter
}

func main() {
	words := []string{"abcd", "efgh", "ijkl", "pqr", "stu", "vwxy"}
	fmt.Println(countPrefixSuffixPairs(words)) // 0

	words = []string{"a", "aba", "ababa", "aa"}
	fmt.Println(countPrefixSuffixPairs(words)) // 4
}
