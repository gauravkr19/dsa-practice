// You have a long flowerbed in which some of the plots are planted, and some are not.
// However, flowers cannot be planted in adjacent plots.
// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1 means not empty, and an integer n,
// return true if n new flowers can be planted in the flowerbed without violating the no-adjacent-flowers rule and false otherwise.
// Hint: For the first and last elements, we need not check the previous and the next adjacent positions respectively.

package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	for i := 0; i < (len(flowerbed)); i++ {
		if flowerbed[i] == 0 {
			next := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			prev := i == 0 || flowerbed[i-1] == 0 // second part is not evaluated due to (logical OR) short-circuiting.

			if prev && next {
				flowerbed[i] = 1
				count++
			}
		}
	}
	return count >= n
}

func main() {
	flowerbed := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Println(canPlaceFlowers(flowerbed, n))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	fmt.Println(canPlaceFlowers([]int{1, 0, 0, 0, 0, 0, 1}, 2))
}
