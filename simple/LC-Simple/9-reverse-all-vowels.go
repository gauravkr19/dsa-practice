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
	for i < j {
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
	fmt.Println(reverseVowels("hello"))
}
