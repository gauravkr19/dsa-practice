// Fundamental
// 		Slices
// 		strings
// 		Sorting
// 		Two Pointer
// 		Stack
// 		Linked List
// 		Matrix
// 		Simulation

// Intermediate
// 		Maps / Hash Table
// 		Depth-First Search
// 		Breadth-First Search
// 		Tree
// 		Binary Tree
// 		Binary Search
// 		Greedy
// 		Math

// Advanced
// 		Dynamic Programming
// 		Backtracking
// 		Union Find
// 		Divide and Conquer
// 		Monotonic Stack
// 		Data Stream
// 		Trie
// 		Topological Sort

package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Test")
	// str := "abcdef"
	str := []byte("abcdef")

	// strRune := []rune(str)
	l := len(str)

	for i := 0; i < l/2; i++ {
		j := l - i - 1

		str[i], str[j] = str[j], str[i]
	}

	fmt.Printf("%s \n", string(str))
}
