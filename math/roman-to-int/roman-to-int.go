package main

import "fmt"

func romanToInt(s string) int {
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	res := 0

	for i := 0; i < n; i++ {
		value := values[rune(s[i])]

		if i < n-1 && value < values[rune(s[i+1])] {
			res -= value
		} else {
			res += value
		}
	}

	return res
}

func main() {
	roman := "MCMXCVII"
	fmt.Println(romanToInt(roman)) // Output: 1997
}
