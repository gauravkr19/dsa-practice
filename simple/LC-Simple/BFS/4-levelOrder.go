package main

// Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[9,20],[15,7]]

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{root}
	// front := 0

	for len(queue) > 0 {
		LevelSize := len(queue)
		ElementsInSameLevel := []int{}

		for i := 0; i < LevelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			ElementsInSameLevel = append(ElementsInSameLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, ElementsInSameLevel)

	}
	return result
}

// Optimized
func LevelOrderOptimized(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int // use this declataion instead of result := [][]int{}
	queue := []*TreeNode{root}
	front := 0

	for front < len(queue) {
		levelSize := len(queue) - front // lowercase for variable names, LevelSize → levelSize
		// levelElements := []int{}  // Don't use this initialization. use make([]int,0,levelSize)
		levelElements := make([]int, 0, levelSize) // This avoids unnecessary reallocation overhead with append()

		for i := 0; i < levelSize; i++ {
			node := queue[front]
			front++
			levelElements = append(levelElements, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelElements)

	}
	return result
}

// root = [3,9,20,null,null,15,7]
func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(LevelOrder(root))
	fmt.Println(LevelOrderOptimized(root))
}
