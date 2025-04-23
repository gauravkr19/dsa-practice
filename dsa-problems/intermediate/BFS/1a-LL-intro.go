package main

// ✅ queue := []*TreeNode{root} is correct because it stores pointers to nodes.
// ❌ queue := []TreeNode{root} fails due to type mismatch (TreeNode vs. *TreeNode).
// ❌ queue := []&TreeNode{root} is invalid syntax.
// 	   queue := []TreeNode{*root} // ✅ Works, but copies the value, not the pointer.

import "fmt"

// Node represents a queue node
type Node struct {
	val  int
	next *Node
}

// Queue represents a queue using linked list
type Queue struct {
	front, rear *Node
}

// Enqueue adds an element to the queue
func (q *Queue) Enqueue(val int) {
	newNode := &Node{val: val}
	if q.rear == nil {
		q.front, q.rear = newNode, newNode
		return
	}
	q.rear.next = newNode
	q.rear = newNode
}

// Dequeue removes an element from the queue
func (q *Queue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	val := q.front.val
	q.front = q.front.next
	if q.front == nil { // Queue becomes empty
		q.rear = nil
	}
	return val, nil
}

// BFS using a linked list queue
func BFS(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := &Queue{}
	queue.Enqueue(start)
	visited[start] = true

	for queue.front != nil {
		node, _ := queue.Dequeue()
		fmt.Print(node, " ")

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue.Enqueue(neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 5, 6},
		3: {1},
		4: {1},
		5: {2},
		6: {2},
	}

	fmt.Println("BFS Traversal (Linked List Queue):")
	BFS(graph, 0)
}
