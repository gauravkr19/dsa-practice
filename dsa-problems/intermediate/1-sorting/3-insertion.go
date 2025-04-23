// Insertion Sort is a simple and intuitive sorting algorithm that works similarly to sorting playing cards in your hand. The array is divided into a sorted and an unsorted part.
// Initially, the first element is considered sorted.
// Then, the algorithm picks the next element from the unsorted part and inserts it into its correct position in the sorted part by shifting the other elements.
// This process repeats until all elements are sorted.

// Logic and Steps:
// Initialize: Assume the first element is already sorted.
// Pick Element: Starting from the second element, pick an element from the unsorted part.
// Find Correct Position: Compare the picked element with elements in the sorted part and shift elements to the right until you find the correct position.
// Insert Element: Insert the picked element at its correct position in the sorted part.
// Repeat: Repeat the process for the remaining unsorted elements until the entire array is sorted.

// Example:
// For the array [12, 11, 13, 5, 6]:

// First Iteration: Consider the first element 12 as sorted.
// Second Iteration: Pick 11, compare with 12, and insert 11 before 12 to get [11, 12].
// Third Iteration: Pick 13, it's already larger than the sorted elements, so no change.
// Fourth Iteration: Pick 5, insert it at the beginning to get [5, 11, 12, 13].
// Fifth Iteration: Pick 6, insert it between 5 and 11 to get [5, 6, 11, 12, 13].

// Explanation of Code:
// Outer Loop (i): Starts from the second element (i = 1), since the first element is already considered sorted.
// Key Element (key): This is the element to be inserted in the sorted portion of the array.
// Inner Loop (j): Compares the key with the elements of the sorted portion (arr[0..i-1]) and shifts larger elements one position to the right to make space for the key.
// Insertion: Once the correct position is found (when arr[j] <= key), the key is placed at arr[j+1].
// Time Complexity:
// Worst Case: O(n²) – Occurs when the array is sorted in reverse order.
// Best Case: O(n) – Occurs when the array is already sorted.
// Average Case: O(n²)

package main

import "fmt"

// InsertionSort function to sort the array
func InsertionSort(arr []int) {
	n := len(arr)
	// Outer loop for each element in the array starting from the second element
	for i := 1; i < n; i++ {
		key := arr[i] // The element to be inserted in the sorted part
		j := i - 1    // Start comparing with the last element of the sorted part

		// Move elements of arr[0..i-1], that are greater than key, to one position ahead
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j] // Shift element to the right
			j = j - 1         // Move to the previous element
		}
		// Insert the key at its correct position
		arr[j+1] = key
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6}
	fmt.Println("Original array:", arr)

	InsertionSort(arr)

	fmt.Println("Sorted array:", arr)
}
