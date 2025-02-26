package main

import "fmt"

// Node structure for Doubly Linked List
type Node struct {
	val  int
	next *Node
	prev *Node
}

// Function to merge two sorted DLLs
func mergeSortedDLL(l1, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// Dummy node to act as the start of the merged DLL
	dummy := &Node{val: 0}
	tail := dummy

	// Merge the two lists
	for l1 != nil && l2 != nil {
		if l1.val <= l2.val {
			tail.next = l1
			l1.prev = tail
			l1 = l1.next
		} else {
			tail.next = l2
			l2.prev = tail
			l2 = l2.next
		}
		tail = tail.next
	}

	// Attach the remaining nodes
	if l1 != nil {
		tail.next = l1
		l1.prev = tail
	} else if l2 != nil {
		tail.next = l2
		l2.prev = tail
	}

	// Return the head of the merged DLL
	mergedHead := dummy.next
	if mergedHead != nil {
		mergedHead.prev = nil
	}
	return mergedHead
}

// Function to print the DLL
func printDLL(head *Node) {
	for head != nil {
		fmt.Printf("%d ", head.val)
		head = head.next
	}
	fmt.Println()
}

// Helper function to create a DLL from a slice
func createDLL(arr []int) *Node {
	if len(arr) == 0 {
		return nil
	}
	head := &Node{val: arr[0]}
	current := head
	for i := 1; i < len(arr); i++ {
		newNode := &Node{val: arr[i]}
		current.next = newNode
		newNode.prev = current
		current = newNode
	}
	return head
}

func main() {
	// Create two sorted DLLs
	l1 := createDLL([]int{1, 3, 5, 7})
	l2 := createDLL([]int{2, 4, 6, 8})

	fmt.Print("List 1: ")
	printDLL(l1)
	fmt.Print("List 2: ")
	printDLL(l2)

	// Merge and print the sorted DLL
	mergedHead := mergeSortedDLL(l1, l2)
	fmt.Print("Merged DLL: ")
	printDLL(mergedHead)
}
