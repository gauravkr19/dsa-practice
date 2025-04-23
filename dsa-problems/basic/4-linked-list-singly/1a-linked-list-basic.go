package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// InsertAtBeginning adds a new node at the start of the linked list  (using a double pointer)
func insertAtBeginning(head **ListNode, value int) {
	// Step 1: Create the new node
	newNode := &ListNode{Val: value}

	fmt.Printf("New node created:\n(newNode)=%p, Data (newNode.Val)=%d, Next (newNode.Next)=%p, &head=%p\n", newNode, newNode.Val, newNode.Next, &head)

	// Print old head node
	if *head != nil {
		fmt.Printf("old head (*head): %p, Data (*head).Val: %d, **head: %v\n", *head, (*head).Val, **head)
	} else {
		fmt.Printf("old head (*head): %p \n", *head)
	}

	// Step 2: Point the New node's Next to the old head or nil for 1st node
	// *head retrieves the value stored in head, which is the address of the old head node.
	fmt.Println()
	newNode.Next = *head
	fmt.Printf("New node after assignment:\n(newNode): %p, Data (newNode.Val): %d, Next (newNode.Next) -> %p\n", newNode, newNode.Val, newNode.Next)

	// Step 3: Update head to point to newNode.
	*head = newNode
	fmt.Println()
	fmt.Printf("Updated head (*head): %p, Data (*head).Val: %d\n\n", *head, (*head).Val)
	fmt.Println("=========================================================")
	fmt.Println()
}

// Alternative function using single pointer (returns the new head)
func insertAtBeginningSingle(head *ListNode, value int) *ListNode {
	// Step 1: Create the new node
	newNode := &ListNode{Val: value}

	fmt.Printf("New node created: newNode=%p, Data: %d, newNode.Next -> %p\n", newNode, newNode.Val, newNode.Next)

	// Print old head node
	if head != nil {
		fmt.Printf("old head (head)=%p, (*head)=%v, Data=%d\n", head, *head, head.Val)
	} else {
		fmt.Println("old head: nil")
	}

	fmt.Println("=========================================================")
	fmt.Println()
	// Step 2: Point the new node's Next to the old head
	newNode.Next = head
	fmt.Printf("After updating: old (head)=%p, newNode.Next=%p  Data=%d, newNode.Next -> %p\n", head, newNode, newNode.Val, newNode.Next)

	fmt.Println()
	// Step 3: Return the new head
	fmt.Printf("Updated head: %p, Data: %d\n", newNode, newNode.Val)
	return newNode
}

// Function to print the linked list
func printList(head *ListNode) {
	fmt.Print("Linked List: ")
	for head != nil {
		fmt.Printf("[%d] -> ", head.Val)
		head = head.Next
	}
	fmt.Println("nil")
}

// InsertAtEnd adds a new node at the end of the linked list
// By using &head and passing it as **ListNode, you allow the function to modify the original head.
func insertAtEnd(head **ListNode, value int) {
	// Step 1: Create the new node
	newNode := &ListNode{Val: value, Next: nil}

	// Step 2: If the list is empty, make the new node the head
	if *head == nil {
		*head = newNode
		return
	}

	// Step 3: Traverse to the last node
	current := *head
	for current.Next != nil {
		current = current.Next
	}

	// Step 4: Point the last node's Next to the new node
	current.Next = newNode
}

func main() {
	fmt.Println("***Using double pointer approach:***")
	var head *ListNode
	// Print address of `head` in `main`
	fmt.Printf("[main] &head (original head address in main): %p\n", &head)

	insertAtBeginning(&head, 10) // head -> [10] -> nil
	insertAtBeginning(&head, 20) // head -> [20] -> [10] -> nil
	insertAtBeginning(&head, 30) // head -> [30] -> [20] -> [10] -> nil
	printList(head)
	fmt.Println()

	fmt.Println("\n***Using single pointer approach***:")
	fmt.Println()
	var head2 *ListNode

	fmt.Printf("[main] head2 (original head2 address in main): %p\n", head2)
	head2 = insertAtBeginningSingle(head2, 10) // head2 -> [10] -> nil

	fmt.Printf("[main] head2 (original head2 address in main): %p\n", head2)
	head2 = insertAtBeginningSingle(head2, 20) // head2 -> [20] -> [10] -> nil

	fmt.Printf("[main] head2 (original head2 address in main): %p\n", head2)
	head2 = insertAtBeginningSingle(head2, 30) // head2 -> [30] -> [20] -> [10] -> nil

	fmt.Printf("[main] head2 (original head2 address in main): %p\n", head2)
	printList(head2)

	var head3 *ListNode
	insertAtEnd(&head3, 30)
	insertAtEnd(&head3, 20)
	insertAtEnd(&head3, 10)
	printList(head3)
}
