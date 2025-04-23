// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once.
// The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
// Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
// Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Input: nums = [1,1,2]
// Output: 2, nums = [1,2,_]
package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1 // Pointer for the next unique element position

	// Iterate through the array starting from the second element
	for i := 1; i < len(nums); i++ {
		// If the current element is not equal to the previous one, it's unique
		if nums[i] != nums[i-1] {
			nums[k] = nums[i] // Place it at the k-th position
			k++               // Increment k for the next unique position
		}
	}

	// k is the count of unique elements
	return k
}

func main() {
	nums := []int{1, 1, 2, 3, 3, 4, 4, 5}

	k := removeDuplicates(nums)

	// Output the result
	fmt.Println("The first", k, "unique elements are:", nums[:k])
	fmt.Println("Length of the array after removing duplicates:", k)
}

// Code Output:
// The first 5 unique elements are: [1 2 3 4 5]
// Length of the array after removing duplicates: 5
