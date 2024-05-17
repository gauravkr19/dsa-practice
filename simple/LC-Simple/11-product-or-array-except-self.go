package main

import "fmt"

// func productExceptSelf(nums []int) []int {
// 	n := len(nums)
// 	leftProduct := make([]int, n)
// 	rightProduct := make([]int, n)
// 	answer := make([]int, n)

// 	leftProduct[0] = 1
// 	for i := 1; i < n; i++ {
// 		leftProduct[i] = leftProduct[i-1] * nums[i-1]
// 	}

// 	rightProduct[n-1] = 1
// 	for i := n - 2; i >= 0; i-- {
// 		rightProduct[i] = rightProduct[i+1] * nums[i+1]
// 	}

// 	for i := 0; i < n; i++ {
// 		answer[i] = leftProduct[i] * rightProduct[i]
// 	}

// 	return answer
// }

// func main() {
// 	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
// }

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Step 1: Calculate prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	// Step 2: Calculate suffix products and multiply with prefix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}

	return answer
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums1)) // Output: [24, 12, 8, 6]

	nums2 := []int{2, 3, 4, 5}
	fmt.Println(productExceptSelf(nums2)) // Output: [60, 40, 30, 24]

	nums3 := []int{1, 1, 1, 1}
	fmt.Println(productExceptSelf(nums3)) // Output: [1, 1, 1, 1]

	nums4 := []int{10, 3, 5, 6, 2}
	fmt.Println(productExceptSelf(nums4)) // Output: [180, 600, 360, 300, 900]
}
