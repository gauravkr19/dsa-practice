package main

import "fmt"

func main() {
	// Step 1: Declare an int variable
	var num int = 42
	fmt.Println("Value of num:", num)    // Output: 42
	fmt.Println("Address of num:", &num) // Example Output: 0xc000014098

	// Step 2: Create a pointer to num
	var ptr *int = &num
	fmt.Println("\nValue of ptr (address of num):", ptr)        // Example Output: 0xc000014098
	fmt.Println("Value at address stored in ptr (*ptr):", *ptr) // Output: 42

	// Step 3: Create a double pointer (pointer to ptr)
	var doublePtr **int = &ptr
	fmt.Println("\nValue of doublePtr (address of ptr):", doublePtr)                 // Example Output: 0xc000012030
	fmt.Println("Value at address stored in doublePtr (*doublePtr):", *doublePtr)    // Output: 0xc000014098
	fmt.Println("Value at address stored in *doublePtr (**doublePtr):", **doublePtr) // Output: 42

	// Step 4: Modify num using the pointer
	*ptr = 100
	fmt.Println("\nAfter changing *ptr:")
	fmt.Println("Value of num:", num)                 // Output: 100
	fmt.Println("Value at *ptr:", *ptr)               // Output: 100
	fmt.Println("Value at **doublePtr:", **doublePtr) // Output: 100

	// Step 5: Modify num using the double pointer
	**doublePtr = 200
	fmt.Println("\nAfter changing **doublePtr:")
	fmt.Println("Value of num:", num)                 // Output: 200
	fmt.Println("Value at *ptr:", *ptr)               // Output: 200
	fmt.Println("Value at **doublePtr:", **doublePtr) // Output: 200
}
