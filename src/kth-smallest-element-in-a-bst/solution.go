package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack []*TreeNode

func (stk stack) Push(node *TreeNode) stack {
	return append(stk, node)
}

func (stk stack) Pop() (stack, *TreeNode) {
	return stk[:len(stk)-1], stk[len(stk)-1]
}

func kthSmallest(root *TreeNode, k int) int {
	// left, visit, right
	var result []int
	if root == nil {
		return 0
	}
	var stk stack
	for {
		if root.Left != nil {
			stk = stk.Push(root)
			root = root.Left
			continue
		}
		result = append(result, root.Val)
		if root.Right != nil {
			root = root.Right
			continue
		}
		if len(stk) != 0 {
			stk, root = stk.Pop()
			root.Left = nil
			continue
		}
		break
	}
	return result[k-1]
}

func main() {
	k := 3
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}
	fmt.Println(kthSmallest(root, k))
}
