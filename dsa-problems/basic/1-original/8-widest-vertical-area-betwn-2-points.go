// Given n points on a 2D plane where points[i] = [xi, yi], Return the widest vertical area between two points such that no points are inside the area.
// A vertical area is an area of fixed-width extending infinitely along the y-axis (i.e., infinite height).
// The widest vertical area is the one with the maximum width.
// Note that points on the edge of a vertical area are not considered included in the area.
// Imagine you drop vertical walls from the sky at each x-point.
// Then ask: Between which two walls is the widest open space (no walls between them)?
package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

// Same complexity: O(n log n) sort + O(n) max gap scan.
func maxWidthOfVerticalAreaImprovedFunc(points [][]int) int {
	// Sort the points based on the x-coordinate using slices.SortFunc
	// slices.SortFunc: Handles sorting more ergonomically by operating on elements directly instead of indices.

	// 	This is a comparison-based sort. Internally, Go's standard library will repeatedly call your comparison function to decide the order of elements during sorting.
	// That means: it will repeatedly call func(a, b []int) int with different pairs of elements from points, and sort based on the return value:

	slices.SortFunc(points, func(a, b []int) int { // based on New go lib 1.21,  a = {8,7}, b = {9,9}
		return cmp.Compare(a[0], b[0]) // Only x (first element of each inner slice) is used to sort. Y doesn't matter at all in this logic.
	})
	// After cmp.Compare, the input becomes: points = [][]int{{7,4},{8,7},{9,9},{9,7}}

	maxWidth := 0
	for i := 1; i < len(points); i++ {
		width := points[i][0] - points[i-1][0]
		if width > maxWidth {
			maxWidth = width
		}
	}
	return maxWidth
}

// using old func
func maxWidthOfVerticalArea(points [][]int) int {
	// Sort the points based on x-coordinates

	// sort.Ints() or slices.SortFunc() both use QuickSort (hybrid) internally, So sorting takes O(n log n)
	sort.Slice(points, func(i, j int) bool { // based on Old go lib 1.8
		// The sort is not guaranteed to be stable: equal elements may be reversed from their original order. For a stable sort, use SliceStable
		return points[i][0] < points[j][0]
	})

	// Initialize the maximum width
	maxWidth := 0

	// Iterate through the sorted points and calculate the width, Just one pass over the sorted data → O(n)
	for i := 1; i < len(points); i++ {
		width := points[i][0] - points[i-1][0]
		if width > maxWidth {
			maxWidth = width
		}
	}

	return maxWidth
}

func maxWidthOfVerticalAreaNew(points [][]int) int {
	// Step 1: Extract all x-coordinates, O(N) to extract x-coordinates
	xCoords := make([]int, len(points))
	for i, point := range points {
		xCoords[i] = point[0]
	}

	// Step 2: Sort the x-coordinates, O(N log N) to sort x-coordinates
	// by sorting in ascending order, it makes sure no other point lies between consecutive elements.
	sort.Ints(xCoords)

	// Step 3: Compute the maximum difference between consecutive x values, O(N) to scan for max gap
	maxGap := 0
	for i := 1; i < len(xCoords); i++ {
		gap := xCoords[i] - xCoords[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}

func main() {
	// Example usage
	points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	output := maxWidthOfVerticalAreaImprovedFunc(points) //  Total = O(n log n)
	fmt.Println(output)

	output = maxWidthOfVerticalArea(points) //  Total = O(n log n)
	fmt.Println(output)

	output = maxWidthOfVerticalAreaNew(points) // /  Total: O(N log N)
	fmt.Println(output)
}
