// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.
// Examples:
// Input: nums = [3,2,3]
// Output: 3
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

package main

import "fmt"

func majorityElement(nums []int) (int, int) {
	candidate := 0
	count := 0

	// First pass: Find the majority element candidate
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}

	// Second pass: Count the occurrences of the candidate (majority element)
	occurrence := 0
	for _, num := range nums {
		if num == candidate {
			occurrence++
		}
	}

	return candidate, occurrence
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	candidate, count := majorityElement(nums)
	fmt.Println(candidate, count)
}
