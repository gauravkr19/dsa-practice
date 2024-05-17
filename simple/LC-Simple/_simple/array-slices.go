package main

import "fmt"

//array vs slices

// arr1 := [4]string{'a', 'b', 'c', 'x'}
// arr2 := [4]string{'x', 'y', 'z'}

// arr3 := [...]int{1,2,3,4,4,2,2,2}

// s = make([]string, 5)
const (
	winter = 4
	summer = 5
)

func main() {

	var (
		books  [4]string
		wBooks [winter]string
		sBooks [summer]string
	)

	// fmt.Printf("Books  : %T\n", books)  // Prints the type
	// fmt.Printf("Books  : %q\n", books)  // Print in quotes
	// fmt.Printf("Books  : %#v\n", books) // prints type and element togther

	// output from above code
	// Books  : [3]string
	// Books  : ["" "" ""]
	// Books  : [3]string{"", "", ""}

	books[0] = "abc"
	books[1] = "xyz"
	books[2] += books[0] + " 2nd Ed"

	for i := 0; i < len(books); i++ {
		wBooks[i] = books[i]
	}

	fmt.Printf("wBooks: %#v\n", books)
	fmt.Printf("wBooks: %#v\n", wBooks)

	for i := range books {
		sBooks[i] = books[i]
	}

	fmt.Printf("wBooks: %#v\n", books)
	fmt.Printf("wBooks: %#v\n", wBooks)
	fmt.Printf("sBooks: %#v\n", sBooks)

	var published [len(wBooks)]bool
}
