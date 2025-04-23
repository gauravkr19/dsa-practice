package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	data int
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

func checkPalindome(head *Node) bool {
	//empty DLL, one node DLL
	if head == nil || head.next == nil {
		return true
	}

	// Find the tail of DLL
	current := head
	for current.next != nil {
		current = current.next
	}

	// 	the loop now stops when:
	// front == back (odd-length list, middle reached)
	// front.prev == back (even-length list, pointers crossed)
	// Use two-pointer approach to check for palindrome
	front, back := head, current
	for front != back && front.prev != back { // Stop when pointers cross or meet
		if front.data == back.data {
			front = front.next
			back = back.prev
		} else {
			return false
		}
	}
	return true
}

// find the middle node in DLL
func findMiddle(head *Node) (*Node, *Node) {

	slow, fast := head, head
	// fast.next = detects odd length
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	if fast == nil {
		return slow.prev, slow
	}
	return slow, nil
}

func main() {
	var head *Node
	insertAtBeginning(&head, 50)
	insertAtBeginning(&head, 40)
	insertAtBeginning(&head, 30)
	insertAtBeginning(&head, 30)
	insertAtBeginning(&head, 40)
	insertAtBeginning(&head, 50)
	// insertAtBeginning(&head, 60)

	fmt.Println(checkPalindome(head))
	fmt.Println("Palindrome")
	PrintForwardDLL(head)

	fmt.Println("Find Middle")
	fmt.Println(findMiddle(head))
	// PrintForwardDLL(head)
}
