package main

import "strings"

func prefixCount(words []string, pref string) int {
	counter := 0
	for _, v := range words {
		if strings.HasPrefix(v, pref) {
			counter++
		}
	}

	return counter
}

func main() {
	words := []string{"pay", "attention", "practice", "attend"}
	pref := "app"
	count := prefixCount(words, pref)
	println(count) // 2
}
