package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	// Trim any trailing spaces
	s = strings.TrimRight(s, " ")

	// Find the last space in the string
	lastSpace := strings.LastIndex(s, " ")

	// If no space is found, the whole string is the last word
	if lastSpace == -1 {
		return len(s)
	}

	// Return the length of the last word
	return len(s) - lastSpace - 1
}

func main() {
	s := "Hello World"
	fmt.Println(lengthOfLastWord(s)) // Output: 5 (World)

	s = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s)) // Output: 4 (moon)
}
