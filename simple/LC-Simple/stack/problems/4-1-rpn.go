package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, token := range tokens {
		if calculate, exist := operators[token]; exist {
			// Ensure at least two elements are available for the operation
			if len(stack) < 2 {
				panic("Invalid expression: not enough operands")
			}

			// Pop last two elements (op2 first, then op1)
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			// Perform calculation and push result back
			stack = append(stack, calculate(a, b))
		} else {

			// Convert string to integer with error handling
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	// Final stack must contain exactly one element (the result)
	if len(stack) != 1 {
		panic("Invalid expression: stack should have exactly one result")
	}
	return stack[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
