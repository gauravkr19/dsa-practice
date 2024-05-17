package main

import (
	"fmt"
	"strings"
)

// CountAndPositions calculates the count and positions of each unique character in the input string.
func CountAndPositions(input string) map[string][]int {
	// Split the input into a slice of strings
	characters := strings.Split(input, ",")

	// Initialize a map to store the count and positions of each character
	charPositions := make(map[string][]int)

	// Iterate through the characters and update positions
	for i, char := range characters {
		// Trim any leading or trailing whitespaces
		char = strings.TrimSpace(char)

		// Update positions
		charPositions[char] = append(charPositions[char], i)
	}

	fmt.Println(charPositions)
	fmt.Println()

	return charPositions
}

func main() {
	// Get input from the user
	fmt.Print("Enter comma-separated alphanumeric characters (15-20): ")
	var input string
	fmt.Scan(&input)

	// Calculate count and positions using the function
	result := CountAndPositions(input)

	// Print the results
	for char, positions := range result {
		fmt.Printf("%s : %d : %v\n", char, len(positions), positions)
	}
}
