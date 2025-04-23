package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root} // Always enqueue left to right
	var result [][]int
	right := true
	front := 0

	for front < len(queue) {
		levelSize := len(queue) - front
		levelElements := list.New() // Deque for zigzag insertion

		for i := 0; i < levelSize; i++ {
			node := queue[front]
			front++

			if right {
				levelElements.PushBack(node.Val) // Insert normally
			} else {
				levelElements.PushFront(node.Val) // Reverse insert
			}

			// Always enqueue left to right
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// Convert deque to slice
		levelSlice := make([]int, levelElements.Len())
		i := 0
		for e := levelElements.Front(); e != nil; e = e.Next() {
			levelSlice[i] = e.Value.(int)
			i++
		}
		result = append(result, levelSlice)
		right = !right
	}
	return result
}

func main() {
	// root = [3,9,20,null,null,15,7]
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(root))
}
