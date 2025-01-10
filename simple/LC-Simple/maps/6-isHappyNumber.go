// Write an algorithm to determine if a number n is happy.

// A happy number is a number defined by the following process:

// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.
// Input: n = 19
// Output: true
// Explanation:
// 12 + 92 = 82
// 82 + 22 = 68
// 62 + 82 = 100
// 12 + 02 + 02 = 1

package main

import (
	"fmt"
)

func isHappy(num int) bool {
	seen := make(map[int]bool)

	for num != 1 {
		if seen[num] {
			return false
		}

		seen[num] = true
		num = sumOfSquares(num)
	}

	return true
}

func sumOfSquares(num int) int {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
}

// Important Notes: if  seen[num] returns value
// In the case above, the if condition doesn't explicitly check the boolean value that indicates whether the key exists.
//  It implicitly checks the value stored at seen[num] and uses that for the condition.
// If you want to check whether the key exists in the map explicitly, you can use the two-value assignment:
// if _, exists := seen[num]; exists
