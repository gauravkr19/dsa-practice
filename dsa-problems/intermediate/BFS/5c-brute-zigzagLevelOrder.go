// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right,
// then right to left for the next level and alternate between)

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
	front := 0
	var right bool
	right = true

	for front < len(queue) {
		levelSize := len(queue) - front
		// reset levelElements slice
		levelElements := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[front]
			front++
			levelElements = append(levelElements, node.Val)

			if !right {
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
			} else {
				if node.Right != nil {
					queue = append(queue, node.Right)
				}
				if node.Left != nil {
					queue = append(queue, node.Left)
				}
			}
		}
		right = !right
		result = append(result, levelElements)
	}
	return result
}

func main() {
	// root = [3,9,20,null,null,15,7]
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(root))
}
