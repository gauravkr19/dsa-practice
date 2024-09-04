// Given two strings s and t, determine if they are isomorphic.
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
// All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character, but a character may map to itself.
// Example 1:
// Input: s = "egg", t = "add"
// Output: true
// Explanation:
// The strings s and t can be made identical by:
// Mapping 'e' to 'a'.
// Mapping 'g' to 'd'.
// Example 2:
// Input: s = "foo", t = "bar"
// Output: false
// Explanation:
// The strings s and t can not be made identical as 'o' needs to be mapped to both 'a' and 'r'.
// Example 3:
// Input: s = "paper", t = "title"
// Output: true

package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	// Initialize two maps for character mapping
	map_s_to_t := make(map[byte]byte)
	map_t_to_s := make(map[byte]byte)

	// Iterate over the characters of both strings simultaneously
	for i := 0; i < len(s); i++ {
		char_s := s[i]
		char_t := t[i]

		// Check if there's an existing mapping for char_s in map_s_to_t
		if mappedChar, exists := map_s_to_t[char_s]; exists {
			if mappedChar != char_t {
				return false
			}
		} else {
			map_s_to_t[char_s] = char_t
		}

		// Check if there's an existing mapping for char_t in map_t_to_s
		if mappedChar, exists := map_t_to_s[char_t]; exists {
			if mappedChar != char_s {
				return false
			}
		} else {
			map_t_to_s[char_t] = char_s
		}
	}

	// If we reach here, the strings are isomorphic
	return true
}

func main() {
	s1 := "foo"
	t1 := "bar"
	fmt.Println(isIsomorphic(s1, t1)) // Output: false

	s2 := "paper"
	t2 := "title"
	fmt.Println(isIsomorphic(s2, t2)) // Output: true
}
