// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]

// Formula Explanation:
// i: This is the index of the current element in the original array nums.
// (i + k) % n: This computes the new index in the result array where the current element from nums[i] should be placed.
// By adding k, we shift the element k steps to the right.
// The % n ensures that the index wraps around when it exceeds the array bounds, i.e., it makes the rotation circular.
// k = k % n => 3 % 7 = 3

// (i+k)%n => 0+3 % 7 = 3		nums[0]
// (i+k)%n => 1+3 % 7 = 4		nums[1]
// (i+k)%n => 2+3 % 7 = 5		nums[2]
// (i+k)%n => 3+3 % 7 = 6		nums[3]
// (i+k)%n => 4+3 % 7 = 0		nums[4]
// (i+k)%n => 5+3 % 7 = 1		nums[5]
// (i+k)%n => 6+3 % 7 = 2		nums[6]
// Input: [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]

package main

import "fmt"

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // Normalize k to be within the array length

	// Create a new array to store the result
	result := make([]int, n)

	// Place elements in their correct positions
	for i := 0; i < n; i++ {
		result[(i+k)%n] = nums[i]
	}

	// Copy the result back into the original array
	fmt.Println("result", result)
	copy(nums, result)
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

// Iteration Steps:
// Iteration 1 (i = 0):

// nums[0] = 1
// New index = (0 + 3) % 7 = 3
// Place 1 at result[3]: result = [_, _, _, 1, _, _, _]
// Iteration 2 (i = 1):

// nums[1] = 2
// New index = (1 + 3) % 7 = 4
// Place 2 at result[4]: result = [_, _, _, 1, 2, _, _]
// Iteration 3 (i = 2):

// nums[2] = 3
// New index = (2 + 3) % 7 = 5
// Place 3 at result[5]: result = [_, _, _, 1, 2, 3, _]
// Iteration 4 (i = 3):

// nums[3] = 4
// New index = (3 + 3) % 7 = 6
// Place 4 at result[6]: result = [_, _, _, 1, 2, 3, 4]
// Iteration 5 (i = 4):

// nums[4] = 5
// New index = (4 + 3) % 7 = 0
// Place 5 at result[0]: result = [5, _, _, 1, 2, 3, 4]
// Iteration 6 (i = 5):

// nums[5] = 6
// New index = (5 + 3) % 7 = 1
// Place 6 at result[1]: result = [5, 6, _, 1, 2, 3, 4]
// Iteration 7 (i = 6):

// nums[6] = 7
// New index = (6 + 3) % 7 = 2
// Place 7 at result[2]: result = [5, 6, 7, 1, 2, 3, 4]
