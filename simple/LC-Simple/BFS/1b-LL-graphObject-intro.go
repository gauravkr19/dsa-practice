package main

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

// use BFS a method of Graph
type Graph struct {
	adjacencyList map[int][]int
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

func (g *Graph) BFS(start int) {
	visited := make(map[int]bool)
	queue := &Queue{}
	queue.Enqueue(start)
	visited[start] = true

	for queue.front != nil {
		node, _ := queue.Dequeue()
		fmt.Println(node)
		for _, neighbor := range g.adjacencyList[node] {
			if !visited[neighbor] {
				queue.Enqueue(neighbor)
				visited[neighbor] = true
			}
		}
	}
}

func main() {
	g := Graph{adjacencyList: map[int][]int{
		0: {1, 2},
		1: {2, 3},
		2: {4},
		3: {},
		4: {},
	}}
	g.BFS(0)
}
