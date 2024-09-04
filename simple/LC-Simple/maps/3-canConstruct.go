// Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.
// Example 1:
// Input: ransomNote = "a", magazine = "b"
// Output: false
// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
// Output: false
// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
// Output: true

package main

import (
	"fmt"
	"strings"
)

func canConstructSlice(s1, s2 string) bool {
	freqMap := make(map[string]int)
	charSlice1 := strings.Split(s1, "")
	charSlice2 := strings.Split(s2, "")

	for _, char := range charSlice2 {
		freqMap[char]++
	}

	for _, char := range charSlice1 {
		if freqMap[char] > 0 {
			freqMap[char]--
		} else {
			return false
		}
	}
	return true
}

func canConstructRune(s1, s2 string) bool {
	freqMap := make(map[rune]int)
	// charSlice1 := strings.Split(s1, "")
	// charSlice2 := strings.Split(s2, "")

	for _, char := range s2 {
		freqMap[char]++
	}

	for _, char := range s1 {
		if freqMap[char] > 0 {
			freqMap[char]--
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(canConstructSlice("abcd", "baca"))
	fmt.Println(canConstructRune("x", "baca"))
}
