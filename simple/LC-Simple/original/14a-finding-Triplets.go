package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) (bool, []int, []int) {
	if len(nums) < 3 {
		return false, nil, nil
	}

	min := math.MaxInt32 // Initialize min with the largest 32-bit integer
	mid := math.MaxInt32 // Initialize mid with the largest 32-bit integer

	minIndex, midIndex := -1, -1 // Store the indices of min and mid

	for i, num := range nums {
		if num <= min {
			// Update min and its index
			min = num
			minIndex = i
		} else if num <= mid {
			// Update mid and its index
			mid = num
			midIndex = i
		} else {
			// Return true along with the indices and the corresponding values
			return true, []int{minIndex, midIndex, i}, []int{min, mid, num}
		}
	}

	// If no triplet is found, return false
	return false, nil, nil
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	found, indices, values := increasingTriplet(nums)

	if found {
		fmt.Println("Triplet found at indices:", indices, "with values:", values)
	} else {
		fmt.Println("No increasing triplet found")
	}
}
