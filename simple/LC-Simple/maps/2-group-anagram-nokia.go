// Employee Given a sequence of words in a array, print all anagrams together

// INPUT [“cat”, “dog”, “tac”, “god”, “act”] {"cat", "dog", "tac", "god", "act"}
// OUTPUT[[“cat”, “tac”,“act”],[ “dog”,“god”]]
// a word or phrase that is made by arranging the letters of another word or phrase in a different order.
// HINT:  We loop through each word, sort it, and add it to the corresponding group in the map.
// Steps: 	1. input= ["cat", "dog", "tac", "god", "act]
// 			2. range the slice to extract each word, example cat
// 			3. pass this each word to a func to sort the word in ascending order
// 			Sorting logic: calls sorting func and pass string word, example "cat" that returns string "act" in asc order
//					i.   Use strings.Split to convert string to []string. This creates [c a t]
//					ii.  Use sort.Strings which accepts []string, [c a t] and returns sorted slice [a c t]
//					iii. Use strings.Join which accepts []string and converts to string "act"
//

package main

import (
	"fmt"
	"sort"
	"strings"
)

// Function to sort the characters in a string
func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

// Function to group anagrams
func groupAnagrams(words []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, word := range words {
		// Sort the word and use the sorted word as the key
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	fmt.Println(anagramMap)
	// Convert the map values to a slice of slices
	// The subsequent loop iterates over the anagramMap, where each group is a slice of strings (e.g., ["cat", "tac", "act"]).
	// The append function adds each group to the result slice.
	result := [][]string{}
	for _, group := range anagramMap {
		result = append(result, group)
	}
	return result
}

func main() {
	words := []string{"cat", "dog", "tac", "god", "act"}
	groupedAnagrams := groupAnagrams(words)
	fmt.Println(groupedAnagrams)
}
