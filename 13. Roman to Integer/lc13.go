package main

// https://leetcode.com/problems/roman-to-integer/
//
// 13. Roman to Integer
//
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:
//
// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.
//
//
// Example 1:
//
// Input: s = "III"
// Output: 3
// Explanation: III = 3.
//
// Example 2:
// Input: s = "LVIII"
// Output: 58
// Explanation: L = 50, V= 5, III = 3.
//
// Example 3:
// Input: s = "MCMXCIV"
// Output: 1994
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
//
// Constraints:
// 1 <= s.length <= 15
// s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// It is guaranteed that s is a valid roman numeral in the range [1, 3999].

import (
	"fmt"
	"regexp"
)

func validRomanNumber(s string) bool {
	match, _ := regexp.MatchString("^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$", s)
	return match
}

func romanToInt(s string) int {

	if !validRomanNumber(s) {
		return 0
	}

	rmap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	slen := len(s)

	for i := 0; i < slen; i++ {
		v1 := rmap[s[i]]

		if i+1 < slen {
			v2 := rmap[s[i+1]]

			if v1 >= v2 {
				result += v1
			} else {
				result += v2 - v1
				i++
			}

		} else {
			result += v1
		}
	}

	return result
}

func test(roman_num string, expected int) {
	result := romanToInt(roman_num)

	fmt.Printf("Input: roman number = %s\nOutput: %d\n", roman_num, result)

	if result != expected {
		fmt.Printf("Result failed, should be: %d\n", expected)
	}
}

func main() {

	test("MMMDCCXXIV", 3724)

	test("I", 1)
	test("III", 3)
	test("LVIII", 58)
	test("MCMXCIV", 1994)
	test("MMMCMXCIX", 3999)

}
