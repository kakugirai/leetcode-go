package main

import "fmt"

//ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

func main() {
	node := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	deleteNode(node)
	for i := node; i != nil; i = i.Next {
		fmt.Println(i.Val)
	}
}
