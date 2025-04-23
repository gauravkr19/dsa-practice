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

// nums = [3,2,3]
// 1) num = candidate = 3, count = 1
// 2) num = 2, candidate = 3, count = 0
// 3) num = candidate = 3, count = 1

// Boyer-Moore Majority Vote Algorithm,  O(n) time and O(1) space.
func majorityElement(nums []int) (int, int) {
	candidate := 0 // single candidate
	count := 0     // vote counter

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

// t’s O(n) time and O(n) space.
func majorityElementMap(nums []int) int {
	counts := make(map[int]int)
	half := len(nums) / 2
	for _, num := range nums {
		counts[num]++
		if counts[num] > half {
			return num
		}
	}
	// Given the problem constraints, majority always exists.
	return -1 // Unreachable under normal conditions.
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	candidate, count := majorityElement(nums)
	fmt.Println(candidate, count)

	fmt.Println(majorityElementMap(nums))
}
