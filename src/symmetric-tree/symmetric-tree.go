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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return getSymmetric(root.Left, root.Right)
}

func getSymmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val &&
		getSymmetric(left.Right, right.Left) &&
		getSymmetric(left.Left, right.Right)
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
	fmt.Println(isSymmetric(tree))
}
