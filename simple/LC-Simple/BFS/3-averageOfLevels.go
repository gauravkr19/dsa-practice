// // Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.
// Input: root = [3,9,20,null,null,15,7]
// Output: [3.00000,14.50000,11.00000]
// Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
// Hence return [3, 14.5, 11].
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := []float64{}
	// front := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0 // initialize sum to 0 before starting processing each level

		for i := 0; i < levelSize; i++ {
			node := queue[0]  // Get the top node
			queue = queue[1:] // Dequeue
			sum += node.Val   // process the node

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(levelSize))
	}
	return result
}

// We can solve this using a naïve BFS with slice shifting (O(N^2)), but that’s inefficient.
// Instead, we use a queue with a two-pointer approach, which keeps the time complexity at O(N) while reducing extra space. Let me implement that.
// optimized using 2 pointer to avoid dequeue operation and hence avoidging O(n^2) slice shitfing each elements
func averageOfLevelsOptimized(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}

	var result []float64
	queue := []*TreeNode{root}
	front := 0 // Track front without slice shifting

	for front < len(queue) {
		levelSize := len(queue) - front
		sum := 0

		for i := 0; i < levelSize; i++ {
			node := queue[front]
			front++
			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, float64(sum)/float64(levelSize))
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
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(averageOfLevels(root)) // Output: [1 2.5 4.5]
	/*
			 3
			/ \
		   9  20
			 /  \
			15   7
	*/
	root = &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(averageOfLevelsOptimized(root)) // Output: [[3 14.5 11]
}
