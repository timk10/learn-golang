package main

import (
	"fmt"
)
/*
A palindromic number reads the same both ways. The largest palindrome made from the product of
two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */

func isPalindrome(prod int) bool {
	s := fmt.Sprintf("%d", prod)
	length := len(s)
	for i := 0; i < length / 2; i++ {
		if s[i] != s[length - i - 1] {
			return false;
		}
	}

	return true
}

func main() {
	max := 0
	for i := 999; i > 99; i-- {
		for j := 999; j > 99; j-- {
			prod := i*j
			if isPalindrome(prod) && prod > max {
				max = prod
			}

		}
	}
	fmt.Printf("%d\n", max)
}