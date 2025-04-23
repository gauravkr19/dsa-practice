// Given the root of a binary tree, imagine yourself standing on the right side of it,
// return the values of the nodes you can see ordered from top to bottom.

// Example 1:
// Input: root = [1,2,3,null,5,null,4]
// Output: [1,3,4]

// Intuition
// The goal is to create a view of a binary tree from the right side, which means capturing the rightmost node at each level of the tree.
// This suggests using a level-order traversal (breadth-first search) approach, where we ensure that for each level, the last node encountered in the traversal is recorded.

// Approach
// The rightSideView function uses a queue to perform a level-order traversal of the tree. The process involves:

// > Enqueuing the root node initially.
// > Iterating over each level of the tree. For each level:
// 		Track the number of nodes (nodesInCurrentLevel).
// 		Traverse through these nodes, updating a variable prev with the value of the current node.
// 		After finishing each level, the value in prev will be the rightmost node of that level, which is then appended to the ans slice.
// > The loop continues until all levels are processed and the queue is empty.

// Steps
// 1. Return nil if the root is nil, as there's no tree to view.
// 2. Initialize a slice ans to store the right side view.
// 3. Use a queue to perform a level-order traversal, starting with the root.
// 4. While the queue is not empty:
// 		Determine the number of nodes at the current level (nodesInCurrentLevel).
// 		Iterate through these nodes:
// 			Dequeue each node and update prev with its value.
// 			Enqueue the left and right children of the dequeued node (if they exist).
// 		After completing the level, append prev to ans.
// Return ans, which contains the right side view of the tree.

// Complexity
// Time complexity: O(n), where n is the number of nodes in the tree. Each node is processed exactly once.
// Space complexity: O(m), where m is the maximum number of nodes at any level of the tree. This space is required for the queue used in the level-order traversal.

// rightSideView returns the right-side view of the binary tree.
package main

import (
	"container/list" // using this package which makes dequeue O(n) instead of O(n^2) - Best Solution but will need an extra package
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	queue := list.New() // Using linked list for queue
	queue.PushBack(root)

	for queue.Len() > 0 {
		levelSize := queue.Len()
		var rightmost int

		for i := 0; i < levelSize; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) // Dequeue O(1)
			rightmost = node.Val

			if node.Left != nil {
				queue.PushBack(node.Left) // Enqueue O(1)
			}
			if node.Right != nil {
				queue.PushBack(node.Right) // Enqueue O(1)
			}
		}
		result = append(result, rightmost) // Capture rightmost node
	}

	return result
}

func main() {
	// Input: root = [1,2,3,null,5,null,4]
	/*
	     1
	    / \
	   2   3
	    \   \
	     5   4
	*/
	root := &TreeNode{1,
		&TreeNode{2, nil, &TreeNode{5, nil, nil}},
		&TreeNode{3, nil, &TreeNode{4, nil, nil}},
	}
	fmt.Println(rightSideView(root)) // Output: [1, 3, 4]
}
