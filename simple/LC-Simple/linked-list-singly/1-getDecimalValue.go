// Given head which is a reference node to a singly-linked list.
// The value of each node in the linked list is either 0 or 1.
// The linked list holds the binary representation of a number.
// Return the decimal value of the number in the linked list.
// The most significant bit is at the head of the linked list.

package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Function to compute the decimal value from the binary linked list
func getDecimalValue(head *ListNode) int {
	result := 0

	// Traverse the linked list
	for head != nil {
		// Left shift the result by 1 (multiply by 2) and add the current node's value
		result = result*2 + head.Val
		// Move to the next node
		head = head.Next
	}

	return result
}

func main() {
	// Example Linked List: 1 -> 0 -> 1 (Binary: 101) 1, 2+0=2, 2*2+1=5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 0}
	head.Next.Next = &ListNode{Val: 1}
	// head.Next.Next points to this new node.
	// This means the Next field of the second node (value 0) is updated to point to the new node.

	// Calculate and print the decimal value
	fmt.Println(getDecimalValue(head)) // Output: 5
}
