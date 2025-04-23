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
// gcd(a, b) always makes the second parameter b go to the first position, and then the second becomes a % b
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
	fmt.Println(gcdOfStrings("ABCD", "ABCABC")) // Output: ""
	fmt.Println(gcdOfStrings("ABABAB", "ABAB")) // Output: "AB"
	fmt.Println(gcdOfStrings("LEET", "CODE"))   // Output: ""
}

// Time complexity
// O(N + M) for string work
// O(log(min(N, M)) for integer GCD
// Total: O(N + M + log(min(N, M)))

// Here’s how:
// Euclidean algorithm has a well-known time complexity of O(log(min(a, b)))
// So when you do gcd(6, 3) or gcd(1000000, 3) → it finishes in log(min(6, 3)) or log(3) steps
// Still very fast — the slowest part is string concatenation!

// “Euclidean algorithm makes progress by shrinking the input quickly.
// It terminates fast, typically in log(min(a, b)) steps, even in worst case like Fibonacci numbers.”
