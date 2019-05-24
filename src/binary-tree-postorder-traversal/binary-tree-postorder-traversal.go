package main

import "fmt"

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

func postorderTraversal(root *TreeNode) []int {
	var result []int
	var stk stack
	if root == nil {
		return result
	}
	stk = stk.Push(root)
	for len(stk) != 0 {
		stk, root = stk.Pop()
		if root.Left != nil {
			stk = stk.Push(root.Left)
		}
		if root.Right != nil {
			stk = stk.Push(root.Right)
		}
		result = append(result, root.Val)
	}

	for i := len(result)/2 - 1; i >= 0; i-- {
		opp := len(result) - 1 - i
		result[i], result[opp] = result[opp], result[i]
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
	fmt.Println(postorderTraversal(tree))
}
