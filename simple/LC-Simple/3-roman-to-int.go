package main

import (
	"fmt"
)

func romanToInt(str string) int {
	romanValues := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(str); i++ {
		currentValue := romanValues[str[i]]
		nextValue := 0

		// Check if there is a next character
		if i+1 < len(str) {
			nextValue = romanValues[str[i+1]]
		}

		// If the current value is less than the next value, subtract it
		if currentValue < nextValue {
			result += nextValue - currentValue
			i++ // Skip the next character since it has been considered
		} else {
			result += currentValue
		}
	}

	return result
}

func main() {
	// Example usage:
	romanNumeral := "XXVII"
	result := romanToInt(romanNumeral)
	fmt.Printf("The integer representation of %s is: %d\n", romanNumeral, result)
}
