// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer.
// Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false. head = [3,2,0,-4], pos = 1

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}

	}
	return false
	// time = O(n), space =  O(1)
}

func hasCycleMap(head *ListNode) bool {
	visited := make(map[*ListNode]bool)

	current := head
	for current.Next != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}
	return false
	// time = O(n), space = O(1) due to map
}

func main() {
	node1 := &ListNode{Val: 3}
	node2 := &ListNode{Val: 2}
	node3 := &ListNode{Val: 0}
	node4 := &ListNode{Val: -4}

	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node2

	fmt.Println(hasCycle(node1))
	fmt.Println(hasCycleMap(node1))
}
