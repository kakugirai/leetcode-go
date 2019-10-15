package main

import (
	"fmt"
)

// ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	a := headA
	b := headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}
	return a
}

func main() {
	headA := &ListNode{2, &ListNode{4, nil}}
	headB := &ListNode{9, &ListNode{1, headA}}
	headIntersection := getIntersectionNode(headA, headB)
	for head := headIntersection; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}
