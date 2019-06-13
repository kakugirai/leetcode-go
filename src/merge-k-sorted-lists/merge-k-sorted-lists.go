package main

import (
	"container/heap"
	"fmt"
)

//ListNode is a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// NodeHeap is a priority queue
type NodeHeap []*ListNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push a ListNode to the NodeHeap
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

// Pop a ListNode from the NodeHeap
func (h *NodeHeap) Pop() interface{} {
	old := *h
	x := old[0]
	heaps := old[1:]
	*h = heaps
	return x
}

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	h := new(NodeHeap)
	heap.Init(h)
	for i := range lists {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	//h.ShowNodeHeap()

	root := &ListNode{0, nil}
	node := root
	for h.Len() > 0 {
		node.Next = h.Pop().(*ListNode)
		node = node.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		//fmt.Println(curr.Val)
		heap.Fix(h, 0)
	}
	return root.Next
}

// ShowListNode shows ListNode
func (l *ListNode) ShowListNode() {
	fmt.Println("Show ListNode")
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

// ShowNodeHeap shows NodeHeap
func (h *NodeHeap) ShowNodeHeap() {
	fmt.Println("Show NodeHeap")
	for i := range *h {
		fmt.Println((*h)[i].Val)
	}
}

func main() {
	//linked_list := [[-10,-9,-9,-3,-1,-1,0],[-5],[4],[-8],[],[-9,-6,-5,-4,-2,2,3],[-3,-3,-2,-1,0]]
	x := &ListNode{-10, &ListNode{-9, &ListNode{-9, &ListNode{-3, &ListNode{-1, &ListNode{-1, &ListNode{0, nil}}}}}}}
	y := &ListNode{-5, nil}
	z := &ListNode{4, nil}
	m := &ListNode{-8, nil}
	n := &ListNode{}
	p := &ListNode{-9, &ListNode{-6, &ListNode{-5, &ListNode{-4, &ListNode{-2, &ListNode{2, &ListNode{3, nil}}}}}}}
	q := &ListNode{-3, &ListNode{-3, &ListNode{-2, &ListNode{-1, &ListNode{0, nil}}}}}
	//linked_list := [[-9,-7,-7],[-6,-4,-1,1],[-6,-5,-2,0,0,1,2],[-9,-8,-6,-5,-4,1,2,4],[-10],[-5,2,3]]
	//x := &ListNode{-9, &ListNode{-7, &ListNode{-7, nil}}}
	//y := &ListNode{-6, &ListNode{-4, &ListNode{-1, &ListNode{1, nil}}}}
	//z := &ListNode{-6, &ListNode{-5, &ListNode{-2, &ListNode{0, &ListNode{0, &ListNode{1, &ListNode{2, nil}}}}}}}
	//m := &ListNode{-9, &ListNode{-8, &ListNode{-6, &ListNode{-5, &ListNode{-4, &ListNode{1, &ListNode{2, &ListNode{4, nil}}}}}}}}
	//n := &ListNode{-10, nil}
	//p := &ListNode{-5, &ListNode{2, &ListNode{3, nil}}}
	lists := []*ListNode{x, y, z, m, n, p, q}
	root := mergeKLists(lists)
	root.ShowListNode()
}
