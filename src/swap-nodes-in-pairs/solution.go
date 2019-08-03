package main

import "fmt"

// ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := head.Next
	head.Next = swapPairs(head.Next.Next)
	n.Next = head
	return n
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	ss := swapPairs(head)
	for root := ss; root != nil; root = root.Next {
		fmt.Println(root.Val)
	}
}
