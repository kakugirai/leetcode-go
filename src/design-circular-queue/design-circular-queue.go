package main

import "fmt"

// MyCircularQueue is a circular queue
type MyCircularQueue struct {
	elements []int
	capacity int
	front    int
	rear     int
}

// Constructor initializes the data structure. Set the size of the queue to be k.
func Constructor(k int) MyCircularQueue {
	cq := MyCircularQueue{capacity: k, front: 0, rear: 0}
	cq.elements = make([]int, cq.capacity)
	cq.front = -1
	cq.rear = -1
	return cq
}

// EnQueue inserts an element into the circular queue. Return true if the operation is successful.
func (cq *MyCircularQueue) EnQueue(value int) bool {
	if cq.IsFull() {
		return false
	}
	if cq.IsEmpty() {
		cq.front = 0
	}
	cq.rear = (cq.rear + 1) % cq.capacity
	cq.elements[cq.rear] = value
	return true
}

// DeQueue deletes an element from the circular queue. Return true if the operation is successful.
func (cq *MyCircularQueue) DeQueue() bool {
	if cq.IsEmpty() {
		return false
	}
	if cq.front == cq.rear {
		cq.front = -1
		cq.rear = -1
		return true
	}
	cq.front = (cq.front + 1) % cq.capacity
	return true
}

// Front gets the front item from the queue.
func (cq *MyCircularQueue) Front() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.elements[cq.front]
}

// Rear gets the last item from the queue.
func (cq *MyCircularQueue) Rear() int {
	if cq.IsEmpty() {
		return -1
	}
	return cq.elements[cq.rear]
}

// IsEmpty checks whether the circular queue is empty or not.
func (cq *MyCircularQueue) IsEmpty() bool {
	return cq.front == -1
}

// IsFull checks whether the circular queue is full or not.
func (cq *MyCircularQueue) IsFull() bool {
	return cq.front == (cq.rear+1)%cq.capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */

func main() {
	mcq := Constructor(3)
	fmt.Println(mcq)
	fmt.Println(mcq.rear)
	fmt.Println(mcq.IsFull())

	mcq.EnQueue(1)
	fmt.Println(mcq)
	fmt.Println(mcq.rear)
	fmt.Println(mcq.IsFull())

	mcq.EnQueue(2)
	fmt.Println(mcq)
	fmt.Println(mcq.rear)
	fmt.Println(mcq.IsFull())

	mcq.EnQueue(3)
	fmt.Println(mcq)
	fmt.Println(mcq.rear)
	fmt.Println(mcq.IsFull())

	mcq.DeQueue()
	fmt.Println(mcq)
	fmt.Println(mcq.rear)
	fmt.Println(mcq.IsFull())

	mcq.EnQueue(4)
	fmt.Println(mcq)
	fmt.Println(mcq.rear)
	fmt.Println(mcq.IsFull())
}
