package main

import (
	"fmt"
)

// Function to get elements only in the first list using nested loops
func getElementsOnlyInFirstList(nums1 []int, nums2 []int) []int {
	onlyInNums1 := []int{}
	// Outer loop for each element in nums1
	for i := 0; i < len(nums1); i++ {
		found := false
		// Inner loop to check if nums1[i] exists in nums2
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				found = true
				break
			}
		}
		// If not found in nums2, add to the result list
		if !found {
			// Check for duplicates in the result list
			duplicate := false
			for k := 0; k < len(onlyInNums1); k++ {
				if nums1[i] == onlyInNums1[k] {
					duplicate = true
					break
				}
			}
			if !duplicate {
				onlyInNums1 = append(onlyInNums1, nums1[i])
			}
		}
	}
	return onlyInNums1
}

// Function to find the difference between two lists
func findDifference(nums1 []int, nums2 []int) [][]int {
	return [][]int{getElementsOnlyInFirstList(nums1, nums2), getElementsOnlyInFirstList(nums2, nums1)}
}

func main() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}
	fmt.Println(findDifference(nums1, nums2)) // Output: [[3], []]
}
