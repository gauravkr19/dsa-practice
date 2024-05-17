package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

// Euclidean algorithm works by repeatedly applying the following equality until the remainder is zero:
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// For two strings s and t, we say "t divides s" if and only if s = t + ... + t (i.e., t is concatenated with itself one or more times).
// Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))  // Output: "ABC"
	fmt.Println(gcdOfStrings("ABCD", "ABCABC")) // Output: "ABC"
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // Output: "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // Output: ""
}
