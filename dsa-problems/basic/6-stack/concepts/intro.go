package main

import (
	"fmt"
)

// Stack structure
type Stack struct {
	items []int
}

// Push - Add element to top of stack
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// Pop - Remove top element from stack
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	top := s.items[len(s.items)-1]     // Get top element
	s.items = s.items[:len(s.items)-1] // Remove top element
	return top
}

// Peek - Get top element without removing it
func (s *Stack) Peek() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.items[len(s.items)-1]
}

// IsEmpty - Check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := Stack{}

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Top element:", stack.Peek())    // 30
	fmt.Println("Popped:", stack.Pop())          // 30
	fmt.Println("Popped:", stack.Pop())          // 20
	fmt.Println("Stack empty?", stack.IsEmpty()) // false
	fmt.Println("Popped:", stack.Pop())          // 10
	fmt.Println("Stack empty?", stack.IsEmpty()) // true
}
