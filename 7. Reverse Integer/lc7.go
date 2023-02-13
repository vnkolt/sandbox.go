package main

// 7. Reverse Integer
//
// https://leetcode.com/problems/reverse-integer/
//
// Given a signed 32-bit integer x, return x with its digits reversed.
// If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:
// Input: x = 123
// Output: 321
//
// Example 2:
// Input: x = -123
// Output: -321
//
// Example 3:
// Input: x = 120
// Output: 21

// Constraints:
// -2^31 <= x <= 2^31 - 1

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if !(math.MinInt32 <= x && x <= (math.MaxInt32-1)) {
		return 0
	}

	reversed := 0

	for x != 0 {
		rem := x % 10
		reversed = reversed*10 + rem
		x /= 10
	}

	if !(math.MinInt32 <= reversed && reversed <= (math.MaxInt32-1)) {
		return 0
	}

	return reversed
}

func test(x int, expected int) {
	res := reverse(x)
	fmt.Printf("Input: x = %d\nOutput: %d\n", x, res)

	if res != expected {
		fmt.Printf("Result failed, should be: %d", expected)
	}
}

func main() {
	test(123, 321)
	test(-123, -321)
	test(120, 21)
	test(2147483648, 0)
	test(1534236469, 0)
}
