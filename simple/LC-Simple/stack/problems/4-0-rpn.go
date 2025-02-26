// You are given an array of strings tokens that represents an arithmetic expression in a Reverse Polish Notation.
// Evaluate the expression. Return an integer that represents the value of the expression.
// Note that:
// The valid operators are '+', '-', '*', and '/'.
// Each operand may be an integer or another expression.
// The division between two integers always truncates toward zero.
// Input: tokens = ["2","1","+","3","*"]
// Output: 9
// Explanation: ((2 + 1) * 3) = 9

package main

// type stack struct {
// 	result []int
// }

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	result := []int{}

	for _, token := range tokens {
		if token != "+" && token != "-" && token != "*" && token != "/" {
			num, _ := strconv.Atoi(token)
			result = append(result, num)
		} else {
			op2 := result[len(result)-1]
			op1 := result[len(result)-2]
			result = result[:len(result)-2]
			switch token {
			case "+":
				sum := op1 + op2
				result = append(result, sum)
			case "-":
				sub := op1 - op2
				result = append(result, sub)
			case "*":
				product := op1 * op2
				result = append(result, product)
			case "/":
				div := op1 / op2
				result = append(result, div)
			}

		}
	}
	return result[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
}
