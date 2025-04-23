// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Solve similar problems that involve prefix/suffix ideas, like:
// 1. // Trapping Rain Water
// 2. // Maximum Product Subarray
// Comparing solutions like the two here will help reinforce differences.

// Question: answer[i] = product of all nums[j] where j ≠ i
// Brute force O(n^2)
// for i in 0 to n:
//     product = 1
//     for j in 0 to n:
//         if i != j:
//             product *= nums[j]
//     answer[i] = product
// Instead of removing nums[i] from the total product (which would need division), build the answer from two sides:
// 			From the left, accumulate product of everything before i	(prefix)
//			From the right, accumulate product of everything after i	(suffix)
// Then multiply both.
// answer[i] = product of all elements before i * product of all elements after i

package main

import "fmt"

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// nums = [1,2,3,4],  answer = [1, 1, 2, 6]
	// i=0: prefix=1         => answer[0] = 1
	// i=1: prefix=1*nums[0] => answer[1] = 1
	// i=2: prefix=1*2       => answer[2] = 2
	// i=3: prefix=2*3       => answer[3] = 6

	// Step 1: Calculate prefix product
	prefix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix // store product of all to the left
		prefix *= nums[i]  // update prefix for next index
	}

	// Step 2: Calculate suffix products and multiply with prefix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix // multiply with product of all to the right
		suffix *= nums[i]   // update suffix for next index (backward)
	}
	// i=3: suffix = 1           => answer[3] *= 1        => 6
	// i=2: suffix = 1*4 = 4     => answer[2] *= 4        => 8
	// i=1: suffix = 4*3 = 12    => answer[1] *= 12       => 12
	// i=0: suffix = 12*2 = 24   => answer[0] *= 24       => 24

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
