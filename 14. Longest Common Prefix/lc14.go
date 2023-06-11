package main

// https://leetcode.com/problems/longest-common-prefix/
//
// 14. Longest Common Prefix
//
// Write a function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
//
// Constraints:
//
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters.

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	min_len := len(strs[0])

	for i := 1; i < len(strs); i++ {
		if len(strs[i]) < min_len {
			min_len = len(strs[i])
		}
	}

	if min_len == 0 {
		return ""
	}

	for max_idx := 0; max_idx < min_len; max_idx++ {
		c := strs[0][max_idx]
		for i := 1; i < len(strs); i++ {
			if strs[i][max_idx] != c {
				return strs[0][0:max_idx]
			}
		}
	}

	return strs[0][0:min_len]

}

func test(strs []string, expected string) {
	result := longestCommonPrefix(strs)

	fmt.Printf("Input: strs = %v\nOutput: %s\n", strs, result)

	if result != expected {
		fmt.Printf("Result failed, should be: %s\n", expected)
	}
}

func main() {
	test([]string{"flower", "flow", "flight"}, "fl")
	test([]string{"dog", "racecar", "car"}, "")
}
