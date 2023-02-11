package main

// https://leetcode.com/problems/two-sum/
//
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example 1:
// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Example 2:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Example 3:
// Input: nums = [3,3], target = 6
// Output: [0,1]

import (
	"fmt"
)

func IndexOf(a []int, start int, target int) int {
	for i := start; i < len(a); i++ {
		if a[i] == target {
			return i
		}
	}
	return -1
}

func twoSum(nums []int, target int) []int {
	var res []int

	for i := 0; i < len(nums); i++ {
		index := IndexOf(nums, i+1, target-nums[i])
		if index != -1 {
			res = append(res, i, index)
			break
		}
	}

	return res
}

func testCase(nums []int, target int) {
	fmt.Printf("Input: nums = %v, target = %d\n", nums, target)
	res := twoSum(nums, 6)
	fmt.Printf("Output: = %v\n", res)
}

func main() {
	testCase([]int{2, 7, 11, 15}, 9) // nums = [2,7,11,15], target = 9
	testCase([]int{3, 2, 4}, 6)      // nums = [3,2,4], target = 6
	testCase([]int{3, 3}, 6)         // nums = [3,3], target = 6
}
