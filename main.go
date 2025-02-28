package main

import "fmt"

// check if a string is a palindrome
func checkPalindrome(s string) bool {
	for i := 0; i <= len(s)-1; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

// Write a program to find the HCF of two numbers without using recursion
func hcf(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else {
			b = b - a
		}
	}
	return a
}

func main() {
	s := "racecar"
	fmt.Println(checkPalindrome(s))

	a := 72
	b := 9
	fmt.Println(hcf(a, b))
}
