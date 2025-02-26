package main

import "fmt"

// Stack struct with slice
type Stack struct {
	items []int
}

// ✅ Corrected: Push with pointer receiver
func (s *Stack) Push(val int) {
	s.items = append(s.items, val)
}

// ✅ Corrected: Pop with pointer receiver
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top
}

// Peek method (doesn't modify stack, so value receiver is fine)
func (s Stack) Peek() int {
	if len(s.items) == 0 {
		fmt.Println("Stack is empty")
		return -1
	}
	return s.items[len(s.items)-1]
}

// IsEmpty method (doesn't modify stack, so value receiver is fine)
func (s Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func main() {
	stack := &Stack{} // Use pointer to struct

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
