package main

import (
	"fmt"
)

// There are n kids with candies. You are given an integer array candies, where each candies[i] represents the number of candies the ith kid has,
// and an integer extraCandies, denoting the number of extra candies that you have.

// Return a boolean array result of length n, where result[i] is true if, after giving the ith kid all the extraCandies,
// they will have the greatest number of candies among all the kids, or false otherwise.

func kidsWithMaxmCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	for _, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}
	}

	// var result []bool - use append with this decalaration of result to incr slice size
	result := make([]bool, len(candies))
	for i, v := range candies {
		result[i] = (v+extraCandies >= maxCandies)
		// result = append(result, v+extraCandies >= maxCandies)  //  this works with nil slice initialization
	}
	return result
}

func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	result := kidsWithMaxmCandies(candies, extraCandies)
	fmt.Println(result)
	fmt.Println(kidsWithMaxmCandies([]int{4, 6, 2, 3}, 4))
}
