package main

import "fmt"

func moveZeroes(nums []int) []int {

	lastNonZeroFoundAt := 0

	// shift non-zero elements to left overwriting the existing slice
	for i, num := range nums {
		if num != 0 {
			nums[lastNonZeroFoundAt] = nums[i]
			lastNonZeroFoundAt++
		}
	}

	// assign 0 to the rest of elements
	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}

	return nums
}

// (Space Optimal, Operation Sub-Optimal)
func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}
