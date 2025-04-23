// As we can guess 'a' is a 2D peak if and only if :  a >= b, a >= c, a >= d and a >= e

// Algorithm - Brute force:Greedy Ascent Algorithm - Θ(n²)
// 1. Pick a random (x, y) position in the grid.
// 2. While a neighbor exists with a greater value:
//    a. Move to the neighbor with the highest value.
// 3. Once no better neighbor exists, return (x, y) as a peak.

// So if we try to do the worst case analysis of the algorithm we will find that it would be Θ(nm) where n is the number of rows and m be the number of columns.
// In the case where n = m the worst case complexity would be Θ(n²).

package main

import (
	"fmt"
	"math/rand"
)

// Directions (up, down, left, right)
var directions = [][]int{
	{-1, 0}, // Up
	{1, 0},  // Down
	{0, -1}, // Left
	{0, 1},  // Right
}

// Function to check if a cell is within bounds
func isValid(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

// Greedy Ascent Algorithm, from random element
func findPeakGreedy(grid [][]int, startX, startY int) (int, int) {
	rows, cols := len(grid), len(grid[0]) // to determine out-of-range via isValid func
	x, y := startX, startY

	for {
		bestX, bestY := x, y
		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]
			if isValid(newX, newY, rows, cols) && grid[newX][newY] > grid[bestX][bestY] {
				bestX, bestY = newX, newY
			}
		}

		// If we found a peak, return its coordinates
		if bestX == x && bestY == y {
			return x, y
		}

		// Move to the better neighbor
		x, y = bestX, bestY
	}
}

// Select a random starting point in the grid
func randomStart(rows, cols int) (int, int) {
	return rand.Intn(rows), rand.Intn(cols)
}
func main() {
	grid := [][]int{
		{10, 8, 10, 10},
		{14, 13, 12, 11},
		{15, 9, 11, 21},
		{16, 17, 19, 20},
	}

	// Get random (x, y) without calling rand.Seed()
	startX, startY := randomStart(len(grid), len(grid[0]))

	peakX, peakY := findPeakGreedy(grid, startX, startY)
	fmt.Printf("Greedy Peak found at (%d, %d) with value %d\n", peakX, peakY, grid[peakX][peakY])
}
