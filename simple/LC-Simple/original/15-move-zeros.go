package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)

	// Count the zeroes
	numZeroes := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			numZeroes++
		}
	}

	// Make all the non-zero elements retain their original order.
	ans := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			ans = append(ans, nums[i])
		}
	}

	// Move all zeroes to the end
	for ; numZeroes > 0; numZeroes-- {
		ans = append(ans, 0)
	}

	// Copy the result back to the original array
	for i := 0; i < n; i++ {
		nums[i] = ans[i]
	}
}

// (Space Sub-Optimal)
// Slice - Automatically passes the address and does not need to return as it can see through via addr
func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) // Output: [1 3 12 0 0]
}
