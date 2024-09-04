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
}
