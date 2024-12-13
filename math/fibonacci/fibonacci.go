package main

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	first, sec := 0, 1

	for i := 2; i <= n; i++ {
		first, sec = sec, first+sec
	}

	return sec
}

func main() {
	fib1 := 1
	fib3 := 3
	fib6 := 6

	fmt.Println(fib(fib1)) // Output: 0
	fmt.Println(fib(fib3)) // Output: 2
	fmt.Println(fib(fib6)) // Output: 8
}
