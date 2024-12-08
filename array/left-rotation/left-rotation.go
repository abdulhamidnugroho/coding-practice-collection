package main

import "fmt"

// rotLeft performs d left rotations on the array a.
func rotLeft(arr []int32, d int32) []int32 {
	n := len(arr)
	d %= int32(n) // Handle cases where d >= n
	// Concatenate the array from index d to the end with the beginning up to index d.
	rotated := append(arr[d:], arr[:d]...)
	return rotated
}

func main() {
	// Example input
	a := []int32{1, 2, 3, 4, 5}
	d := int32(3)

	// Perform the left rotation
	result := rotLeft(a, d)

	// Output the result
	fmt.Println("Rotated Array:", result)
}
