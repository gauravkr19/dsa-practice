// Selection Sort is a simple comparison-based sorting algorithm. The idea is to divide the array into two parts: a sorted part and an unsorted part.
// The algorithm repeatedly selects the smallest (or largest) element from the unsorted part and swaps it with the first unsorted element, expanding the sorted part by one element after each iteration.

// Logic and Steps:
// Initialize: Start with the first element in the array.
// Find the Minimum: Search through the unsorted part of the array to find the smallest element. [innes loop]
// Swap: Swap the smallest element with the first unsorted element.
// Expand Sorted Portion: Move the boundary between the sorted and unsorted parts by one element.
// Repeat: Repeat the process for the remaining unsorted portion of the array until all elements are sorted.

// Explanation of Code:
// Outer Loop (i): This loop goes through each element in the array. After each iteration, the element at index i is sorted.
// Inner Loop (j): This loop starts from i + 1 and searches for the smallest element in the unsorted portion of the array.
// Finding Minimum (minIdx): minIdx keeps track of the index of the smallest element found in the unsorted portion. If a smaller element is found, minIdx is updated.
// Swap: After finding the minimum element, it is swapped with the first unsorted element (at index i), thus expanding the sorted portion by one element.

package main

import "fmt"

// 64, 25, 12, 22, 11
// i = min = 0
// j = 1
// min = 1
// 25, 64, 12, 22, 11
// i = min = 1
// j = 2
//  12 < 25, j=min = 2
// 25, 64, 12, 22, 11

// SelectionSort function to sort the array
func SelectionSort(arr []int) {
	n := len(arr)
	// Outer loop for each element in the array
	for i := 0; i < n-1; i++ {
		// Assume the current element is the minimum
		minIdx := i
		// Inner loop to find the minimum element in the unsorted portion
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				// Update minIdx if a smaller element is found
				minIdx = j
			}
		}
		// Swap only if the minimum element is not already in its place
		if minIdx != i {
			arr[i], arr[minIdx] = arr[minIdx], arr[i]
		}
	}
}

func main() {
	arr := []int{64, 25, 12, 22, 11}
	fmt.Println("Original array:", arr)

	SelectionSort(arr)

	fmt.Println("Sorted array:", arr)
}
