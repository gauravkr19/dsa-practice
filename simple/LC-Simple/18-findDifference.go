package main

import (
	"fmt"
)

type Solution struct{}

// Returns the elements in nums1 that don't exist in nums2.
func (s Solution) getElementsOnlyInFirstList(nums1 []int, nums2 []int) []int {

	existsInNums2 := make(map[int]bool) // convert nums2 slice to map
	onlyInNums1 := make(map[int]bool)   // store only the unique element in nums1

	// convert nums2 slice to map
	for _, num := range nums2 {
		existsInNums2[num] = true
	}

	// check if nums1 element exists in existsInNums2 map and
	// store unique element from nums1 into onlyInNums1 map
	for _, num := range nums1 {
		if !existsInNums2[num] {
			onlyInNums1[num] = true
		}
	}

	// Convert to slice.
	result := []int{}
	for num := range onlyInNums1 {
		result = append(result, num)
	}
	return result
}

func (s Solution) findDifference(nums1 []int, nums2 []int) [][]int {
	return [][]int{s.getElementsOnlyInFirstList(nums1, nums2), s.getElementsOnlyInFirstList(nums2, nums1)}
}

func main() {
	sol := Solution{}
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}
	fmt.Println(sol.findDifference(nums1, nums2)) // Output: [[1 3] [4 6]
	fmt.Println(sol.findDifference([]int{1, 2, 3, 3}, []int{1, 1, 2, 2}))
}

// Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
// answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
// answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
// Note that the integers in the lists may be returned in any order.
// Input: nums1 = [1,2,3], nums2 = [2,4,6]
// Output: [[1,3],[4,6]]

// Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// Output: [[3],[]]
