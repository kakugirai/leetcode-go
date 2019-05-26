package main

import (
	"fmt"
)

// TreeNode is a binary tree node
type TreeNode struct {
    Val int
    Left *TreeNode
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

type BSTIterator struct {
	current int
	result []int
}

func Constructor(root *TreeNode) BSTIterator {
	result := inorderTraversal(root)
	return BSTIterator{current: -1, result: result}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	this.current++
	return this.result[this.current]
}


/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if this.current+1 >= len(this.result) {
		return false
	}
	return true
}

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
	fmt.Println(obj.Next())
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */