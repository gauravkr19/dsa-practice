// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
// You must write an algorithm that runs in O(n) time.
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.

// ## Aim to find the start of the sequence with the longest streak

// Step-by-Step Execution:
// Initial Setup:

// Convert the array nums to a set numSet for O(1) lookups.
// numSet will be: {100, 4, 200, 1, 3, 2}.
// Initialize longestStreak = 0.
// Iterate Through Each Number in nums:

// We will iterate through each number in the nums array and determine if it's the start of a new consecutive sequence.
// Processing Each Element:

// For num = 100:

// Check if 100 - 1 = 99 is in numSet. It is not, so 100 is the start of a new sequence.
// Initialize currentNum = 100 and currentStreak = 1.
// Check if 101 is in numSet (i.e., currentNum + 1). It's not, so the sequence ends.
// Update longestStreak = max(1, 0) = 1.

// For num = 4:

// Check if 4 - 1 = 3 is in numSet. It is, so 4 is not the start of a new sequence.
// Skip to the next number.

// For num = 200:

// Check if 200 - 1 = 199 is in numSet. It is not, so 200 is the start of a new sequence.
// Initialize currentNum = 200 and currentStreak = 1.
// Check if 201 is in numSet. It's not, so the sequence ends.
// Update longestStreak = max(1, 1) = 1.

// For num = 1:

// Check if 1 - 1 = 0 is in numSet. It is not, so 1 is the start of a new sequence.
// Initialize currentNum = 1 and currentStreak = 1.
// Check if 2 is in numSet. It is, so increment currentNum to 2 and currentStreak to 2.
// Check if 3 is in numSet. It is, so increment currentNum to 3 and currentStreak to 3.
// Check if 4 is in numSet. It is, so increment currentNum to 4 and currentStreak to 4.
// Check if 5 is in numSet. It's not, so the sequence ends.
// Update longestStreak = max(4, 1) = 4.

// For num = 3:

// Check if 3 - 1 = 2 is in numSet. It is, so 3 is not the start of a new sequence.
// Skip to the next number.

// For num = 2:

// Check if 2 - 1 = 1 is in numSet. It is, so 2 is not the start of a new sequence.
// Skip to the next number.
// Final Result:

// After iterating through all the numbers, the longest sequence found is [1, 2, 3, 4] with a length of 4.
// The function will return 4.
package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	// Edge case: if the array is empty, return 0
	if len(nums) == 0 {
		return 0
	}

	// Convert the array to a set for O(1) lookups
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestStreak := 0

	// Iterate through each number in the array
	for _, num := range nums {
		// Check if num is the start of a sequence
		if !numSet[num-1] { // if num-1 is not in the set, num is the start of a sequence
			currentNum := num
			currentStreak := 1

			// Count the length of the sequence
			for numSet[currentNum+1] { // while currentNum+1 is in the set
				currentNum++
				currentStreak++
			}

			// Update the longest streak
			if currentStreak > longestStreak {
				longestStreak = currentStreak
			}
		}
	}

	return longestStreak
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println("The length of the longest consecutive elements sequence is:", longestConsecutive(nums))
}
