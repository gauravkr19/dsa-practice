// Given an integer array nums sorted in non-decreasing order, remove some duplicates in-place such that each unique element appears at most twice.
// The relative order of the elements should be kept the same.
// Since it is impossible to change the length of the array in some languages, you must instead have the result be placed in the first part of the array nums.
// More formally, if there are k elements after removing the duplicates, then the first k elements of nums should hold the final result.
// It does not matter what you leave beyond the first k elements.
// Return k after placing the final result in the first k slots of nums.
// Do not allocate extra space for another array. You must do this by modifying the input array in-place with O(1) extra memory.Input: nums = [1,1,1,2,2,3]
// Input: nums = [1,1,1,2,2,3]
// Output: 5, nums = [1,1,2,2,3,_]

// [1, 1, 1, 2, 3, 3, 3]
// i	num[i]	num[k-2]	Comparison (!=)	Action			k			num (mutated)
// 2	1		1			false			skip			2			[1, 1, 1, 2, 3, 3, 3]
// 3	2		1			true			num[2] = 2; k++	3			[1, 1, 2, 2, 3, 3, 3]
// 4	3		1			true			num[3] = 3; k++	4			[1, 1, 2, 3, 3, 3, 3]
// 5	3		2			true			num[4] = 3; k++	5			[1, 1, 2, 3, 3, 3, 3]
// 6	3		3			false			skip			5			[1, 1, 2, 3, 3, 3, 3]

package main

import "fmt"

func removeDuplicatesGeneral(nums []int, maxAllowed int) int {
	if len(nums) <= maxAllowed {
		return len(nums)
	}

	k := maxAllowed // First `maxAllowed` elements are always valid

	for i := maxAllowed; i < len(nums); i++ {
		// Compare current num with the num maxAllowed steps back
		if nums[i] != nums[k-maxAllowed] {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3}
	// k := removeDuplicates(nums)
	// fmt.Println(nums[:k])

	k := removeDuplicatesGeneral(nums, 3)
	fmt.Println(nums[:k])
}
