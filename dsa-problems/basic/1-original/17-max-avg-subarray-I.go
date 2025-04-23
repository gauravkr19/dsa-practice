// You are given an integer array nums consisting of n elements, and an integer k.
// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value.
// Any answer with a calculation error less than 10-5 will be accepted.

// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000

// sliding window example

package main

import (
	"fmt"
)

func findMaxAverage(nums []int, k int) float64 {
	sum := 0

	// get the sum of first k-1 (i < k) elements, sliding window
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)
}

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // Output: 12.75
}
