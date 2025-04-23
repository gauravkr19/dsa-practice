// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
// Consider the number of elements in nums which are not equal to val be k, to get accepted, you need to do the following things:
// Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.
// Input: nums = [3,2,2,3], val = 3
// Output: 2, nums = [2,2,_,_]
// Explanation: Your function should return k = 2, with the first two elements of nums being 2.
// It does not matter what you leave beyond the returned k (hence they are underscores).

// **General Pattern to Remember (Template)**
// >	Initialize a write pointer (k) to the first valid write position (often 0 or 1)
// >	Loop with a read pointer (i) to go through the array
// >	At each i, check if current element is valid (unique, or not to be removed, or allowed duplicate)
// >	If yes, write to nums[k] and increment k

package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	k := 0 // Pointer for the position of the next non-val element

	// Iterate through the array
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			// Place the non-val element at index k, k also tracks non val count
			nums[k] = nums[i]
			k++ // Increment k for the next position
		}
	}

	// k is the count of elements that are not equal to val
	return k
}

func main() {
	nums := []int{3, 2, 2, 3, 3}
	val := 3

	k := removeElement(nums, val)

	// Output the result
	fmt.Println("The first", k, "elements after removing", val, "are:", nums[:k])
	fmt.Println("Length of the array after removal:", k)
}
