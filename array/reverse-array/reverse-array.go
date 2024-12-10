package main

import "fmt"

func reverseArray(a []int32) []int32 {
	b := make([]int32, len(a))

	for i, v := range a {
		b[len(a)-1-i] = v
	}

	return b
}

func main() {
	arr := []int32{1, 2, 3, 4, 5}
	reversed := reverseArray(arr)

	fmt.Println(reversed) // Output: [5 4 3 2 1]
}
