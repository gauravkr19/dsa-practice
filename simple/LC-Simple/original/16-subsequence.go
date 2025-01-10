package main

import "fmt"

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
// of the characters without disturbing the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

// when using range ---> for i, s := range t Error: "invalid operation: cannot index s (variable of type rune)""
func subsequence(s string, t string) bool {

	sIndex := 0 // use sIndex to track the elements of slice s
	for i := 0; i < len(t) && sIndex < len(s); i++ {

		if t[i] == s[sIndex] {
			sIndex++
		}
	}

	return sIndex == len(s)
}

func main() {
	fmt.Println(subsequence("abc", "asdsbdsdc"))
	fmt.Println(subsequence("axc", "asdsbdsdc"))
}
