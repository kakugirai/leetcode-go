package main

import (
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}

	mid := findMid(head)
	root := &TreeNode{mid.Val, nil, nil}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(mid.Next)
	return root
}

func findMid(head *ListNode) *ListNode {
	slow := head
	fast := head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	return slow
}

func main() {
	head := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	fmt.Println(sortedListToBST(head))
}
