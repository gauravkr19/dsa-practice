// Bubble Sort is a simple sorting algorithm that repeatedly steps through the list, compares adjacent elements, and swaps them if they are in the wrong order. The pass through the list is repeated until the list is sorted. The algorithm gets its name because smaller elements "bubble" to the top of the list (the beginning), while larger elements sink to the bottom (the end).

// Logic and Steps:
// Initialize: Start from the first element of the array.
// Compare: Compare the current element with the next element.
// Swap: If the current element is greater than the next element, swap them.
// Move to Next Pair: Move to the next pair of elements and repeat the process.
// Iterate: Continue the process for all elements. After each full pass through the array, the 'largest unsorted' element will be placed at its correct position at the end of the array.
// Repeat: Repeat the process for the remaining unsorted elements (excluding the last sorted element) until no swaps are required in a full pass through the array.

// 64, 34, 25, 12, 22, 11, 90
package main

import "fmt"

// BubbleSort function to sort the array
func BubbleSort(arr []int) {
	n := len(arr)
	// Outer loop for the number of passes
	for i := 0; i < n-1; i++ {
		// Flag to check if the array is already sorted
		swapped := false
		// Inner loop to compare adjacent elements
		for j := 0; j < n-i-1; j++ { // n-i-1 => ignores the sorted region of the array
			if arr[j] > arr[j+1] {
				// Swap if the current element is greater than the next
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// If no two elements were swapped in the inner loop, array is sorted
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Println("Original array:", arr)

	BubbleSort(arr)

	fmt.Println("Sorted array:", arr)
}

// Explanation of Code:
// Outer Loop: This loop ensures that the algorithm performs n-1 passes, where n is the length of the array.
// Inner Loop: The inner loop iterates through the array and compares adjacent elements, swapping them if they are in the wrong order.
// Swapped Flag: If no swaps are made during a pass, the array is already sorted, and the algorithm can terminate early.

// Worst Case: O(n²)  Occurs when the array is sorted in reverse order.
// Best Case: O(n) Occurs when the array is already sorted (optimized with the swapped flag).
// Average Case: O(n²)

// Detailed Explanation:
// The outer loop (for i := 0; i < n-1; i++) controls how many passes the algorithm makes. With each pass, the largest unsorted element "bubbles" to its correct position at the end of the array.
// After each pass of the outer loop, the last i elements of the array are sorted. Therefore, in the next pass, we don't need to check those elements again. This is why the inner loop only goes up to n-i-1, so it skips comparing the sorted portion of the array.
