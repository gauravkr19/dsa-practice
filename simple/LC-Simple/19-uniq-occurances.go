package main

import "fmt"

func UniqOccurances(nums []int) bool {

	elementOccurances := make(map[int]int)
	for _, num := range nums {
		elementOccurances[num]++
	}

	elementCount := make(map[int]bool)

	for _, count := range elementOccurances {
		if elementCount[count] {
			return false
		}
		elementCount[count] = true
	}

	return true
}

func main() {
	fmt.Println(UniqOccurances([]int{1, 2, 2, 1, 1, 3, 3}))
}
