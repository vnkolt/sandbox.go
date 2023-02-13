package main

// 5. Longest Palindromic Substring
//
// https://leetcode.com/problems/longest-palindromic-substring/description/
//
// Given a string s, return the longest palindromic substring in s.
//
// Example 1:
//
// Input: s = "babad"
// Output: "bab"
// Explanation: "aba" is also a valid answer.
//
// Example 2:
// Input: s = "cbbd"
// Output: "bb"
//
// Constraints:
// 1 <= s.length <= 1000
// s consist of only digits and English letters.

import (
	"fmt"
)

func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}

	max_left, max_len := 0, 1

	for start := 0; start < len && len-start > max_len/2; {
		left := start
		right := start
		for right < len-1 && s[right+1] == s[right] {
			right++
		}

		start = right + 1
		for right < len-1 && left > 0 && s[right+1] == s[left-1] {
			right++
			left--
		}

		if max_len < right-left+1 {
			max_left = left
			max_len = right - left + 1
		}
	}

	return s[max_left : max_left+max_len]
}

func test(s string) {
	fmt.Printf("\nInput: s = %v\n", s)
	fmt.Printf("Output: %v", longestPalindrome(s))
}

func main() {
	test("babad")               // bab
	test("cbbd")                // bb
	test("what is abuduba?")    // abuduba
	test("madamimadam")         // madamimadam
	test("liderbredil")         // liderbredil
	test("saippuakivikauppias") // saippuakivikauppias
}
