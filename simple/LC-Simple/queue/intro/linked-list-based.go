// For queue implementations, pointer receivers (*Queue) are better since we modify front, rear,
// and size but don’t need to change the queue object itself. For linked lists, double pointers (**Node)** are needed when functions must reassign the head pointer.
package main

import (
	"fmt"
)

// Node structure
type Node struct {
	data int
	next *Node
}

// Queue using Linked List
type LinkedListQueue struct {
	front *Node
	rear  *Node
	size  int // tracks the count of elements
}

// Enqueue - Adds an element at the rear
func (q *LinkedListQueue) Enqueue(val int) {
	newNode := &Node{data: val}
	if q.rear == nil {
		q.front, q.rear = newNode, newNode
	} else {
		q.rear.next = newNode
		q.rear = newNode
	}
	q.size++
}

// Dequeue - Removes and returns the front element
func (q *LinkedListQueue) Dequeue() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	val := q.front.data
	q.front = q.front.next
	if q.front == nil {
		q.rear = nil // Queue is now empty
	}
	q.size--
	return val, nil
}

// Front - Returns the front element without removing it
func (q *LinkedListQueue) Front() (int, error) {
	if q.front == nil {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.front.data, nil
}

// IsEmpty - Checks if the queue is empty
func (q *LinkedListQueue) IsEmpty() bool {
	return q.front == nil
}

// Size - Returns the size of the queue
func (q *LinkedListQueue) Size() int {
	return q.size
}

// Main function to test LinkedListQueue
func main() {
	q := &LinkedListQueue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue()) // 1
	fmt.Println(q.Front())   // 2
	fmt.Println(q.Size())    // 2
}
