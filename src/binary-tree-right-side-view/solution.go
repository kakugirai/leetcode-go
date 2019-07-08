package main

import (
	"fmt"
)

//TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	result := new([]int)
	if root == nil {
		return *result
	}
	getRightSideView(root, result, 0)
	return *result
}

func getRightSideView(root *TreeNode, result *[]int, level int) {
	if root == nil {
		return
	}
	if len(*result) == level {
		*result = append(*result, root.Val)
	}
	getRightSideView(root.Right, result, level+1)
	getRightSideView(root.Left, result, level+1)
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
	fmt.Printf("%#v", rightSideView(tree))
}
