package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack []*TreeNode

func (s stack) Push(node *TreeNode) stack {
	return append(s, node)
}

func (s stack) Pop() (stack, *TreeNode) {
	return s[:len(s)-1], s[len(s)-1]
}

func preorderTraversal(root *TreeNode) []int {
	var result []int
	var rights stack
	for root != nil {
		result = append(result, root.Val)
		if root.Right != nil {
			rights = rights.Push(root.Right)
		}
		root = root.Left
		if root == nil && len(rights) != 0 {
			rights, root = rights.Pop()
		}
	}
	return result
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
	fmt.Printf("%#v", preorderTraversal(tree))
}
