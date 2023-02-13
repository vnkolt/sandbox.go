package main

// 6. Zigzag Conversion
//
// https://leetcode.com/problems/zigzag-conversion/
//
// The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this:
// (you may want to display this pattern in a fixed font for better legibility)
//
//     P   A   H   N
//     A P L S I I G
//     Y   I   R
//
// And then read line by line: "PAHNAPLSIIGYIR"
//
// Write the code that will take a string and make this conversion given a number of rows:
//
// string convert(string s, int numRows);
//
// Example 1:
// Input: s = "PAYPALISHIRING", numRows = 3
// Output: "PAHNAPLSIIGYIR"
//
// Example 2:
// Input: s = "PAYPALISHIRING", numRows = 4
// Output: "PINALSIGYAHRPI"
// Explanation:
// P     I    N
// A   L S  I G
// Y A   H R
// P     I
//
// Example 3:
// Input: s = "A", numRows = 1
// Output: "A"

import (
	"fmt"
	"strings"
)

const (
	Direction_Down    int = 0
	Direction_UpRight int = 1
)

func convert(s string, numRows int) string {
	if numRows <= 1 || len(s) <= 1 {
		return s
	}

	numCols := len(s)

	// In Golang strings are immutable
	// If we want to read/write chars from/to the string we need to convert it to an array
	str := []byte(s)

	result := []byte(strings.Repeat("#", numCols))

	// Actually it is not a matrix (2D-array), but we can work with it like a 2D array
	matrix := []byte(strings.Repeat("#", numRows*numCols))

	crow, ccol := 0, 0

	// We go down the matrix, and then up-right and write each letter from the original string
	// into the corresponding cell of the matrix according to the task conditions
	direction := Direction_Down

	for _, c := range str {
		pos := crow*numCols + ccol // current position in the matrix
		matrix[pos] = c

		if direction == Direction_Down { // Go down
			crow++
			if crow == numRows {
				direction = Direction_UpRight
				crow -= 2
				ccol++
			}
		} else { // Go up-right
			crow--
			ccol++
			if crow < 0 {
				crow = 1
				ccol--
				direction = Direction_Down
			}
		}
	}

	// Read letters from the matrix and fill result string
	for i, j := 0, 0; i < len(matrix) && j < len(result); i++ {
		if matrix[i] != '#' {
			result[j] = matrix[i]
			j++
		}
	}

	return string(result)
}

func test(s string, numRows int, answer string) {
	fmt.Printf("\nInput: s = %v, numRows = %d\n", s, numRows)

	result := convert(s, numRows)
	fmt.Printf("Output: %v\n", result)

	if result != answer {
		fmt.Printf("Result failed, should be: %v", answer)
	}

}

func main() {
	// Example 1:
	// Input: s = "PAYPALISHIRING", numRows = 3
	// Output: "PAHNAPLSIIGYIR"
	test("PAYPALISHIRING", 3, "PAHNAPLSIIGYIR")

	// Example 2:
	// Input: s = "PAYPALISHIRING", numRows = 4
	// Output: "PINALSIGYAHRPI"
	test("PAYPALISHIRING", 4, "PINALSIGYAHRPI")
}
