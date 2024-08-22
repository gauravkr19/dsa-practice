package main

import (
	"fmt"
	"strconv"
	"strings"
)

// example
// Input =  comma separated alphanumeric characters = A,c,f,4,6,j,1,0,1,c,A

// CountAndPositions calculates the count and positions of each unique character in the input string.
func CountAndPositions(input string) map[string][]int {
	// Split the input into a slice of strings
	characters := strings.Split(input, ",")

	charPositions := make(map[string][]int)

	for i, char := range characters {
		char = strings.TrimSpace(char)

		// charPositions[char] = ... updates the map with the new slice containing the appended index.
		charPositions[char] = append(charPositions[char], i)
		// append(charPositions[char] - accesses the slice of indices associated with the  char.
	}

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
	for char, positions := range result { //positions is []int
		// fmt.Printf("%s : %d : %v\n", char, len(positions), positions)

		var csvElements []string
		for _, element := range positions {
			csvElements = append(csvElements, strconv.Itoa(element))
		}
		csvPositions := strings.Join(csvElements, ",")
		fmt.Printf("%s : %d : %v\n", char, len(positions), csvPositions)
	}
}
