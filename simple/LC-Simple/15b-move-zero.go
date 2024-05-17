package main

import "fmt"

func moveZeroes(nums []int) []int {
	lastNonZeroFoundAt := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// Swap non-zero element with the element at lastNonZeroFoundAt
			nums[lastNonZeroFoundAt], nums[i] = nums[i], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) // Output: [1 3 12 0 0]
	fmt.Println(moveZeroes([]int{2, 5, 0, 2, 5, 0}))
}

// Returning slices
// =====================
// In Go, when you declare a function with a return type, it is expected to explicitly return a value of that type.
// If you don't specify a return type in the function signature, it means the function doesn't return any value explicitly,
// although it can still modify data that's passed to it as reference (like slices or maps).
// In the case of your moveZeroes function, it's modifying the slice nums directly and doesn't need to return anything.
// But when you call it like fmt.Println(moveZeroes([]int{2, 5, 0, 2, 5, 0})), the Println function expects a value to print,
// and it receives nothing because moveZeroes doesn't return anything.
// This leads to a mismatch between the expected argument type and the actual argument type for Println, resulting in an error.
// To fix this, you can simply call moveZeroes without trying to print its return value, like this:
