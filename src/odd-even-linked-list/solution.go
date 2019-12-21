package main

import "fmt"

// ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

func main() {
	head := &ListNode{2, &ListNode{1, &ListNode{3, &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{7, nil}}}}}}}
	result := oddEvenList(head)
	for i := result; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}
