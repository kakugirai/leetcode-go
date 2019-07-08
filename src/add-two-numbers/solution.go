package main

import "fmt"

//ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	c1 := l1
	c2 := l2
	var result = new(ListNode)
	root := result
	sum := 0
	for c1 != nil || c2 != nil {
		sum /= 10
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}
		root.Next = &ListNode{sum % 10, nil}
		root = root.Next
	}

	if sum/10 == 1 {
		root.Next = &ListNode{1, nil}
	}
	return result.Next
}

func main() {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	root := addTwoNumbers(l1, l2)
	for i := root; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}
