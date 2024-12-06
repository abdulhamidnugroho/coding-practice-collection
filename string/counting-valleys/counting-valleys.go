package main

import "fmt"

func countingValleys(steps int32, path string) int32 {
	valleys := int32(0)
	altitude := int32(0)

	for _, step := range path {
		if step == 'U' {
			altitude++
		} else if step == 'D' {
			altitude--
		}

		if altitude == 0 && step == 'U' {
			valleys++
		}
	}

	return valleys
}

func main() {
	steps := int32(8)
	path := "DDUUUUDD"

	result := countingValleys(steps, path)

	// Output the result
	fmt.Printf("Steps: %d\n", steps)
	fmt.Printf("Path: %s\n", path)
	fmt.Printf("Number of valleys: %d\n", result) // 1
}
