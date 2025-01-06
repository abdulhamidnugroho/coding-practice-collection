package main

import "reflect"

func isAnagram(s string, t string) bool {
	counter := make(map[string]int, len(s))

	for _, v := range s {
		if currentVal, ok := counter[string(v)]; ok {
			counter[string(v)] = currentVal + 1
			continue
		}

		counter[string(v)] = 1
	}

	counter2 := make(map[string]int, len(t))
	for _, v := range t {
		if currentVal, ok := counter2[string(v)]; ok {
			counter2[string(v)] = currentVal + 1
			continue
		}

		counter2[string(v)] = 1
	}

	if reflect.DeepEqual(counter, counter2) {
		return true
	}

	return false
}

func main() {
	s := "anagram"
	t := "nagaram"

	result := isAnagram(s, t)

	if result {
		println("The strings are anagrams") // true
	} else {
		println("The strings are not anagrams") // false
	}
}
