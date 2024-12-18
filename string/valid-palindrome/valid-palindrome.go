package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	str := ""

	// normalize input: only keep alphanumeric characters and convert to lowercase
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') || (v >= '0' && v <= '9') {
			str += strings.ToLower(string(v))
		}
	}

	// check if the string is a palindrome
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}

func isPalindromeTwo(s string) bool {
	var filtered []rune

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			filtered = append(filtered, unicode.ToLower(r))
		}
	}

	n := len(filtered)
	for i := 0; i < n/2; i++ {
		if filtered[i] != filtered[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	s2 := "0P"

	fmt.Println(isPalindrome(s))  // Output: true
	fmt.Println(isPalindrome(s2)) // Output: false

	fmt.Println(isPalindromeTwo(s))  // Output: true
	fmt.Println(isPalindromeTwo(s2)) // Output: false
}
