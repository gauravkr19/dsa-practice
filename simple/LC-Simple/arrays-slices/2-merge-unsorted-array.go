// make([]int, n)... (The ... operator):

// The ... operator is used to unpack or expand the elements of the slice created by make([]int, n) into individual elements.
// Without the ..., append(nums1, make([]int, n)) would append the entire slice as a single element (i.e., a nested slice).
// With the ..., the elements of the slice [0, 0, 0] are unpacked and appended individually to nums1. So, if nums1 was [1, 2, 3], after appending, it would become [1, 2, 3, 0, 0, 0].

package main

import (
	"fmt"
	"sort"
)

func mergeUnsorted(nums1 []int, nums2 []int) []int {
	// Get the length of nums1 and nums2
	m := len(nums1)
	n := len(nums2)

	// Extend nums1 to accommodate elements from nums2
	nums1 = append(nums1, make([]int, n)...) // Using the variadic operator to unpack the slice elements

	// Copy elements from nums2 into nums1
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}

	// Sort the combined nums1
	sort.Ints(nums1)

	// Return the merged and sorted nums1
	return nums1
}

func main() {
	nums1 := []int{3, 1, 2} // Unsorted nums1
	nums2 := []int{6, 5, 4} // Unsorted nums2

	// Call the mergeUnsorted function
	result := mergeUnsorted(nums1, nums2)

	// Output the result
	fmt.Println(result) // Output: [1 2 3 4 5 6]
}
