// Your program should accept comma separated alphanumeric characters ( 15-20) ,
// print those characters with number of times repeated and their position in sequence.

// example
// Input =  comma separated alphanumeric characters = A,c,f,4,6,j,1,0,1,c,A

// Expected Output:
// A : 2 : 0,10
package main

import (
	"fmt"
	"strings"
)

func main() {
	// Get input from the user
	fmt.Print("Enter comma-separated alphanumeric characters (15-20): ")
	var input string
	fmt.Scan(&input)

	// Split the input into a slice of strings
	characters := strings.Split(input, ",")

	// Initialize a map to store the count and positions of each character
	charCount := make(map[string]int)
	charPositions := make(map[string][]int)

	// Iterate through the characters and update count and positions
	for i, char := range characters {
		// Trim any leading or trailing whitespaces
		char = strings.TrimSpace(char)

		// This increments the count of the character char in the charCount map.
		// If char is encountered for the first time, it will be automatically added to the map with a count of 1.
		// If it already exists, its count will be incremented by 1.
		charCount[char]++

		// This appends the current position i of the character char to the slice in the charPositions map.
		// If char is encountered for the first time, a new slice is created, and the position i is added to it.
		// If it already exists, the position i is appended to the existing slice of positions.
		charPositions[char] = append(charPositions[char], i)
	}

	// fmt.Println(charCount)
	// fmt.Println(charPositions)

	// Print the results
	for char, count := range charCount {
		positions := charPositions[char]
		fmt.Printf("%s : %d : %v\n", char, count, positions)
	}
}

// A : 2 : 0,10

// convert slice of integer to slice of string
