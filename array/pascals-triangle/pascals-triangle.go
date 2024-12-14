package main

import "fmt"

// generate generates Pascal's Triangle for the given number of rows
func generate(numRows int) [][]int {
	res := [][]int{}

	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}

		res = append(res, row)
	}

	return res
}

func printPascalTriangle(triangle [][]int) {
	for _, row := range triangle {
		fmt.Println(row)
	}
}

func main() {
	// Example 1
	numRows1 := 5
	fmt.Println("Pascal's Triangle with 5 rows:")
	triangle1 := generate(numRows1)
	printPascalTriangle(triangle1)

	// Example 2
	numRows2 := 1
	fmt.Println("\nPascal's Triangle with 1 row:")
	triangle2 := generate(numRows2)
	printPascalTriangle(triangle2)
}
