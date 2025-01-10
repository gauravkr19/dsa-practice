// Given a pattern and a string s, find if s follows the same pattern.
// Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in s. Specifically:
// Each letter in pattern maps to exactly one unique word in s.
// Each unique word in s maps to exactly one letter in pattern.
// No two letters map to the same word, and no two words map to the same letter.

// Example 1:
// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
// Explanation:
// The bijection can be established as:
// 'a' maps to "dog".
// 'b' maps to "cat"

package main

import (
	"fmt"
	"strings"
)

func WordPattern(pattern, s string) bool {
	PatternToWords := make(map[byte]string)
	WordsToPattern := make(map[string]byte)

	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		charPattern := pattern[i]
		word := words[i]

		if mappedWord, exists := PatternToWords[charPattern]; exists {
			if mappedWord != word {
				return false
			}
		} else {
			PatternToWords[charPattern] = word
		}

		if mappedChar, exists := WordsToPattern[word]; exists {
			if mappedChar != charPattern {
				return false
			}
		} else {
			WordsToPattern[word] = charPattern
		}
	}

	return true
}

func main() {
	fmt.Println(WordPattern("abba", "dog cat cat dog"))
	fmt.Println(WordPattern("abba", "dog cat cat fish"))
	fmt.Println(WordPattern("aaaa", "dog cat cat dog"))
	fmt.Println(WordPattern("abbac", "dog cat cat dog dog"))
}
