package main

import "fmt"

//ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	if head == nil {
		return dummy
	}

	dummy.Next = head
	slow := dummy
	fast := dummy

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return dummy.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	removeNthFromEnd(head, 2)
	for root := head; root != nil; root = root.Next {
		fmt.Println(root.Val)
	}
}
