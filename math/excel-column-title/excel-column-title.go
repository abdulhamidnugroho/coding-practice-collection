package main

import (
	"fmt"
)

func convertToTitle(columnNumber int) string {
	result := ""

	for columnNumber > 0 {
		columnNumber--

		remainder := columnNumber % 26 // modulo 26

		// convert remainder to corresponding character
		result = string('A'+remainder) + result // most important code
		// move to the next digit
		columnNumber /= 26
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(convertToTitle(1))          // "A"
	fmt.Println(convertToTitle(28))         // "AB"
	fmt.Println(convertToTitle(701))        // "ZY"
	fmt.Println(convertToTitle(2147483647)) // FXSHRXW
}
