package main

import (
	"fmt"
)

// MergeSort function recursively divides and sorts the array
func MergeSort(arr []int, left, right int) {
	/* 	Question: why not left == right
	However, using left >= right is just a safer defensive check:
	Normally, the recursive calls will always follow left == right at the base case.
	> But in some edge cases (like if the function is misused or incorrect arguments are passed), left > right could occur.
	> left >= right prevents infinite recursion in such cases.
	*/
	if left >= right {
		return // Base case: Single element is already sorted
	}

	mid := left + (right-left)/2 // use mid to split the slice into 2 halves
	// However, if left and right are large (int max value close to 2^31 - 1 for 32-bit systems), their sum can overflow, leading to incorrect behavior.
	// instead use left + (right-left)/2 for mid rather than (left+right)/2

	// Recursively sort left and right halves
	MergeSort(arr, left, mid)
	MergeSort(arr, mid+1, right)

	// Merge the sorted halves
	merge(arr, left, mid, right)
}

// merge function merges two sorted halves of the array
func merge(arr []int, left, mid, right int) {
	// Create temporary slices for left and right subarrays
	leftArr := append([]int{}, arr[left:mid+1]...)     // Copy left half
	rightArr := append([]int{}, arr[mid+1:right+1]...) // Copy right half

	i, j, k := 0, 0, left
	// If we mistakenly set k := 0, we would start merging at arr[0], corrupting previous parts of the array.
	// k = left ensures we only overwrite elements in the correct subarray, preserving other sections of arr.

	// Merge two sorted halves
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			arr[k] = leftArr[i]
			i++
		} else {
			arr[k] = rightArr[j]
			j++
		}
		k++
	}

	// Copy remaining elements of leftArr, if any
	for i < len(leftArr) {
		arr[k] = leftArr[i]
		i++
		k++
	}

	// Copy remaining elements of rightArr, if any
	for j < len(rightArr) {
		arr[k] = rightArr[j]
		j++
		k++
	}
}

func main() {
	arr := []int{38, 27, 43, 3, 9, 82, 10}
	fmt.Println("Original array:", arr)

	MergeSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:  ", arr)
}
