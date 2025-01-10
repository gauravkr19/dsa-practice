package main

import "fmt"

func checkPalindrome(num int) (bool, int) {
	var rev int
	temp := num
	for temp != 0 {
		rev = rev*10 + temp%10 // creates rev by extracting a digit from end
		temp = temp / 10       // eliminate the last digit
	}
	if rev == num {
		return true, rev
	}
	return false, rev
}

// 12321 / 10 = 1232
// 12321 % 10 = 1

// 12321
// rev = 0  +  1 = 1, temp = 12321
// rev = 10 +  2 = 12, temp = 1232
// rev = 120 + 3 = 123, temp = 123
// rev = 1230 + 2 = 1232. temp = 12
// rev = 12320 + 1 = 12321. temp = 1

func palindromeString(str string) bool {

	var rev string
	for i := len(str) - 1; i >= 0; i-- {
		rev = rev + string(str[i])
	}
	if rev == str {
		return true
	}
	return false
}

func main() {
	strNumber := "1234321"
	numNumber := 12345

	if palindromeString(strNumber) {
		fmt.Printf("%s Is a palindrome \n", strNumber)
	}

	isPalindrome, rev := checkPalindrome(numNumber)
	if isPalindrome {
		fmt.Printf("%d Is a palindrome \n", rev)
	} else {
		fmt.Printf("%d Is NOT palindrome \n", rev)
	}
}
