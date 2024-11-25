package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		char := strs[0][i] // Get the current character

		// Check if this character matches in all other strings
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != char {
				return strs[0][:i] // Return the prefix up to the previous character
			}
		}
	}

	// If the entire first string is a prefix, return it
	return strs[0]
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs)) // Output: "fl"

	strs = []string{"interspecies", "interstellar", "interstate"}
	fmt.Println(longestCommonPrefix(strs)) // Output: "inters"
}
