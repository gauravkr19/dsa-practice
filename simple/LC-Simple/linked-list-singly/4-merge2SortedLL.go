// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLL(list1, list2 *ListNode) *ListNode {

	dummy := &ListNode{}
	op := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			op.Next = list1
			list1 = list1.Next
		} else {
			op.Next = list2
			list2 = list2.Next
		}
		op = op.Next
	}
	if list1 != nil {
		op.Next = list1
	}

	if list2 != nil {
		op.Next = list2
	}
	return dummy.Next
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
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 4}
	node1.Next = node2
	node2.Next = node3
	l1 = node1

	l2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	mergedLL := mergeTwoLL(l1, l2)
	PrintLL(mergedLL)
}
