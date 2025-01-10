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

func IsAnagram(s, t string) bool {

	countMap := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		countMap[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		countMap[t[i]]--
		if countMap[t[i]] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(canConstructSlice("abcd", "baca"))
	fmt.Println(canConstructRune("x", "baca"))
	fmt.Println(canConstructRune("x", "baca"))
}
