package main

import "fmt"

// Given a string s, find the length of the longest substring without repeating characters.
// Example 1:
// Input: s = "abcabcbb"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createLinkedList(head **ListNode, val int) {
	newNode := &ListNode{Val: val}
	newNode.Next = *head
	*head = newNode
	fmt.Printf("Added %d: Head now points to %d\n", val, newNode.Val)
	if newNode.Next != nil {
		fmt.Printf("Next of %d points to %d\n", newNode.Val, newNode.Next.Val)
	}
}

// Display nodes
func DisplayList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("%d ->", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func main() {
	var head *ListNode
	createLinkedList(&head, 10)
	createLinkedList(&head, 20)
	createLinkedList(&head, 30)

	// Add a node to the front of the list
	newNodeFront := &ListNode{Val: 35}
	newNodeFront.Next = head
	head = newNodeFront

	// Display the list with new front node
	DisplayList(head)
	fmt.Println()

	// Add a node to the end of the list
	newNodeEnd := &ListNode{Val: 5, Next: nil}
	currentEnd := head
	for currentEnd.Next != nil {
		currentEnd = currentEnd.Next
	}
	currentEnd.Next = newNodeEnd

	// Display the list with new last node
	DisplayList(head)
}
