package main

import (
	"fmt"
)

type stack []int

func (stk *stack) Len() int {
	return len(*stk)
}

func (stk *stack) Empty() bool {
	return len(*stk) == 0
}

func (stk *stack) Push(num int) {
	*stk = append(*stk, num)
}

func (stk *stack) Pop() int {
	num := (*stk)[len(*stk)-1]
	*stk = (*stk)[:len(*stk)-1]
	return num
}

func (stk *stack) Peek() int {
	return (*stk)[len(*stk)-1]
}

type MyQueue struct {
	PushStack *stack
	PopStack  *stack
}

// Initialize your data structure here.
func Constructor() MyQueue {
	return MyQueue{
		PushStack: new(stack),
		PopStack:  new(stack),
	}
}

// Push element x to the back of queue.
func (q *MyQueue) Push(x int) {
	q.PushStack.Push(x)
}

// Removes the element from in front of queue and returns that element.
func (q *MyQueue) Pop() int {
	if q.PopStack.Empty() {
		for q.PushStack.Len() > 0 {
			q.PopStack.Push(q.PushStack.Pop())
		}
		return q.PopStack.Pop()
	} else {
		return q.PopStack.Pop()
	}
}

// Get the front element.
func (q *MyQueue) Peek() int {
	if q.PopStack.Empty() {
		for q.PushStack.Len() > 0 {
			q.PopStack.Push(q.PushStack.Pop())
		}
		return q.PopStack.Peek()
	} else {
		return q.PopStack.Peek()
	}
}

// Returns whether the queue is empty.
func (q *MyQueue) Empty() bool {
	return q.PopStack.Empty() && q.PushStack.Empty()
}

func main() {
	obj := Constructor()
	obj.Push(1)

	popped := obj.Pop()
	fmt.Println(popped)

	isEmpty := obj.Empty()
	fmt.Println(isEmpty)
}
