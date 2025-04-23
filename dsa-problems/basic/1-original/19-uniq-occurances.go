// Given an array of integers arr,
// return true if the number of occurrences of each value in the array is unique or false otherwise.

// Example 1:

// Input: arr = [1,2,2,1,1,3]
// Output: true
// Step 2: Check if occurrences are unique

// We then create another map elementCount, where the key is the count of occurrences (from the previous step),
// and the value is a boolean flag indicating whether this count has already been seen.
// We iterate through the counts from elementOccurances, and for each count, we check if this count has already been recorded in elementCount.
// If it has, return false (since two elements have the same count, meaning the occurrences are not unique).
// If it hasn’t been seen, we mark it as seen by setting elementCount[count] = true.
// If we complete the iteration without finding any duplicate counts, return true.

// Iterating through the counts of elementOccurances:
// For count 3: elementCount becomes {3: true}
// For count 2: elementCount becomes {3: true, 2: true}
// For count 1: elementCount becomes {3: true, 2: true, 1: true}
// Since all counts are unique, return true.

package main

import "fmt"

func UniqOccurances(nums []int) bool {

	// Get the occurances in first loop
	elementOccurances := make(map[int]int)
	for _, num := range nums {
		elementOccurances[num]++
	}

	elementCount := make(map[int]bool)

	// Iterate through elementOccurances map and check duplicate value
	for _, count := range elementOccurances {
		if elementCount[count] {
			return false
		}
		elementCount[count] = true
	}

	return true
}

func main() {
	fmt.Println(UniqOccurances([]int{1, 2, 2, 1, 1, 3, 3}))
}
