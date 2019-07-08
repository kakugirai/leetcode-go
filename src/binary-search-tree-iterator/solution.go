package main

import (
	"fmt"
)

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack []*TreeNode

func (stk stack) Push(node *TreeNode) stack {
	return append(stk, node)
}

func (stk stack) Pop() (stack, *TreeNode) {
	return stk[:len(stk)-1], stk[len(stk)-1]
}

func inorderTraversal(root *TreeNode) []int {
	// left, visit, right
	var result []int
	if root == nil {
		return result
	}
	var stk stack
	for {
		if root.Left != nil {
			stk = stk.Push(root)
			root = root.Left
			continue
		}
		result = append(result, root.Val)
		if root.Right != nil {
			root = root.Right
			continue
		}
		if len(stk) != 0 {
			stk, root = stk.Pop()
			root.Left = nil
			continue
		}
		break
	}
	return result
}

// BSTIterator is a binary search tree iterator
type BSTIterator struct {
	current int
	result  []int
}

// Constructor will initialize a BSTIterator
func Constructor(root *TreeNode) BSTIterator {
	result := inorderTraversal(root)
	return BSTIterator{current: -1, result: result}
}

// Next will return the next smallest number
func (iterator *BSTIterator) Next() int {
	iterator.current++
	return iterator.result[iterator.current]
}

// HasNext will return whether we have a next smallest number
func (iterator *BSTIterator) HasNext() bool {
	if iterator.current+1 >= len(iterator.result) {
		return false
	}
	return true
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
func main() {
	tree := &TreeNode{1,
		&TreeNode{2,
			&TreeNode{4,
				&TreeNode{7, nil, nil},
				nil},
			&TreeNode{5, nil, nil}},
		&TreeNode{3,
			&TreeNode{6,
				&TreeNode{8, nil, nil},
				&TreeNode{9, nil, nil}},
			nil}}
	obj := Constructor(tree)
	fmt.Println(obj.current)
	fmt.Println(obj.result)
	fmt.Println(obj.Next())
}
