package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Logic Breakdown
// Initialize Queues:

// curLevelQ: Tracks nodes at the current level.
// nextLevelQ: Stores nodes for the next level.
// averagePerLevel: Stores the average value for each level.
// Level Order Traversal:

// Iterate while curLevelQ is not empty.
// Initialize curLevelSum to accumulate the sum of node values in the current level.
// Traverse each node in curLevelQ:
// Add its value to curLevelSum.
// Append its left and right children (if present) to nextLevelQ.
// Compute level average and store it.
// Move to the next level by assigning curLevelQ = nextLevelQ.
// End Condition:

// When curLevelQ becomes empty, traversal is complete.
// Return averagePerLevel.

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	curLevelQ := []*TreeNode{root}
	averagePerLevel := []float64{}

	for len(curLevelQ) != 0 {
		nextLevelQ := []*TreeNode{}
		curLevelSum := 0.0

		for _, node := range curLevelQ { // ranging through the slice of nodes
			// update current level summation
			if node != nil {
				curLevelSum += float64(node.Val)
			}
			// Get next level,
			// Add left child
			if node.Left != nil {
				nextLevelQ = append(nextLevelQ, node.Left)
			}
			// Add right child
			if node.Right != nil {
				nextLevelQ = append(nextLevelQ, node.Right)
			}
		}
		curAvg := curLevelSum / float64(len(curLevelQ))
		averagePerLevel = append(averagePerLevel, curAvg)
		curLevelQ = nextLevelQ
	}
	//End of level order traversal
	return averagePerLevel
}
