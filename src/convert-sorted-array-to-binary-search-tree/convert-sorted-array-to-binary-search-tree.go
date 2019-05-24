package main

import "fmt"

// TreeNode is a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	return constructBSTRecursive(nums, 0, len(nums)-1)
}

func constructBSTRecursive(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := left + (right-left)/2
	var current TreeNode
	current.Val = nums[mid]
	current.Left = constructBSTRecursive(nums, left, mid-1)
	current.Right = constructBSTRecursive(nums, mid+1, right)
	return &current
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	t := sortedArrayToBST(nums)
	fmt.Println(t.Left.Right)
}
