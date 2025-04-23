package main

import "fmt"

func StrPallindrome(str string) (string, bool) {
	var rev string
	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + string(str[i])
	}
	if rev == str {
		return rev, true
	}
	return rev, false

}

func NumPallindrome(num int) (int, bool) {
	var rev int
	temp := num
	for temp != 0 {
		rev = rev*10 + temp%10
		temp = temp / 10
	}
	if rev == num {
		return rev, true
	}
	return rev, true
}

func CheckTypeAndPallindrome(value interface{}) {

	switch v := value.(type) {
	case string:
		fmt.Println(StrPallindrome(v))
	case int:
		fmt.Println(NumPallindrome(v))
	default:
		fmt.Println("type unknown")
	}
}

func main() {
	// fmt.Println(StrPallindrome("madam"))
	// fmt.Println(StrPallindrome("madamme"))
	// fmt.Println(NumPallindrome(12321))

	CheckTypeAndPallindrome("Hello")
	CheckTypeAndPallindrome(3.12)

}
