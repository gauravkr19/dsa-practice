// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
// Example 1:
// Input: prices = [7,1,5,3,6,4]
// Output: 5
// Steps:
// Track the minimum price: As you iterate through the array, keep track of the lowest price seen so far.
// Calculate the profit: For each price, calculate the profit by subtracting the current minimum price from the current price.
// Update the maximum profit: If the calculated profit is higher than the current maximum profit, update the maximum profit.
// Return the result: After iterating through all prices, return the maximum profit.

// Start from i := 1 to ensure you are always buying on an earlier day and selling on a later day.
// If you start from i := 0, you'll incorrectly compare the price on the same day, which doesn't allow you to maximize the profit across different days.

package main

// It's a greedy approach:
// You're greedily updating:
// At each step:
// > 	Check if today's price is a new minimum
// >	Or calculate profit from selling today after buying at minPrice

//  "Buy and Sell Stock II" version (multiple transactions allowed)??
import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i] // Update minPrice if we find a lower price
		} else {
			profit := prices[i] - minPrice // Calculate the profit
			if profit > maxProfit {
				maxProfit = profit // Update maxProfit if the new profit is larger
			}
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

// Example Walkthrough:
// For prices = [7, 1, 5, 3, 6, 4]:

// Start with minPrice = 7, maxProfit = 0.
// On day 2, prices[1] = 1, update minPrice = 1.
// On day 3, prices[2] = 5, profit is 5 - 1 = 4, update maxProfit = 4.
// On day 4, prices[3] = 3, profit is 3 - 1 = 2, no update to maxProfit.
// On day 5, prices[4] = 6, profit is 6 - 1 = 5, update maxProfit = 5.
// On day 6, prices[5] = 4, profit is 4 - 1 = 3, no update to maxProfit.
