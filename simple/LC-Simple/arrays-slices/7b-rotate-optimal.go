// For the example nums = [1,2,3,4,5,6,7] and k = 3, the steps are:
// Reverse the entire array: [7,6,5,4,3,2,1]
// Reverse the first 3 elements: [5,6,7,4,3,2,1]
// Reverse the remaining 4 elements: [5,6,7,1,2,3,4]

package main

import "fmt"

// Total work done = O(n) + O(k) + O(n-k) = O(2n) = O(n)
// It uses a fixed number of variables (start, end, temporary swap variable), so auxiliary space = O(1)
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n // Normalize k to be within the array length

	reverse(nums, 0, n-1) // Reverse the entire array → O(n)
	reverse(nums, 0, k-1) // Reverse the first k elements → O(k)
	reverse(nums, k, n-1) // Reverse the remaining elements → O(n-k)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
