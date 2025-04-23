package main

import (
	"fmt"
)

// Struct for comparing two lists of numbers
type NumComparator struct {
	nums1 []int
	nums2 []int
}

func (nc NumComparator) getElementsOnlyInFirstList(nums1 []int, nums2 []int) []int {
	existsInNums2 := make(map[int]bool)
	onlyInNums1 := make(map[int]bool)

	for _, num := range nums2 {
		existsInNums2[num] = true
	}

	for _, num := range nums1 {
		if !existsInNums2[num] {
			onlyInNums1[num] = true
		}
	}

	result := []int{}
	for num := range onlyInNums1 {
		result = append(result, num)
	}
	return result
}

func (nc NumComparator) findDifference() [][]int {
	return [][]int{nc.getElementsOnlyInFirstList(nc.nums1, nc.nums2), nc.getElementsOnlyInFirstList(nc.nums2, nc.nums1)}
}

func main() {
	comparator := NumComparator{
		nums1: []int{1, 2, 3},
		nums2: []int{2, 4, 6},
	}

	fmt.Println(comparator.findDifference()) // Output: [[1 3] [4 6]]

	comparator2 := NumComparator{nums1: []int{1, 2, 3, 3}, nums2: []int{1, 1, 2, 2}}
	fmt.Println(comparator2.findDifference()) // Output: [[3] []]
}
