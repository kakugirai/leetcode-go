package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var q []*TreeNode
	var levelNodes []int
	q = append(q, root)
	levelSize := len(q)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		levelSize--
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		levelNodes = append(levelNodes, node.Val)
		if levelSize == 0 {
			tmp := make([]int, len(levelNodes))
			copy(tmp, levelNodes)
			levelNodes = nil
			result = append(result, tmp)
			levelSize = len(q)
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
	fmt.Println(levelOrder(tree))
}
