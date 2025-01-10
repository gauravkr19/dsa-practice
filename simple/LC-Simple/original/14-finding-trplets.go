package main

import (
	"fmt"
	"math"
)

// Given an integer array nums, return true if there exists a triple of indices (i, j, k)
// such that i < j < k and nums[i] < nums[j] < nums[k]. If no such indices exists, return false.

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	min := math.MaxInt32 // Initialize min and mid with largest 32bit integar possible.
	mid := math.MaxInt32

	fmt.Println(min)

	for _, num := range nums {
		if num <= min {
			min = num
		} else if num <= mid {
			mid = num
		} else {
			return true // Found a triplet
		}
	}

	return false
}

func main() {
	// Example usage:
	nums1 := []int{1, 2, 3, 4, 5}
	fmt.Println(increasingTriplet(nums1)) // Output: true

	nums2 := []int{5, 4, 3, 2, 1}
	fmt.Println(increasingTriplet(nums2)) // Output: false

	nums3 := []int{2, 1, 5, 0, 4, 6}
	fmt.Println(increasingTriplet(nums3)) // Output: true
}
