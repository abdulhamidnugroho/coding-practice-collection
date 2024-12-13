package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}

	return curr
}

func main() {
	n := 2
	n2 := 3
	n11 := 11

	fmt.Println(climbStairs(n))   // Output: 2
	fmt.Println(climbStairs(n2))  // Output:	3
	fmt.Println(climbStairs(n11)) // Output: 144 distinct ways
}
