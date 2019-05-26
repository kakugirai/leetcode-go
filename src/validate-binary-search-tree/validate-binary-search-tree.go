package main

import (
	"fmt"
)

type stack []*TreeNode

func (stk stack) Push(node *TreeNode) stack {
	return append(stk, node)
}

func (stk stack) Pop() (stack, *TreeNode) {
	return stk[:len(stk)-1], stk[len(stk)-1]
}

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var result []int
	if root == nil {
		return true
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

	if len(result) == 1 {
		return true
	}
	for i := 1; i < len(result); i++ {
		if result[i] <= result[i-1] {
			return false
		}
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
	fmt.Println(isValidBST(tree))
}
