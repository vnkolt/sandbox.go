package main

// 8. String to Integer (atoi)
//
// https://leetcode.com/problems/string-to-integer-atoi/

// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer (similar to C/C++'s atoi function).

// The algorithm for myAtoi(string s) is as follows:

// Read in and ignore any leading whitespace.
// Check if the next character (if not already at the end of the string) is '-' or '+'. Read this character in if it is either. This determines if the final result is negative or positive respectively. Assume the result is positive if neither is present.
// Read in next the characters until the next non-digit character or the end of the input is reached. The rest of the string is ignored.
// Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
// If the integer is out of the 32-bit signed integer range [-2^31, 2^31-1], then clamp the integer so that it remains in the range.
// Specifically, integers less than -2^31 should be clamped to -2^31, and integers greater than 2^31 - 1 should be clamped to 2^31-1.
// Return the integer as the final result.
// Note:

// Only the space character ' ' is considered a whitespace character.
// Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.

// Example 1:
// Input: s = "42"
// Output: 42
// Explanation: The underlined characters are what is read in, the caret is the current reader position.
// Step 1: "42" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "42" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "42" ("42" is read in)
//            ^
// The parsed integer is 42.
// Since 42 is in the range [-2^31, 2^31-1], the final result is 42.
//
// Example 2:
// Input: s = "   -42"
// Output: -42
// Explanation:
// Step 1: "   -42" (leading whitespace is read and ignored)
//             ^
// Step 2: "   -42" ('-' is read, so the result should be negative)
//              ^
// Step 3: "   -42" ("42" is read in)
//                ^
// The parsed integer is -42.
// Since -42 is in the range [-2^31, 2^31-1], the final result is -42.
//
// Example 3:

// Input: s = "4193 with words"
// Output: 4193
// Explanation:
// Step 1: "4193 with words" (no characters read because there is no leading whitespace)
//          ^
// Step 2: "4193 with words" (no characters read because there is neither a '-' nor '+')
//          ^
// Step 3: "4193 with words" ("4193" is read in; reading stops because the next character is a non-digit)
//              ^
// The parsed integer is 4193.
// Since 4193 is in the range [-2^31, 2^31-1], the final result is 4193.

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {

	if len(s) < 1 {
		return 0
	}

	str := []rune(s)

	result, sign_mult, i := 0, 1, 0

	// Skip any leading whitespace
	for i < len(str) && unicode.IsSpace(str[i]) {
		i++
	}

	// Check if the next character (if not already at the end of the string) is '-' or '+'.
	// Read this character in if it is either. This determines if the final result is negative or positive respectively.
	// Assume the result is positive if neither is present.
	if i < len(str) {
		if str[i] == '+' {
			sign_mult = 1
			i++
		} else if str[i] == '-' {
			sign_mult = -1
			i++
		}
	}

	for i < len(str) && unicode.IsDigit(str[i]) {
		result *= 10
		if result < 0 {
			break
		} // integer overflow?
		result += int(str[i] - '0')
		if result < 0 {
			break
		} // integer overflow?
		i++
	}

	if result < 0 { // integer overflow?
		if sign_mult == 1 {
			result = math.MaxInt32
		} else {
			result = math.MinInt32
		}
	} else {
		result *= sign_mult
		if result < math.MinInt32 {
			result = math.MinInt32
		} else if result > math.MaxInt32 {
			result = math.MaxInt32
		}
	}

	return result
}

func test(s string, expected int) {
	res := myAtoi(s)
	fmt.Printf("Input: s = %v\nOutput: %d\n", s, res)

	if res != expected {
		fmt.Printf("Result failed, should be: %d", expected)
	}
}

func main() {
	test("42", 42)
	test("-42", -42)
	test("4193 with words", 4193)
	test("Invalid number 1234", 0)
	test("1234 valid number", 1234)
	test("+-+- another invalid number", 0)
	test("18446744073709551617", 2147483647) // overflow test case
	test("9223372036854775808", 2147483647)  // overflow test case
}
