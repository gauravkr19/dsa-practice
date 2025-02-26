package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

// Prints in forward direction
func PrintForwardDLL(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("[%d] -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

// Inserts a new node at beginning of list
func insertAtBeginning(head **Node, val int) {
	// Create a new node with next pointing to the current head
	nn := &Node{data: val, next: *head}

	// If the list is not empty, update the old head's prev to point to new node
	if *head != nil {
		(*head).prev = nn
	}
	// Update head to point to the new node
	*head = nn
}

// Insert at end of DLL
func insertAtEnd(head **Node, val int) {
	nn := &Node{data: val}

	if *head == nil {
		*head = nn
		return
	}
	current := *head
	// Traversal, to get the last node
	for current.next != nil {
		current = current.next
	}
	nn.prev = current
	current.next = nn
}

// Insert after a node with int 10
func insertAfter(head **Node, val int) {
	current := *head
	// Traversal, to check if data exist in last node
	for current != nil {
		if current.data == 10 {
			break
		}
		current = current.next
	}
	// If 10 is not found, return
	if current == nil {
		return
	}
	nn := &Node{data: val, next: current.next, prev: current}
	// If current.next == nil {..} (case: 10 is the last node), this will cause a nil pointer dereference without if
	if current.next != nil {
		current.next.prev = nn
	}
	current.next = nn
}

// Insert a new node before a given node 20
func insertBefore(head **Node, val int) {
	current := *head
	for current != nil {
		if current.data == 20 {
			break
		}
		current = current.next
	}
	// data 20 not in DLL
	if current == nil {
		return
	}
	nn := &Node{data: val, next: current, prev: current.prev}

	// Edge case: If inserting before the first node, update head
	if current.prev == nil {
		*head = nn
	} else {
		current.prev.next = nn
	}
	current.prev = nn
}

// Deletion: 3 cases : Delete a node from the beginning of DLL
// The DLL has multiple nodes → It correctly moves head to the next node and clears .prev.
// The DLL has only one node → It correctly sets head = nil, making the list empty.
// The DLL is already empty → It simply returns without errors
func deleteAtBeginning(head **Node) {
	if *head == nil {
		return // No node to delete
	}

	// Move head to the next node
	*head = (*head).next

	// If the new head exists, remove prev link to old head
	if *head != nil {
		(*head).prev = nil
	}
}

// Delete a node from the end of DLL
func deleteAtEnd(head **Node) {
	//empty list
	if *head == nil {
		return
	}
	// delete the only node
	if (*head).next == nil {
		*head = nil
		return
	}
	current := *head
	for current.next != nil {
		current = current.next
	}
	current.prev.next = nil
}

// Delete a node by value 15 from DLL
func deleteNodeByValue(head **Node, val int) {
	// Handle empty list
	if *head == nil {
		return
	}

	current := *head
	// Traverse the list to find the node
	for current != nil && current.data != val {
		current = current.next
	}

	// Node not found
	if current == nil {
		return
	}

	// If the node to delete is the head
	if current.prev != nil {
		current.prev.next = current.next
	} else {
		// Update the head
		*head = current.next
	}

	// If the node is not the last node, update the next node's prev pointer
	if current.next != nil {
		current.next.prev = current.prev
	}

	// Optional: clean up the current node's pointers for safety
	current.next = nil
	current.prev = nil
}

// Delete a node by position (index) from DLL
func deleteNodeByPos(head **Node, position int) {
	if *head == nil || position < 0 { // Handle empty list or invalid position
		return
	}
	current := *head
	index := 0

	// Traverse to the node at the given position
	for current != nil && index < position {
		current = current.next
		index++
	}
	// If position is out of bounds
	if current == nil {
		return
	}
	// Is it the first node
	if current.prev == nil {
		*head = current.next
	} else {
		current.prev.next = current.next
	}
	// It should not be the last node
	if current.next != nil {
		current.next.prev = current.prev
	}
	// Optional: Clean up the node
	current.next = nil
	current.prev = nil
}

// Reverse a DLL
func reverseDLL(head **Node) {
	if *head == nil {
		return
	}
	var temp *Node
	current := *head
	// swap next and prev for all nodes
	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}
	// Update head to the last processed node (previous `temp`)
	if temp != nil {
		*head = temp.prev
	}
}

// Creating a DLL from a slice is more efficient than repeated insertAtEnd calls.
// Since we track current, we avoid O(n) traversal per insertion.
func createDLLFromSlice(list []int) *Node {
	head := &Node{data: list[0]}
	current := head
	for i := 0; i < len(list); i++ {
		nn := &Node{data: list[i]}
		current.next = nn
		nn.prev = current
		current = nn
	}
	return head
}

func main() {
	var head *Node

	// Insert Ops
	insertAtBeginning(&head, 10)
	insertAtBeginning(&head, 5)
	insertAtEnd(&head, 20)  // [5] -> [10] -> [20] -> nil
	insertAfter(&head, 15)  // Need to insert 15 after node 10, [5] -> [10] -> [15] -> [20] -> nil
	insertBefore(&head, 17) // insert before node 20

	// Delete ops
	deleteAtBeginning(&head)
	deleteAtEnd(&head)
	deleteNodeByValue(&head, 15) // delete node with value 15
	deleteNodeByPos(&head, 1)    // delete node with position 1

	insertAtBeginning(&head, 7)
	insertAtBeginning(&head, 5)
	insertAtEnd(&head, 15) // [5] -> [7] -> [10] -> [15] -> nil

	reverseDLL(&head) // output:[15] -> [10] -> [7] -> [5] -> nil
	PrintForwardDLL(head)

	// Create DLL from a Slice
	createDLLFromSlice([]int{3, 7, 8})
	PrintForwardDLL(head)
}
