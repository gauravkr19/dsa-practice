package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	var result [][]int
	right := true
	front := 0

	for front < len(queue) {
		levelSize := len(queue) - front
		levelElements := make([]int, levelSize)
		tempQueue := []*TreeNode{} // Temporary slice to avoid duplication

		for i := 0; i < levelSize; i++ {
			node := queue[front]
			front++

			if right {
				levelElements[i] = node.Val
			} else {
				levelElements[levelSize-1-i] = node.Val // Reverse insert
			}

			// Always enqueue left to right
			if node.Left != nil {
				tempQueue = append(tempQueue, node.Left)
			}
			if node.Right != nil {
				tempQueue = append(tempQueue, node.Right)
			}
		}

		// Append new level and update queue
		result = append(result, levelElements)
		queue = append(queue, tempQueue...)
		right = !right
	}
	return result
}

func main() {
	// root = [3,9,20,null,null,15,7]
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(root))
}
