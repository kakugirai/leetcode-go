package main

import "fmt"

//TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)

	if root.Left == nil {
		return
	}

	node := root.Left
	for node.Right != nil {
		node = node.Right
	}

	node.Right = root.Right
	root.Right = root.Left
	root.Left = nil
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
	flatten(tree)

	for node := tree; node != nil; node = node.Right {
		fmt.Println(node.Val)
	}
}
