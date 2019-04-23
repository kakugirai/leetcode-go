package main

import "fmt"

type stack []*TreeNode

func (stk stack) Push(node *TreeNode) stack {
	return append(stk, node)
}

func (stk stack) Pop() (stack, *TreeNode) {
	return stk[:len(stk)-1], stk[len(stk)-1]
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// left, visit, right
	var result []int
	//var stk stack
	for root != nil {
		if root.Left == nil {
			if root.Right == nil {
				result = append(result, root.Val)
			}
			root = root.Right
		}
		root = root.Left
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
	fmt.Println(inorderTraversal(tree))
}
