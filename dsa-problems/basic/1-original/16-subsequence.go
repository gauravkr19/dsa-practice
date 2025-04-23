package main

import "fmt"

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
// A subsequence of a string is a new string that is formed from the original string by deleting some (can be none)
// of the characters without disturbing the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).

// It's a common two-pointer pattern — one for each string — and a key takeaway you should definitely remember for interviews.
// You're checking if all characters in s appear in t, 'in order' — but 'not necessarily contiguously'.
// So as you walk through t, you need to 'remember how much of s you've matched so far.' That's exactly what sIndex is for.

// You need to scan t fully, but only move forward in s when there's a match

// **Algo**
// Think of it as a "match progress tracker":
// Start at the first char of s
// Scan through t (using t[i] == ...)
// When you see a match: move sIndex forward
// If sIndex == len(s) by the end → All characters matched in order → ✅ subsequence

// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted.

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
