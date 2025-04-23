// Reoccurance Relation
// T(n, m) = T(n, m/2) + Θ(n) (to find global maximum in a column — (n rows))
// T(n, m) = (Θ(n) + . . . + Θ(n)) *  log m = Θ(n log m)

// Divide and Conquer Approach (Θ(n log m))
// Algorithm
// Select the middle column (m/2).
// Find the maximum value in this column.
// If this maximum is a peak, return it.
// Otherwise:
// 		If the larger neighbor is left, recurse on the left half.
// 		If the larger neighbor is right, recurse on the right half.

// Attempt # 2 - Recursive

// Pick middle column j = m/2
// Find global maximum on column j at (i, j)
// Compare (i, j − 1),(i, j),(i, j + 1)
// Pick left columns if(i, j − 1) > (i, j)
// Similarly for right if (i, j) < (i, j + 1)
// (i, j) is a 2D-peak if neither condition holds
// Solve the new problem with half the number of columns.
// When you have a single column, find global maximum and you‘re done.

package main

import (
	"fmt"
)

// Function to find max element in a column
func findMaxInColumn(grid [][]int, col int) (int, int) {
	maxRow := 0
	for i := 1; i < len(grid); i++ {
		if grid[i][col] > grid[maxRow][col] {
			maxRow = i
		}
	}
	return maxRow, col
}

// Divide and Conquer Peak Finder
func findPeakDivideConquer(grid [][]int, left, right int) (int, int) {
	// Choose middle column
	midCol := (left + right) / 2

	// Find the row with max value in midCol
	maxRow := 0
	for i := 0; i < len(grid); i++ {
		if grid[i][midCol] > grid[maxRow][midCol] {
			maxRow = i
		}
	}

	// Get left and right neighbors
	leftNeighbor := -1
	rightNeighbor := -1
	if midCol > 0 {
		leftNeighbor = grid[maxRow][midCol-1]
	}
	if midCol < len(grid[0])-1 {
		rightNeighbor = grid[maxRow][midCol+1]
	}

	// Check if max in column is a peak, (midCol == 0 ) and (midCol == len(grid[0])-1) act as circuit breakers to avoid out-sof-bound
	// end if left and right neighbor are smaller, peak found!
	if (midCol == 0 || grid[maxRow][midCol] >= leftNeighbor) &&
		(midCol == len(grid[0])-1 || grid[maxRow][midCol] >= rightNeighbor) {
		return maxRow, midCol
	}

	// Move left or right based on the largest neighbor
	if midCol > 0 && (midCol == len(grid[0])-1 || leftNeighbor > rightNeighbor) {
		if midCol-1 >= left { // Out-of-bounds check
			return findPeakDivideConquer(grid, left, midCol-1) // Search left
		}
	}

	if midCol+1 <= right { // Out-of-bounds check
		return findPeakDivideConquer(grid, midCol+1, right) // Search right
	}

	return -1, -1 // Should never reach here
}

func main() {
	grid := [][]int{
		{10, 8, 10, 10},
		{14, 13, 12, 11},
		{15, 9, 11, 21},
		{16, 17, 19, 20},
	}

	peakX, peakY := findPeakDivideConquer(grid, 0, len(grid[0])-1)
	fmt.Printf("Divide & Conquer Peak found at (%d, %d) with value %d\n", peakX, peakY, grid[peakX][peakY])
}
