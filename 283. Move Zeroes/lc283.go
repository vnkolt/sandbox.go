package main

// https://leetcode.com/problems/move-zeroes/description/
//
// 283. Move Zeroes
// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.
//
// Example 1:
// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:
// Input: nums = [0]
// Output: [0]

import "fmt" // We need this here because we use fmt.Printf in this package

func moveZeroes(nums []int) {
  j := 0 // In place declaration with implicit type
  // Move non-zero items to the left
  for i := 0; i < len(nums); i++ {
    if nums[i] != 0 {
      nums[j] = nums[i];
      j++;
    }
  }

  // Fill in the rest of the items with zeros
  for j < len(nums) {
    nums[j] = 0 // I'd like to write nums[j++] = 0 here
    j++         // but Go does not support it
  }
}

func testCase1() {
  nums := []int{0,1,0,3,12} // They call it "slice" (a fixed size array)
  fmt.Printf("Example 1:\nInput: nums = %v\n", nums) // nums will be printed in square brackets like [0 1 0 3 12]
  moveZeroes(nums)
  fmt.Printf("Output: = %v\n", nums)
}

func testCase2() {
  nums := []int{0}
  fmt.Printf("Example 2:\nInput: nums = %v\n", nums)
  moveZeroes(nums)
  fmt.Printf("Output: = %v\n", nums)
}

func main() {
  testCase1()
  testCase2()
}
