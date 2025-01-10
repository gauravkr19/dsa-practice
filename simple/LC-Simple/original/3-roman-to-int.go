package main

import (
	"fmt"
)

func romanToInt(s string) int {

	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	total := 0

	// Iterate through the string
	for i := 0; i < n; i++ {
		// Get the value of the current Roman numeral
		value := romanMap[s[i]]

		// Check the next numeral (if exists) for subtraction rule
		if i+1 < n && value < romanMap[s[i+1]] {
			total -= value // Apply subtraction
		} else {
			total += value // Apply addition
		}
	}

	return total
}

func main() {
	fmt.Println(romanToInt("III"))     // Output: 3
	fmt.Println(romanToInt("IV"))      // Output: 4
	fmt.Println(romanToInt("IX"))      // Output: 9
	fmt.Println(romanToInt("LVIII"))   // Output: 58
	fmt.Println(romanToInt("MCMXCIV")) // Output: 1994
	fmt.Println(romanToInt("XXVII"))   // Output: 27
	romanNumeral := "XXVII"
	result := romanToInt(romanNumeral)
	fmt.Printf("The integer representation of %s is: %d\n", romanNumeral, result)
}
