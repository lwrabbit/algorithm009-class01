package main

func PreorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	r1 := PreorderTraversal(root.Left)
	r2 := PreorderTraversal(root.Right)
	return append(append([]int{root.Val}, r1...), r2...)
}
