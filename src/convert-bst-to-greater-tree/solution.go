package main

import "fmt"

//TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	dfs(root, &sum)
	return root
}

func dfs(root *TreeNode, sum *int) {
	if root == nil {
		return
	}
	dfs(root.Right, sum)
	root.Val += *sum
	*sum = root.Val
	dfs(root.Left, sum)
}

func main() {
	tree := &TreeNode{5, &TreeNode{2, nil, nil}, &TreeNode{13, nil, nil}}
	fmt.Println(convertBST(tree))
}
