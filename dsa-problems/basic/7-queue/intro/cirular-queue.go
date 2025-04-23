// A circular queue is a 'fixed-size' queue that connects the end back to the beginning, forming a circle.
// It avoids shifting elements like in a slice-based queue and efficiently utilizes available space.
// Insertion is via Front, it tracks insertion. Deletion is via rear
// special cases:
// Insert - IsFull
// Deletion - IsEmpty + when only 1 element exists, to set front, rear t0 -1
// Peek - IsEmpty
package main

import "fmt"

// CircularQueue struct
type CircularQueue struct {
	data     []int
	front    int
	rear     int
	capacity int
	size     int
}

// Initialize queue
func NewCircularQueue(cap int) *CircularQueue {
	return &CircularQueue{
		data:     make([]int, cap),
		front:    -1,
		rear:     -1,
		capacity: cap,
		size:     0,
	}
}

// Enqueue operation
func (q *CircularQueue) Enqueue(val int) bool {
	if q.IsFull() {
		fmt.Println("Queue is full!")
		return false
	}
	if q.IsEmpty() {
		q.front = 0 // required for dequeue to know that element exists
	}
	q.rear = (q.rear + 1) % q.capacity
	q.data[q.rear] = val
	q.size++
	return true
}

// Dequeue operation
func (q *CircularQueue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		fmt.Println("Queue is empty!")
		return 0, false
	}
	val := q.data[q.front]
	if q.front == q.rear { // Only one element left
		q.front, q.rear = -1, -1
	} else {
		q.front = (q.front + 1) % q.capacity
	}
	q.size--
	return val, true
}

// Peek (Get Front Element)
func (q *CircularQueue) Peek() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}
	return q.data[q.front], true
}

// IsEmpty check
func (q *CircularQueue) IsEmpty() bool {
	return q.size == 0
	// or q.front == -1
}

// IsFull check
func (q *CircularQueue) IsFull() bool {
	return q.size == q.capacity
	// or return q.size == q.capacity
}

// GetSize
func (q *CircularQueue) GetSize() int {
	return q.size
}

// Main function
func main() {
	q := NewCircularQueue(5)

	q.Enqueue(10)
	q.Enqueue(20)
	q.Enqueue(30)
	q.Enqueue(40)
	q.Enqueue(50)
	fmt.Println("Queue full:", q.IsFull()) // True

	val, _ := q.Dequeue()
	fmt.Println("Dequeued:", val) // 10

	q.Enqueue(60)
	front, _ := q.Peek()                    // 20
	fmt.Println("Peek - Front Val:", front) // 10
	fmt.Println("Queue size:", q.GetSize()) // 5 (still full)
}
