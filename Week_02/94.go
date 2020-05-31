package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res []int

func InorderTraversal(root *TreeNode) []int {
	Inorder(root)
	return res
}

func Inorder(root *TreeNode) {
	if root != nil {
		Inorder(root.Left)
		res = append(res, root.Val)
		Inorder(root.Right)
	}
}

func InorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r1 := InorderTraversal1(root.Left)
	r2 := InorderTraversal1(root.Right)
	return append(append(r1, root.Val), r2...)
}
