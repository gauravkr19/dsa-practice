package main

import "fmt"

// 1. Linear Time Complexity (O(n)) - Simple Iterative Approach

// Algorithm:
// If the array has only one element, return that element.
// If the first or last element is a peak, return it.
// Iterate through the array from index 1 to n-2:
// If arr[i] is greater than both arr[i-1] and arr[i+1], return arr[i].

// Time Complexity: O(n) (worst case: scan the entire array)
// Space Complexity: O(1)

func findPeakLinear(arr []int) int {
	n := len(arr)
	if n == 1 {
		return arr[0]
	}
	if arr[0] >= arr[1] {
		return arr[0]
	}
	if arr[n-1] >= arr[n-2] {
		return arr[n-1]
	}

	for i := 1; i < n-1; i++ {
		if arr[i] >= arr[i-1] && arr[i] >= arr[i+1] {
			return arr[i]
		}
	}
	return -1 // Should never reach here
}

// 2. Logarithmic Time Complexity (O(log n)) - Binary Search Approach

// Algorithm:
// Pick the middle element, arr[mid].
// If arr[mid] is greater than both neighbors (arr[mid-1] and arr[mid+1]), return arr[mid].
// If arr[mid-1] > arr[mid], a peak must exist in the left half, so search in low...mid-1.
// If arr[mid+1] > arr[mid], a peak must exist in the right half, so search in mid+1...high.

// Reoccurance Relation is T(n) = T(n/2) + Θ(1)
// Time Complexity: O(log n)
// Space Complexity: O(log n) (due to recursion, but can be optimized to O(1) using iteration)

func findPeakBinary(arr []int, low, high int) int {
	mid := low + (high-low)/2
	n := len(arr)

	// Check if mid is a peak, use circuit-breakers when testing with 1st and last element
	if (mid == 0 || arr[mid] >= arr[mid-1]) &&
		(mid == n-1 || arr[mid] >= arr[mid+1]) {
		return arr[mid]
	}

	// If the left neighbor is greater, move left
	if mid > 0 && arr[mid-1] > arr[mid] {
		return findPeakBinary(arr, low, mid-1)
	}

	// Otherwise, move right
	return findPeakBinary(arr, mid+1, high)
}

func findPeakLogN(arr []int) int {
	return findPeakBinary(arr, 0, len(arr)-1)
}

// For an iterative O(log n) solution, we can replace recursion with a while loop:
// Time Complexity: O(log n)
// Space Complexity: O(1) (iterative version)

func findPeakIterative(arr []int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		// Check if mid is a peak, use circuit-breakers when testing with 1st and last element
		if (mid == 0 || arr[mid] >= arr[mid-1]) &&
			(mid == len(arr)-1 || arr[mid] >= arr[mid+1]) {
			return arr[mid]
		}

		if mid > 0 && arr[mid-1] > arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	arr := []int{3, 5, 8, 9, 2, 1}
	fmt.Println(findPeakLinear(arr))
	fmt.Println(findPeakLogN(arr))
	fmt.Println(findPeakIterative(arr))
}
