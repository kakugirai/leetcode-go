package main

import (
	"fmt"
	"strconv"
)

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var result []string
	if root == nil {
		return result
	}
	var temp string
	result = dfs(root, temp, result)
	return result
}

func dfs(root *TreeNode, temp string, result []string) []string {
	temp = temp + strconv.Itoa(root.Val)

	if root.Left == nil && root.Right == nil {
		result = append(result, temp)
		return result
	}

	if root.Left != nil {
		result = dfs(root.Left, temp+"->", result)
	}
	if root.Right != nil {
		result = dfs(root.Right, temp+"->", result)
	}
	return result
}

func main() {
	t := &TreeNode{1, &TreeNode{2, &TreeNode{5, nil, nil}, nil}, &TreeNode{3, nil, nil}}
	fmt.Printf("%#v", binaryTreePaths(t))
}
