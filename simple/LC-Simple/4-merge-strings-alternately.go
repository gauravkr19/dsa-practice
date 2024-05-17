package main

import (
	"fmt"
	"math"
)

func MergeAlternate(word1 string, word2 string) string {
	merged := ""

	for i := 0; i < int(math.Max(float64(len(word1)), (float64(len(word2))))); i++ {

		if i < len(word1) {
			merged += string(word1[i])
		}

		if i < len(word2) {
			merged += string(word2[i])
		}
	}
	return merged
}

func main() {
	fmt.Println(MergeAlternate("a", "xyz"))
	fmt.Println(MergeAlternate("ab", "pqrs"))
	fmt.Println(MergeAlternate("abcd", "pq"))
}

// You are given two strings word1 and word2. Merge the strings by adding letters in alternating order,
// starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.
// Return the merged string.
