package main

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, toDelete []int) []*TreeNode {
	var remaining []*TreeNode
	m := make(map[int]struct{})
	// save value to delete into a hash set
	for _, v := range toDelete {
		m[v] = struct{}{}
	}

	// recursively remove nodes
	removeNodes(root, m, &remaining)

	// append the root node to remaining if root is not in toDelete
	if _, ok := m[root.Val]; !ok {
		remaining = append(remaining, root)
	}

	return remaining
}

func removeNodes(root *TreeNode, m map[int]struct{}, remaining *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = removeNodes(root.Left, m, remaining)
	root.Right = removeNodes(root.Right, m, remaining)

	if _, ok := m[root.Val]; ok {
		if root.Left != nil {
			*remaining = append(*remaining, root.Left)
		}
		if root.Right != nil {
			*remaining = append(*remaining, root.Right)
		}
		return nil
	}

	return root
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	toDelete := []int{5, 3}
	fmt.Println(delNodes(root, toDelete))
}
