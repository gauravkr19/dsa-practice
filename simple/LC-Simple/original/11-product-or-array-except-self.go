// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums
//  except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Solve similar problems that involve prefix/suffix ideas, like:
// 1. // Trapping Rain Water
// 2. // Maximum Product Subarray
// Comparing solutions like the two here will help reinforce differences.

package main

import "fmt"

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// nums = [1,2,3,4]
	// Step 1: Calculate prefix products , 	// Answer with Prefix: answer= [1 1 2 6]
	// answer[0] = 				Prefix = 1
	// Prefix = Prefix * nums[0] = 1*1 = 1
	// answer[1] = Prefix = 1
	// Prefix = Prefix * nums[1] = 1*2 = 2
	// answer[2] = Prefix = 2
	// Prefix = Prefix * nums[2] = 2*3 = 6
	// answer[3] = Prefix = 6
	// Prefix = Prefix * nums[3] = 6*4 = 24

	// Step 1: Calculate prefix product
	prefix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix
		prefix *= nums[i] // prefix = prefix * num[i]
	}

	// Step 2: Calculate suffix products and multiply with prefix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix // answer[i] = answer[i] * suffix
		suffix *= nums[i]   // suffix = suffix * num[i]
	}
	// Input = nums= []int{1, 2, 3, 4}
	// Answer with Prefix: answer= [1 1 2 6]
	// Now: Suffix part
	// answer[3] = answer[3] * suffix = 	6*1=6
	// suffix = suffix * nums[3] = 1*4=4
	// answer[2] = answer[2] * suffix = 	2*4=8
	// suffix = suffix * nums[2] = 4*3=12
	// answer[1] = answer[1] * suffix = 	1*12 = 12
	// suffix = suffix * nums[1] = 	12*2=24
	// answer[0] = answer[0] * suffix = 	1*24 = 24
	// suffix = suffix * nums[0] = 24*1=24
	return answer
}

func productExceptSelf1(nums []int) []int {
	n := len(nums)
	leftProduct := make([]int, n)
	rightProduct := make([]int, n)
	answer := make([]int, n)

	leftProduct[0] = 1
	// At index  i, leftProduct[i] stores the product of all elements before index
	for i := 1; i < n; i++ {
		leftProduct[i] = leftProduct[i-1] * nums[i-1]
	}

	rightProduct[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rightProduct[i] = rightProduct[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		answer[i] = leftProduct[i] * rightProduct[i]
	}

	fmt.Println("DONE")
	return answer
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf1(nums1)) // Output: [24, 12, 8, 6]

	nums2 := []int{2, 3, 4, 5}
	fmt.Println(productExceptSelf1(nums2)) // Output: [60, 40, 30, 24]

	// nums3 := []int{1, 1, 1, 1}
	// fmt.Println(productExceptSelf1(nums3)) // Output: [1, 1, 1, 1]

	nums4 := []int{10, 3, 5, 6, 2}
	fmt.Println(productExceptSelf2(nums4)) // Output: [180, 600, 360, 300, 900]
}
