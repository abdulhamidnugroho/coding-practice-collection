package main

import "fmt"

func isIsomorphic(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS, mapT := [256]int{}, [256]int{}

	for i := 0; i < len(s); i++ {
		charS, charT := s[i], t[i]

		if mapS[charS] != mapT[charT] {
			return false
		}

		mapS[charS] = i + 1
		mapT[charT] = i + 1
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))     // true
	fmt.Println(isIsomorphic("foo", "bar"))     // false
	fmt.Println(isIsomorphic("paper", "title")) // true
}
