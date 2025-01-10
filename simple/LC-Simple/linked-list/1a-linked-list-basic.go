package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// InsertAtBeginning adds a new node at the start of the linked list
func insertAtBeginning(head **ListNode, value int) {
	// Existing linked list head -> [10] -> [20] -> [30] -> nil
	// Step 1: Create the new node
	newNode := &ListNode{Val: value}

	// Step 2: Point the New node's Next to the current head,
	// newNode -> [5] ->> (head) [10] -> [20] -> [30] -> nil
	// head -> [10] -> [20] -> [30] -> nil
	newNode.Next = *head

	// Step 3: Update head to point to newNode.
	*head = newNode
	// head -> [5] -> [10] -> [20] -> [30] -> nil
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
	// Initialize an empty linked list
	var head *ListNode

	// Insert values at the beginning
	// By using &head and passing it as **ListNode, you allow the function to modify the original head in main as well.
	insertAtBeginning(&head, 10) // head -> [10] -> nil
	insertAtBeginning(&head, 20) // head -> [20] -> [10] -> nil
	insertAtBeginning(&head, 30) // head -> [30] -> [20] -> [10] -> nil

	// Print the linked list
	// current := head
	// for current != nil {
	// 	fmt.Printf("%d -> ", current.Val)
	// 	current = current.Next
	// }
	// fmt.Println("nil")
}
