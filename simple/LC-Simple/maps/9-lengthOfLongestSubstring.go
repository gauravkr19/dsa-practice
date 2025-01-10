// Given a string s, find the length of the longest substring without repeating characters.
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	// Map to store the last index of each character
	charMap := make(map[byte]int)

	// Initialize the start of the window and the max length
	start, maxLength := 0, 0

	// Iterate over the string
	for end := 0; end < len(s); end++ {
		char := s[end]

		// If the character is already in the map and its index is within the current window, abcabcbb
		if lastIndex, exists := charMap[char]; exists && lastIndex >= start {
			// Move the start of the window to exclude the repeated character
			start = lastIndex + 1
		}

		// Update the last seen index of the character
		charMap[char] = end

		// Calculate the length of the current substring
		maxLength = max(maxLength, end-start+1)
	}

	return maxLength
}

// Helper function to find the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test cases
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
	fmt.Println(lengthOfLongestSubstring(""))         // Output: 0
}

// Sliding Window Concept:

// Use two pointers (start and end) to represent the current substring (window) under consideration.
// Expand the window by moving end to the right, character by character.
// Adjust the start pointer to shrink the window whenever a duplicate character is encountered, ensuring the substring remains valid (i.e., without repeating characters).
// Tracking Characters:

// Use a hash map (or slice for ASCII strings) to store the last seen index of each character in the current substring.
// This helps efficiently detect duplicate characters and decide how much to shrink the window.
// Calculate Maximum Length:

// At each step, calculate the length of the current substring (end - start + 1) and update the maximum length seen so far.
