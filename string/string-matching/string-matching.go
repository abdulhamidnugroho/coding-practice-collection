package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	res := []string{}
	wholeString := strings.Join(words, "!")

	for _, v := range words {
		if strings.Count(wholeString, v) > 1 {
			res = append(res, v)
		}
	}

	return res
}

func main() {
	words := []string{"mass", "as", "hero", "superhero"}
	res := stringMatching(words)
	fmt.Println(res)

	words = []string{"leetcode", "et", "code"}
	res = stringMatching(words)

	fmt.Println(res)
}
