package main

import (
	"container/heap"
	"fmt"
)

// Point in a coordinate system
type Point struct {
	x        int
	y        int
	distance int
	index    int
}

// PriorityQueue contains points
type PriorityQueue []*Point

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].distance > pq[j].distance
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push point into PriorityQueue
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	point := x.(*Point)
	point.index = n
	*pq = append(*pq, point)
}

// Pop point from PriorityQueue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	point := old[n-1]
	point.index = -1 // for safety
	*pq = old[0 : n-1]
	return point
}

func kClosest(points [][]int, K int) [][]int {
	pq := make(PriorityQueue, len(points))
	for i := range points {
		pq[i] = &Point{
			x:        points[i][0],
			y:        points[i][1],
			distance: (points[i][0]*points[i][0] + points[i][1]*points[i][1]) * -1,
			index:    i,
		}
	}
	heap.Init(&pq)

	var result [][]int
	for K > 0 {
		point := heap.Pop(&pq).(*Point)
		result = append(result, []int{point.x, point.y})
		K--
	}
	return result
}

func main() {
	points := [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	fmt.Printf("%#v", kClosest(points, 2))
}
