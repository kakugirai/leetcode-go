package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}
	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}
	return root.Val + rangeSumBST(root.Left, L, R) + rangeSumBST(root.Right, L, R)
}

func main() {
	root := &TreeNode{10, &TreeNode{7, &TreeNode{5, nil, nil}, nil}, &TreeNode{15, nil, nil}}
	fmt.Println(rangeSumBST(root, 7, 15))
}
