// Given an input string s, reverse the order of the words.
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
// Return a string of the words in reverse order concatenated by a single space.
// Note that s may contain leading or trailing spaces or multiple spaces between two words.
// The returned string should only have a single space separating the words. Do not include any extra spaces.

package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// Trim leading and trailing spaces
	s = strings.TrimSpace(s)

	// Split the string into words
	words := strings.Fields(s) //[a good example]
	fmt.Println(words)

	// Reverse the order of words
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	// Join the reversed words into a single string
	result := strings.Join(words, " ")

	return result
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))  // Output: "blue is sky the"
	fmt.Println(reverseWords("  hello world  "))  // Output: "world hello"
	fmt.Println(reverseWords("a good   example")) // Output: "example good a"

	fmt.Println("-----Calix Interview---------")
	fmt.Println(reverseWordsCharsInterview("the sky is blue"))  // Output: "example good a"
	fmt.Println(reverseWordsCharsInterview("a good   example")) // Output: "example good a"
}

// calix - reverse strings
func reverseWordsCharsInterview(str string) string {
	words := []byte(str)
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	return string(words)
}
