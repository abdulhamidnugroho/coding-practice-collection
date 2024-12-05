package main

import "fmt"

func sockMerchant(n int32, ar []int32) int32 {
	count := make(map[int]int)
	pairs := 0

	for _, v := range ar {
		count[int(v)]++
		if count[int(v)]%2 == 0 {
			pairs++
		}
	}

	return int32(pairs)
}

func main() {
	var n int32

	// Reading the number of socks
	fmt.Println("Enter the number of socks:")
	fmt.Scan(&n)

	ar := make([]int32, n)
	fmt.Println("Enter the colors of the socks:")
	for i := 0; int32(i) < n; i++ {
		fmt.Scan(&ar[i])
	}

	// Calculate and print the number of pairs
	result := sockMerchant(n, ar)
	fmt.Println("Number of pairs:", result)
}
