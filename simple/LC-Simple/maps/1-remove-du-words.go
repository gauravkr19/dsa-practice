// Write a Go function that takes a string as input and returns a map where the keys are the distinct words in the string,
// and the values are the number of times each word appears. Assume words are separated by spaces.

// For the input string "hello world hello", the function should return a map like:
// map["hello": 2, "world": 1]

package main

import (
	"fmt"
	"strings"
)

func removeDupWords(str string) map[string]int {

	str = strings.TrimSpace(str)

	result := make(map[string]int)

	strSlice := strings.Split(str, " ")
	for _, word := range strSlice {

		result[word]++
	}

	return result
}

func removeDupChar(str string) map[string]int {
	str = strings.TrimSpace(str)
	result := make(map[string]int)

	for i := range str {
		result[string(str[i])]++
	}
	return result
}

func main() {
	fmt.Println(removeDupWords("    hello world hello word world "))

	fmt.Println(removeDupChar("  hello "))

	// Write a Go function that takes a string as input and returns a map where the keys are the distinct characters in the string and the values are the number of times each character appears.

	// Example:
	// For the input string "hello", the function should return a map like:

	// map[h:1 e:1 l:2 o:1]

}
