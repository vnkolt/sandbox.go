package main

// 11. Container With Most Water
//
// https://leetcode.com/problems/container-with-most-water/
//

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).
// Find two lines that together with the x-axis form a container, such that the container contains the most water.
// Return the maximum amount of water a container can store.
//
// Example 1:
//
// Input: height = [1,8,6,2,5,4,8,3,7]
// Output: 49
// Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
//
// Example 2:

// Input: height = [1,1]
// Output: 1

// Constraints:

// n == height.length
// 2 <= n <= 10^5
// 0 <= height[i] <= 10^4

import (
	"fmt"
)

func min(a int, b int) int { // math.min requres float type, that's why we need this func for int
	if a < b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	result := 0

	for ldx, rdx := 0, len(height)-1; ldx < rdx; {
		min_height := min(height[ldx], height[rdx])
		cont_square := min_height * (rdx - ldx)
		if cont_square > result {
			result = cont_square
		}

		if height[ldx] < height[rdx] {
			ldx++
		} else {
			rdx--
		}
	}

	return result
}

func test(height []int, expected int) {
	res := maxArea(height)
	fmt.Printf("Input: x = %d\nOutput: %v\n", height, res)

	if res != expected {
		fmt.Printf("Test failed, result must be: %v", expected)
	}
}

func main() {
	test([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49)
	test([]int{1, 1}, 1)
}
