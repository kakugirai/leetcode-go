package main

import "fmt"

func main() {
	var l1, l2 *ListNode
	*l1 = struct {
		Val  int
		Next *ListNode
	}{Val: 1, Next: nil}

	*l2 = struct {
		Val  int
		Next *ListNode
	}{Val: 2, Next: nil}
	fmt.Println(*mergeTwoLists(l1, l2))
}

// ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// return the other linked list if nil
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	// find the link list started with smaller node
	// set the first node as head
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	// loop through l1 and l2
	// shift to the next node
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			l1, node.Next = l1.Next, l1
		} else {
			l2, node.Next = l2.Next, l2
		}
		node = node.Next
	}

	// if one of the link lists ended, link to the other one
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}

	return head
}
