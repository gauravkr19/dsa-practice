// Inversion Brute O(n^2)
// Inversion Recursive O (n logn)

package main

import (
	"fmt"
)

// mergeAndCount merges two sorted subarrays and counts split inversions.
func mergeAndCount(left, right []int) ([]int, int) {
	merged := make([]int, 0, len(left)+len(right)) // Merged array
	i, j := 0, 0
	splitInvCount := 0
	leftSize := len(left)

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			merged = append(merged, left[i]) // No inversion, take from left
			i++
		} else {
			merged = append(merged, right[j]) // Inversion found
			splitInvCount += (leftSize - i)   // Count all remaining elements in left[]
			j++
		}
	}

	// Append any remaining elements
	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged, splitInvCount
}

// sortAndCount sorts the array and counts the number of inversions
func sortAndCount(arr []int) ([]int, int) {
	if len(arr) <= 1 {
		return arr, 0 // Base case: No inversions in a single-element array
	}

	mid := len(arr) / 2

	// Recursively sort left and right halves and count their inversions
	left, leftInv := sortAndCount(arr[:mid])   // (A) Gets sorted left half & its inversion count
	right, rightInv := sortAndCount(arr[mid:]) // (B) Gets sorted right half & its inversion count

	// Merge sorted halves and count split inversions
	merged, splitInv := mergeAndCount(left, right) // (C) Counts inversions during merge

	return merged, leftInv + rightInv + splitInv // (D) Total inversions in this step
}

func main() {
	arr := []int{2, 3, 8, 6, 1}
	sortedArr, invCount := sortAndCount(arr)
	fmt.Println("Sorted Array:", sortedArr)
	fmt.Println("Number of Inversions:", invCount)
}
