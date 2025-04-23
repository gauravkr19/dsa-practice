// Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
// answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
// answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
// Note that the integers in the lists may be returned in any order.
// Input: nums1 = [1,2,3], nums2 = [2,4,6]
// Output: [[1,3],[4,6]]

// Input: nums1 = [1,2,3,3], nums2 = [1,1,2,2]
// Output: [[3],[]]

package main

import "fmt"

func findDisjoint(nums1 []int, nums2 []int) [][]int {
	// Function to remove duplicates from a slice
	removeDuplicates := func(nums []int) []int {
		seen := make(map[int]bool)
		result := []int{}
		for _, num := range nums {
			if !seen[num] {
				result = append(result, num)
				seen[num] = true
			}
		}
		return result
	}

	// Remove duplicates from nums1 and nums2
	nums1Unique := removeDuplicates(nums1)
	nums2Unique := removeDuplicates(nums2)

	// Create maps to store the occurrence of elements in nums1 and nums2
	freq1 := make(map[int]bool)
	freq2 := make(map[int]bool)

	// Populate maps with unique elements
	for _, num := range nums1Unique {
		freq1[num] = true
	}
	for _, num := range nums2Unique {
		freq2[num] = true
	}

	// Initialize lists to store distinct integers not present in the other array
	notInNums2 := []int{}
	notInNums1 := []int{}

	// Iterate over nums1 and nums2 to find distinct integers not present in the other array
	for _, num := range nums1Unique {
		if _, found := freq2[num]; !found {
			notInNums2 = append(notInNums2, num)
		}
	}

	for _, num := range nums2Unique {
		if _, found := freq1[num]; !found {
			notInNums1 = append(notInNums1, num)
		}
	}

	return [][]int{notInNums2, notInNums1}
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	fmt.Println(findDisjoint(nums1, nums2)) // Output: [[3], []]
}
