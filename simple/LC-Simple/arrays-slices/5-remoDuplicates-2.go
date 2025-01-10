// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice.
// The relative order of the elements should be kept the same.
// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums.
// More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result.
// It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.Input: nums = [1,1,1,2,2,3]
// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]

// [1,1,1,2,2,3]
// k = 1, i = 2, 1 != 1 || 1 != 1 false
// k = 1, i = 3, 2 != 1 || 2 != 1 true => k++, k = 2,  [1,1,2,2,2,3]
// k = 2, i = 4, 2 != 1 || 2 != 2, true
// k = 3   [1,1,2,2,2,3]
// k = 3, i = 5, 5 != 2
// k = 4  [1,1,2,2,3,3]

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1 // Start at the second element

	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[k-1] || nums[i] != nums[k] {
			k++
			nums[k] = nums[i]
		}
	}

	return k + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k := removeDuplicates(nums)
	fmt.Println(nums[:k])
}
