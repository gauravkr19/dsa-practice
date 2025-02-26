// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.

// Hints to solve the problem:
// Use a dummy node to simplify the result list creation.
// Use a carry variable to handle sums greater than 9.
// Traverse both linked lists simultaneously, adding corresponding digits and considering the carry.
// If one list is shorter, assume missing digits as 0.
// At the end, if there's a carry left, create a new node for it.

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {

	// Initialize dummyHead (creates empty node) and current to keep track of the result list.
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	// 	carry != 0 is needed in for loop consition,	When does for l1 != nil || l2 != nil fail?
	// The issue occurs when both l1 and l2 are fully processed, but a carry remains.
	// without carry != 0 in the for loop, the program ends without creating extra node for carry. example: [9,9]+[1] will give [0,0] instead of [1,0,0]
	for l1 != nil || l2 != nil || carry != 0 {

		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := x + y + carry
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	// If there's a remaining carry, add it as a new node or append to current
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	// dummyHead returns the list starting from 0 >
	return dummyHead.Next
}

func PrintLL(head *ListNode) {

	for head != nil {
		fmt.Printf("[%d] > ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

func main() {

	var l1 *ListNode
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = nil
	l1 = node1

	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}

	resultList := addTwoNumbers(l1, l2)
	PrintLL(resultList)
}
