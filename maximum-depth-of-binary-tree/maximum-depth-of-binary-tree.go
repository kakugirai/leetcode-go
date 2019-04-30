package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	answer := 1
	var getMaxDepth func(root *TreeNode, depth int)
	getMaxDepth = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			if depth > answer {
				answer = depth
			}
		}
		getMaxDepth(root.Left, depth+1)
		getMaxDepth(root.Right, depth+1)
	}
	getMaxDepth(root, answer)
	return answer
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
	fmt.Println(maxDepth(tree))
}
