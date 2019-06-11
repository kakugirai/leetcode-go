package main

import (
	"container/heap"
	"fmt"
)

type kv struct {
	Key   int
	Value int
}

// KVHeap is a heap of kv
type KVHeap []kv

func (h KVHeap) Len() int           { return len(h) }
func (h KVHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }
func (h KVHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push kv to heap
func (h *KVHeap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

// Pop hv out from heap
func (h *KVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	h := &KVHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, kv{k, v})
	}

	var result []int
	for i := 0; i < k; i++ {
		result = append(result, heap.Pop(h).(kv).Key)
	}
	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(nums, 3))
}
