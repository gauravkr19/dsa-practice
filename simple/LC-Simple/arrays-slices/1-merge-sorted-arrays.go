// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n,
// where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n
// Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]

package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) { // no return required as dealing with []int
	// Initialize three pointers
	p1 := m - 1    // Pointer to the last element in the meaningful part of nums1
	p2 := n - 1    // Pointer to the last element in nums2
	p := m + n - 1 // Pointer to the last element in nums1

	// While there are still elements to compare in nums1 and nums2
	for p1 >= 0 && p2 >= 0 {
		if nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1] // If nums1 element is greater, place it at the end
			p1--
		} else {
			nums1[p] = nums2[p2] // If nums2 element is greater or equal, place it at the end
			p2--
		}
		p-- // Move the pointer for the next element placement
	}

	// If there are remaining elements in nums2, copy them to nums1
	// No need to copy remaining elements from nums1 as they are already in place
	for p2 >= 0 {
		nums1[p] = nums2[p2]
		p--
		p2--
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3

	merge(nums1, m, nums2, n)

	fmt.Println(nums1) // Output: [1 2 2 3 5 6]
}
