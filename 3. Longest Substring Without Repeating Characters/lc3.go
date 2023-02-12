package main

// https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
// Given a string s, find the length of the longest substring without repeating characters.
//
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

import (
	"fmt"
)

func IsUnique(s string, pos int, len int) bool {
	for i := pos; i < pos+len-1; i++ {
		for j := i + 1; j < pos+len; j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}
	return true
}

func lengthOfLongestSubstring(s string) int {
	if len(s) < 1 {
		return 0
	}

	max_longest_substring_len := 1

	for start := 0; start < len(s); start++ {
		for end := start + 1; end < len(s); end++ {
			if end-start+1 > max_longest_substring_len {
				if !IsUnique(s, start, end-start+1) {
					break
				}
				if end-start+1 > max_longest_substring_len {
					max_longest_substring_len = end - start + 1
				}
			}
		}
	}

	return max_longest_substring_len
}

func test(s string, answer int) {
	fmt.Printf("Input string: s = %v\n", s)
	res := lengthOfLongestSubstring(s)

	fmt.Printf("Lenght of the longest substring without repeating characters is %d\n", res)
	if answer != res {
		fmt.Printf("Test failed for string %v, result=%d, expected result = %d\n", s, res, answer)
	}

}

func main() {
	test("abcabcbb", 3)
	test("bbbbb", 1)
	test("pwwkew", 3)
}
