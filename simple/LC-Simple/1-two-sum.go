// Given an array of integers nums and an integer target,
// return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]
			if sum == target {
				result = append(result, i, j)
				return result
			}
		}
	}
	return result
}

func elementSum(mySlice []int, target int) []int {
	myMap := make(map[int]int)

	for i, num := range mySlice {
		complement := target - num
		if idx, found := myMap[complement]; found {
			return []int{idx, i}
		}
		// create a myMap, map[4:0 6:2 9:1, 4:2, 6:3, 9:3]
		myMap[num] = i
	}
	return []int{}
}

// Ex-1: Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
	fmt.Println(elementSum(nums, target))
}
