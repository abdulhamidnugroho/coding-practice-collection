package main

import "fmt"

func isValid(s string) bool {
	stack := []rune{}

	for _, char := range s {
		switch char {
		case '(': // Push corresponding closing bracket into the stack
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		default: // For closing brackets
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false // Mismatch or empty stack
			}
			stack = stack[:len(stack)-1] // Pop the stack
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))     // Output: true
	fmt.Println(isValid("()[]{}")) // Output: true
	fmt.Println(isValid("(]"))     // Output: false
	fmt.Println(isValid("([)]"))   // Output: false
}
