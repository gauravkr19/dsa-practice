package main

import (
	"fmt"
)

// Definition of TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to build a binary tree from a slice
func BuildTreeFromSlice(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	// Create the root node
	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	// Process nodes level by level
	for i < len(values) {
		current := queue[0] // Get the front node
		queue = queue[1:]   // Dequeue

		// Left Child
		if i < len(values) && values[i] != nil {
			current.Left = &TreeNode{Val: values[i].(int)}
			queue = append(queue, current.Left) // Enqueue left child
		}
		i++

		// Right Child
		if i < len(values) && values[i] != nil {
			current.Right = &TreeNode{Val: values[i].(int)}
			queue = append(queue, current.Right) // Enqueue right child
		}
		i++
	}

	return root
}

// Helper function to print tree in level-order (BFS)
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("Empty Tree")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:] // Dequeue

		if current != nil {
			fmt.Print(current.Val, " ")
			queue = append(queue, current.Left, current.Right) // Enqueue children
		} else {
			fmt.Print("nil ")
		}
	}
	fmt.Println()
}

// Example usage
func main() {
	values := []interface{}{1, 2, 3, nil, 5, nil, 4} // Input array
	root := BuildTreeFromSlice(values)
	PrintTree(root) // Output the constructed tree
}
