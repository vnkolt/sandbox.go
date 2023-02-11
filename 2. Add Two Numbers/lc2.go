package main

// https://leetcode.com/problems/add-two-numbers/
//
// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order,
// and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
// Example 1:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [7,0,8]
// Explanation: 342 + 465 = 807.
//
// Example 2:
// Input: l1 = [0], l2 = [0]
// Output: [0]
//
// Example 3:
// Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
// Output: [8,9,9,9,0,0,0,1]
//
// Constraints:
// The number of nodes in each linked list is in the range [1, 100].
// 0 <= Node.val <= 9
// It is guaranteed that the list represents a number that does not have leading zeros.

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func (s *ListNode) AddBack(val int) {
	newListNode := &ListNode{val, nil}
	s.Next = newListNode
}

func add_and_next(node* ListNode, val* int) *ListNode {
	if node != nil { // Unlike C or C++ golang does not allow "if node" without additional operand
		*val += node.Val
		return node.Next;
	}
	return nil
}
 
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}

	node1 := l1
	node2 := l2

	node := result
	remainder := 0;

	for node1 != nil || node2 != nil {
		res := remainder
		node1 = add_and_next(node1, &res)
		node2 = add_and_next(node2, &res)
		node.Val = res % 10;
		remainder = res / 10;
		if remainder != 0 || node1 != nil || node2 != nil {
			node.Next = &ListNode{0, nil}
			node = node.Next
			node.Val = remainder
		}
	}
		
	return result;
 }
 

func initList(nums []int) *ListNode {

	if (len(nums) < 1) {
		return nil;
	}

	root := &ListNode{nums[0], nil}
	sl := root
	for i := 0; i < len(nums)-1; i++ {
		sl.Next = &ListNode{nums[i+1], nil}
		sl = sl.Next;
	}

	return root;
}

func test(nums1 []int, nums2 []int) {
	l1 := initList(nums1)
	l2 := initList(nums2)

	fmt.Printf("\nInput: l1 = %v, l2 = %d\n", nums1, nums2)
	result := addTwoNumbers(l1, l2)
	node := result
	fmt.Printf("Output:")
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
}

func main() {
	test([]int{2,4,3}, []int{5,6,4}) 					 // l1 = [2,4,3], l2 = [5,6,4]
	test([]int{0}, []int{0})         					 // l1 = [0], l2 = [0]
	test([]int{9,9,9,9,9,9,9}, []int{9,9,9,9}) // l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
}
