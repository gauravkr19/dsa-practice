package main

import "fmt"

// nums = []int{1, 0, 3, 0, 12}
// i=0, num[0] = 1, Non-Zero element so swap with insertPos = 0. [1, 0, 3, 0, 12], insertPos++ = 1
// i=1, num[1] = 0, Zero element so skip [1, 0, 3, 0, 12]
// i=2, num[2] = 1, Non-Zero element so swap with insertPos = 1. [1, 3, 0, 0, 12], insertPos++ = 2
// i=3, num[3] = 0, Zero element so skip [1, 3, 0, 0, 12]
// i=4, num[4] = 12, Non-Zero element so swap with insertPos = 2. insertPos++ = 3

func moveZeroes(nums []int) []int {
	insertPos := 0 // We use a pointer insertPos to keep track of where the next non-zero element should go.

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// Swap non-zero element with the element at insertPos
			nums[insertPos], nums[i] = nums[i], nums[insertPos]
			insertPos++
		}
	}
	return nums
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) // Output: [1 3 12 0 0]

	moveZeroesBrute(nums)
	fmt.Println(nums)
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

func moveZeroesBrute(nums []int) []int {

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
