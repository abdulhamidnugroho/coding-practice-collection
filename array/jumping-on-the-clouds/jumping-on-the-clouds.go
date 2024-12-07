package main

import "fmt"

func jumpingOnClouds(c []int32) int32 {
	var res int32

	for i := 0; i < len(c)-1; i++ {
		if c[i] == 0 {
			i++
		}
		res++
	}
	return res
}

func main() {
	c := []int32{0, 1, 0, 0, 0, 1, 0}

	result := jumpingOnClouds(c)

	fmt.Printf("Clouds: %v\n", c)
	fmt.Printf("Minimum number of jumps: %d\n", result)
}
