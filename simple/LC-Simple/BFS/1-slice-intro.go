// The primary goal of BFS is to visit every node in the shortest path order (layer by layer) while
// ensuring no node is visited twice.

// ***Algorithm for BFS***
// 1. Initialize a Queue: Start by pushing the starting node into a queue.
// 2. Mark the Node as Visited: Use a visited set/map/array to keep track of visited nodes.
// 3. Loop until the Queue is Empty:
// 		a. Dequeue (remove) a node from the front of the queue.
// 		b. Process the current node (print/store the value, etc.).
// 		c. Push all its unvisited neighbors into the queue and mark them as visited.
// 4. Repeat until all reachable nodes are visited.

// ***Pseudocode for BFS***
// BFS(graph, startNode):
//     Initialize an empty queue Q
//     Initialize a set or array to track visited nodes
//     Add startNode to Q and mark it as visited

//     while Q is not empty:
//         current = Q.dequeue()  // Remove front node from queue
//         Process current node   // (print/store the value)

//         for each neighbor in graph[current]:  // Iterate over neighbors
//             if neighbor is not visited:
//                 Add neighbor to Q
//                 Mark neighbor as visited

package main

import "fmt"

// BFS function to traverse a graph using a queue
func BFS(graph map[int][]int, start int) {
	visited := make(map[int]bool) // Track visited nodes
	queue := []int{start}         // Initialize queue with the start node
	visited[start] = true         // Mark start node as visited

	for len(queue) > 0 {
		node := queue[0]     // Get the front element
		queue = queue[1:]    // Dequeue a node
		fmt.Print(node, " ") // Process the current node

		// Visit all unvisited neighbors
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor) // Enqueue node's children
				visited[neighbor] = true        // Mark as visited
			}
		}
	}
}

func main() {
	// Graph representation (Adjacency List)
	graph := map[int][]int{
		0: {1, 2},
		1: {0, 3, 4},
		2: {0, 5, 6},
		3: {1},
		4: {1},
		5: {2},
		6: {2},
	}

	fmt.Println("BFS Traversal:")
	BFS(graph, 0) // Start BFS from node 0
}
