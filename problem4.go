package main

import "strconv"

/*
 * Largest palindrome product
 *
 * A palindromic number reads the same both ways. The largest palindrome made
 * from the product of two 2-digit numbers is 9009 = 91 × 99.
 *
 * Find the largest palindrome made from the product of two 3-digit numbers.
 */
func problem4() int {
	// Unlike most languages, Go supports multiple encodings. Hence no unified reverse.
	// UTF-8 is used by default.
	reverse := func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}
	isPalindrome := func(n int) bool {
		s := strconv.Itoa(n)
		if s == reverse(s) {
			return true
		}
		return false
	}

	var max int

	for i := 100; i < 1000; i++ {
		for j := 100; j < 1000; j++ {
			prod := i * j
			if isPalindrome(prod) && prod > max {
				max = prod
			}
		}
	}
	return max
}
