package main

import "fmt"

//Definition for TreeNode.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return helper(root, p.Val, q.Val)
}

func helper(root *TreeNode, p, q int) *TreeNode {
	r := root.Val
	if p < r && q < r {
		return helper(root.Left, p, q)
	} else if p > r && q > r {
		return helper(root.Left, p, q)
	}
	return root
}

func main() {
	tree := &TreeNode{6, &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}}, &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}}
	p := &TreeNode{3, nil, nil}
	q := &TreeNode{9, nil, nil}
	fmt.Println(lowestCommonAncestor(tree, p, q))
}
