package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Test")
	// str := "abcdef"
	str := []byte("abcdef")

	// strrune := []rune(str)
	l := len(str)

	for i := 0; i < l/2; i++ {
		j := l - i - 1

		str[i], str[j] = str[j], str[i]
	}

	fmt.Printf("%s \n", string(str))
}
