package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

func PrintDLL(head *Node) {
	current := head
	fmt.Printf("head/nil <-> ")
	for current != nil {
		fmt.Printf("[%d] <-> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func createList1(head **Node, val int) {
	nn := &Node{data: val, next: *head}
	if *head != nil {
		(*head).prev = nn
	}
	*head = nn
}

func mergeDLL(list1, list2 *Node) *Node {

	dummy := &Node{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.data <= list2.data {
			current.next = list1
			list1.prev = current
			list1 = list1.next
		} else {
			current.next = list2
			list2.prev = current
			list2 = list2.next
		}
		current = current.next
	}

	if list1 != nil {
		current.next = list1
	} else if list2 != nil {
		current.next = list2
	}

	// if list1 != nil {
	// 	current.next = list1
	// 	list1.prev = current
	// 	// current = current.next
	// }

	// if list2 != nil {
	// 	current.next = list2
	// 	list2.prev = current
	// }
	return dummy.next
}

func main() {
	// var list1 *Node
	// var list2 *Node
	// Create first sorted DLL: 1 <-> 3 <-> 5
	l1 := &Node{data: 1}
	l1.next = &Node{data: 3, prev: l1}
	l1.next.next = &Node{data: 5, prev: l1.next}

	// Create second sorted DLL: 2 <-> 4 <-> 6
	l2 := &Node{data: 2}
	l2.next = &Node{data: 4, prev: l2}
	l2.next.next = &Node{data: 6, prev: l2.next}

	PrintDLL(l1)
	PrintDLL(l2)

	merged := mergeDLL(l1, l2)
	PrintDLL(merged)
}
