package main

// 9. Palindrome Number
//
// https://leetcode.com/problems/palindrome-number/
//
// Given an integer x, return true if x is a  palindrome, and false otherwise.
//
// Example 1:
// Input: x = 121
// Output: true
// Explanation: 121 reads as 121 from left to right and from right to left.

// Example 2:
//
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not a palindrome.

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false // Negative numbers aren't polindromes
	}
	if x < 10 {
		return true // Any number from 0 to 9 inclusive is a palindrome
	}

	numbers := [32]int{} // 32 is more than enough, but I like 32
	len := 0

	// put the numbers of the decimal representation of number "x" into array "numbers", e.g. if x = 1234 then array "numbers" contains 1, 2, 3, 4
	for x != 0 { // Golang syntax doesn't allow "for x" without comparing x with something, in C/C++ we can write "while(x)" that means "while(x!=0)"
		numbers[len] = x % 10
		len++
		x /= 10
	}

	for i := 0; i < len/2; i++ {
		if numbers[i] != numbers[len-i-1] { // Items equidistant from the middle must be equal.
			return false
		}
	}

	return true
}

func test(x int, expected bool) {
	res := isPalindrome(x)
	fmt.Printf("Input: x = %d\nOutput: %v\n", x, res)

	if res != expected {
		fmt.Printf("Test failed, result must be: %v", expected)
	}
}

func main() {
	test(121, true)
	test(-121, false)
	test(10, false)
	test(123454321, true)
}
