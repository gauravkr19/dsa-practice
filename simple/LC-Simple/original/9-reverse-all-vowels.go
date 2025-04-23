// Given a string s, reverse only all the vowels in the string and return it.
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases, more than once.

// Input: s = "hello"
// Output: "holle"
// Example 2:
// Input: s = "leetcode"
// Output: "leotcede"

package main

import "fmt"

func reverseVowels(str string) string {
	vowels := "aeiouAEIOU"
	vowelSet := make(map[byte]bool)

	// Populate the vowel set for efficient vowel checking
	for i := range vowels {
		vowelSet[vowels[i]] = true
	}

	// Convert string to a slice of characters for easy swapping
	characters := []byte(str)

	// Use two pointers to swap vowels from the two ends towards the center
	i, j := 0, len(characters)-1

	// Outer loop is O(n) as each char is processed at most once
	for i < j {

		// inner loop skips non-vowel chars
		// In the worst case (all non-vowel characters), the inner loops will also traverse the string once from both ends.
		// This does not increase the time complexity beyond O(n), because the total work is proportional to the string length.
		for i < j && !vowelSet[characters[i]] {
			i++
		}
		for i < j && !vowelSet[characters[j]] {
			j--
		}

		// Swap vowels
		characters[i], characters[j] = characters[j], characters[i]
		i++
		j--
	}

	// Convert the modified slice of characters back to a string
	return string(characters)
}

func main() {
	fmt.Println(reverseVowels("helloo"))
}
