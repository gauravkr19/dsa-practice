// Given two strings s and t, return true if t is an anagram of s, and false otherwise.

// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

package main

import "fmt"

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
	fmt.Println(IsAnagram("anagrama", "nagarama"))
}

// Time Complexity
//Counting Approach:  O(n), where 𝑛 n is the length of the strings, This is more efficient than the O(n logn) sorting method.
