package main

import (
	"container/heap"
	"fmt"
)

// IntHeap is an array
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push int into IntHeap
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop int out of IntHeap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := IntHeap(nums)
	heap.Init(&h)
	if k == 1 {
		return h[0]
	}
	for i := 1; i < k; i++ {
		heap.Remove(&h, 0)
	}
	return h[0]
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//nums := []int{2, 1}
	fmt.Println(findKthLargest(nums, 4))
}
