// Given n points on a 2D plane where points[i] = [xi, yi], Return the widest vertical area between two points such that no points are inside the area.
// A vertical area is an area of fixed-width extending infinitely along the y-axis (i.e., infinite height).
// The widest vertical area is the one with the maximum width.
// Note that points on the edge of a vertical area are not considered included in the area.

package main

import (
	"fmt"
	"sort"
)

func maxWidthOfVerticalArea(points [][]int) int {
	// Sort the points based on x-coordinates
	sort.Slice(points, func(i, j int) bool { // The sort is not guaranteed to be stable: equal elements may be reversed from their original order. For a stable sort, use SliceStable.
		return points[i][0] < points[j][0]
	})

	// Initialize the maximum width
	maxWidth := 0

	// Iterate through the sorted points and calculate the width
	for i := 1; i < len(points); i++ {
		width := points[i][0] - points[i-1][0]
		if width > maxWidth {
			maxWidth = width
		}
	}

	return maxWidth
}

func main() {
	// Example usage
	points := [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}
	output := maxWidthOfVerticalArea(points)
	fmt.Println(output)
}
