package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var IsBalanced bool

func isBalanced(root *TreeNode) bool {
	IsBalanced = true
	getHeight(root)
	return IsBalanced
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)

	if abs(left-right) > 1 {
		IsBalanced = false
	}

	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func main() {
	t := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(isBalanced(t))
}
