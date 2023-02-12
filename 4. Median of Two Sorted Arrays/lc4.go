package main

// 4. Median of Two Sorted Arrays
//
// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
//
// Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
// The overall run time complexity should be O(log (m+n)).
//
// Example 1:
// Input: nums1 = [1,3], nums2 = [2]
// Output: 2.00000
// Explanation: merged array = [1,2,3] and median is 2.
//
// Example 2:
// Input: nums1 = [1,2], nums2 = [3,4]
// Output: 2.50000
// Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
//
// Constraints:
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10^6 <= nums1[i], nums2[i] <= 10^6
//

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64

	// build an array to merge given arrays in ascending order
	merged := make([]int, len(nums1)+len(nums2))

	for i, idx1, idx2 := 0, 0, 0; i < len(merged); i++ {
		if idx1 < len(nums1) && idx2 < len(nums2) {
			if nums1[idx1] < nums2[idx2] {
				merged[i] = nums1[idx1]
				idx1++
			} else {
				merged[i] = nums2[idx2]
				idx2++
			}
		} else if idx1 < len(nums1) {
			merged[i] = nums1[idx1]
			idx1++
		} else {
			merged[i] = nums2[idx2]
			idx2++
		}
	}

	if (len(merged) & 1) != 0 {
		res = float64(merged[len(merged)/2])
	} else {
		res = (float64(merged[len(merged)/2-1]) + float64(merged[len(merged)/2])) / 2
	}

	return res

}

func test(nums1 []int, nums2 []int) {
	fmt.Printf("\nInput: nums1 = %v, nums2 = %d\n", nums1, nums2)
	fmt.Printf("Output: %v", findMedianSortedArrays(nums1, nums2))
}

func main() {
	test([]int{1, 3}, []int{2})
	test([]int{1, 2}, []int{3, 4})
}
