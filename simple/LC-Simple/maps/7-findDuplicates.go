//Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.
// Example 1:

// Input: nums = [1,2,3,1], k = 3
// Output: true
// Example 2:

// Input: nums = [1,0,1,1], k = 1
// Output: true
// Example 3:

// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false

package main

import "fmt"

func findDuplicate(nums []int, k int) bool {
	dupMap := make(map[int]int)

	for i, num := range nums {
		if j, exist := dupMap[num]; exist && i-j <= k {
			return true
		}

		dupMap[num] = i
	}
	return false
}

func main() {
	fmt.Println(findDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
