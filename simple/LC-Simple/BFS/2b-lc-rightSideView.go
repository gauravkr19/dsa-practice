package main

import "fmt"

// Definition for a binary tree node.
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
	queue := []*TreeNode{root} // Using slice for queue

	for len(queue) > 0 {
		levelSize := len(queue) // update queue length for each level
		var rightmost int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			//  (O(n) per operation or O(n^2) due to shifting)
			queue = queue[1:]    // Dequeue front element, so left is rightmost element
			rightmost = node.Val // rightmost is finally updated with the correct value with the last node of the level (or the rightmost node)
			// Even though it may initially hold the leftmost node of the level (2 at level 2), it keeps getting updated and
			// finally gets the correct rightmost node in the last iteration of each level. Before the last iteration, the wrong rightmost val gets eliminated by dequeue.
			// And when we process the last node of the level (3 at level 2, 4 at level 3), rightmost gets the correct value.

			if node.Left != nil {
				queue = append(queue, node.Left) // Enqueue (O(1) amortized)
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // Enqueue (O(1) amortized)
			}
		}
		result = append(result, rightmost) // Capture rightmost node
	}

	return result
}

// Approach									Best for:
// Linked List Queue (container/list)		Large trees, consistent O(N)
// Slice Queue (Original)					Small inputs, but can degrade to O(N²)
// Optimized Slice (Two-pointer technique)	Best alternative to container/list

// **Identifying the Two Pointers in rightSideViewOptimized**
// In the optimized right-side view BFS solution using two pointers, the two pointers are:

// >>front → Marks the beginning of the current level (tracks the first node in the queue for each level).
// >>queue end (len(queue)) → Acts as the dynamically growing end of the queue.
//
// How These Two Pointers Work
// >>front moves forward as nodes are processed.
// >>queue keeps growing as new nodes (children of current level nodes) are appended.
// The difference between len(queue) and front gives the number of nodes in the current level (levelSize).

// Why is this efficient?
// Avoids expensive slice shifting (O(N^2)) by using a logical boundary (levelSize) instead of popping elements.

// Uses 2 pointer approach, Uses two-pointer approach (front pointer for dequeue), making it O(N) time complexity
func rightSideViewOptimized(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var ans []int
	queue := []*TreeNode{root}
	front := 0 // Track front of queue (avoid slice shifting)
	// At the start of a level, 'front' marks the first node of that level.
	// By the end of processing that level, 'front' has moved past all nodes in that level, ensuring correct traversal of the next level.

	for front < len(queue) { // eventually front == len(queue), ending the loop.
		levelSize := len(queue) - front
		// 'levelSize' dynamically adjusts per level. It calculates the number of nodes currently in this level that need to be processed.
		// Since 'front' tracks processed nodes, 'len(queue) - front' ensures we only process this level before moving to the next.

		var rightmost int

		for i := 0; i < levelSize; i++ {
			node := queue[front] // Get the current front node of this level.
			front++              // Move 'front' forward, marking this node as processed.
			rightmost = node.Val // 'rightmost' - gets updated throughout the loop, but the final stored value is the rightmost node
			// Because its the final value retained in rightmost variable before completing the last iteration.

			if node.Left != nil {
				queue = append(queue, node.Left) // Enqueue left child
			}
			if node.Right != nil {
				queue = append(queue, node.Right) // Enqueue right child
			}
		}
		ans = append(ans, rightmost) // Add rightmost node of the level to the result.
	}
	return ans
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
	fmt.Println(rightSideViewOptimized(root)) // Output: [1, 3, 4]
	fmt.Println(rightSideView(root))          // Output: [1, 3, 4]
}
