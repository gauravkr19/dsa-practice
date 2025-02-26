package main

import (
	"fmt"
)

// Queue using slice
type SliceQueue struct {
	items []int
}

// Enqueue - Adds an element to the queue
func (q *SliceQueue) Enqueue(val int) {
	q.items = append(q.items, val)
}

// Dequeue - Removes and returns the front element
func (q *SliceQueue) Dequeue() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	val := q.items[0]
	q.items = q.items[1:] // Shift elements left
	return val, nil
}

// Front - Returns the front element without removing it
func (q *SliceQueue) Front() (int, error) {
	if len(q.items) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	return q.items[0], nil
}

// IsEmpty - Checks if the queue is empty
func (q *SliceQueue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size - Returns the size of the queue
func (q *SliceQueue) Size() int {
	return len(q.items)
}

// Main function to test the SliceQueue
func main() {
	q := &SliceQueue{}

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	fmt.Println(q.Dequeue()) // 1
	fmt.Println(q.Front())   // 2
	fmt.Println(q.Size())    // 2
}
