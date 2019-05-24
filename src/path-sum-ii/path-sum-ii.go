package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	var current []int
	findPaths(root, sum, current, &paths)
	return paths
}

func findPaths(root *TreeNode, sum int, current []int, paths *[][]int) {
	if root == nil {
		return
	}
	current = append(current, root.Val)
	if root.Val == sum && root.Left == nil && root.Right == nil {
		*paths = append(*paths, current)
		return
	}
	findPaths(root.Left, sum-root.Val, current, paths)
	findPaths(root.Right, sum-root.Val, current, paths)
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
	fmt.Println(pathSum(tree, 14))
}
