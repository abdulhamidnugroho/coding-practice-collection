package main

import "fmt"

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1 // first element is always 1

	for i := 1; i <= rowIndex; i++ {
		prev := row[i-1]
		row[i] = prev * (rowIndex - i + 1) / i
	}

	return row
}

func main() {
	rowIndex1 := 3
	fmt.Println("Row 3 of Pascal's Triangle:", getRow(rowIndex1))

	rowIndex2 := 0
	fmt.Println("Row 0 of Pascal's Triangle:", getRow(rowIndex2))

	rowIndex3 := 1
	fmt.Println("Row 1 of Pascal's Triangle:", getRow(rowIndex3))
}
