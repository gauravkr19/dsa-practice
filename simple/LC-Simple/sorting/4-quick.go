// When the first element of the array is selected as the pivot, the partitioning logic and behavior of the algorithm change slightly, but the overall structure remains the same.
// The pivot will still divide the array into two parts: elements smaller than the pivot on the left and elements greater than the pivot on the right.
// The pivot is then recursively applied to these sub-arrays.
// Explanation:
// Pivot Selection: Instead of choosing the last element as the pivot, we select the first element of the array (arr[low]).
// Partitioning: We will rearrange the array so that all elements less than the pivot are placed on its left and all elements greater than the pivot are placed on its right.
// Recursive Sorting: The sub-arrays (elements to the left and right of the pivot) are sorted recursively.
// Explanation of the Code:
// Partition Function:
// Pivot: The first element (arr[low]) is chosen as the pivot.
// i: i starts from low + 1 (the second element) and keeps track of elements that are greater than the pivot.
// j: j starts from high and moves leftward, tracking elements smaller than the pivot.
// Swapping: Elements that are on the wrong side of the pivot are swapped (i.e., an element greater than the pivot on the left side and an element smaller than the pivot on the right side are swapped).
// Once the elements have been rearranged, the pivot is swapped into its correct position (arr[j]), and the index j is returned.
// QuickSort Function:
// The QuickSort function calls partition to partition the array and then recursively applies the sorting algorithm to the sub-arrays to the left and right of the pivot.

package main

import "fmt"

// Partition function that takes the first element as the pivot
func partition(arr []int, low, high int) int {
	pivot := arr[low] // Choose the first element as the pivot
	i := low + 1      // Index for elements greater than pivot
	j := high         // Index for elements less than pivot

	for i <= j {
		// Find the first element greater than the pivot from the left side
		for i <= high && arr[i] <= pivot {
			i++
		}
		// Find the first element smaller than the pivot from the right side
		for arr[j] > pivot {
			j--
		}
		// If there are elements on the wrong side, swap them
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Place the pivot at the correct position
	arr[low], arr[j] = arr[j], arr[low]
	return j
}

// QuickSort function
func QuickSort(arr []int, low, high int) {
	if low < high {
		// Partition the array
		pi := partition(arr, low, high)

		// Recursively sort elements before and after partition
		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 80, 30, 90, 40, 50, 70}
	fmt.Println("Original array:", arr)

	QuickSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:", arr)
}
