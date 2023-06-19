package main

// https://leetcode.com/problems/3sum/
//
// 15. 3Sum
//
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.
//
// Example 1:
//
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
//
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
//
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.
//
// Constraints:
//
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 105

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	zeros := make([][]int, 0)

	sort.Ints(nums)

	for idx := 0; idx < len(nums)-2; idx++ {
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		for lidx, ridx := idx+1, len(nums)-1; lidx < ridx; {
			sum := nums[idx] + nums[lidx] + nums[ridx]

			if sum < 0 {
				lidx++
			} else if sum > 0 {
				ridx--
			} else {
				zeros = append(zeros, []int{nums[idx], nums[lidx], nums[ridx]})

				for lidx < ridx && nums[lidx] == nums[lidx+1] {
					lidx++
				}
				for lidx < ridx && nums[ridx] == nums[ridx-1] {
					ridx--
				}

				lidx++
				ridx--

			}
		}
	}
	return zeros
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func Equal2D(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !Equal(a[i], b[i]) {
			return false
		}
	}

	return true
}

func test(nums []int, expected [][]int) {
	result := threeSum(nums)

	fmt.Printf("Input: nums = %v\nOutput: %v\n", nums, result)

	if !Equal2D(result, expected) {
		fmt.Printf("Result failed, should be: %v\n", expected)
	}
}

func main() {
	test([]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}})
	test([]int{0, 1, 1}, [][]int{})
}
