package main

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//错误代码，bst的定义是左右子树，而非左右节点
//func isValidBST(root *TreeNode) bool {
//	if root == nil{
//		return true
//	}
//	leftValid := true
//	rightValid := true
//	if root.Left != nil{
//		isValidBST(root.Left)
//		leftValid = root.Val > root.Left.Val
//	}
//	if root.Right != nil{
//		isValidBST(root.Right)
//		rightValid = root.Val < root.Right.Val
//	}
//	return leftValid && rightValid
//}

// 递归判断每个节点的上界和下界
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower int, upper int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

//中序遍历=升序遍历
var last *TreeNode

func isValidBST1(root *TreeNode) bool {
	last = &TreeNode{Val: math.MinInt64}
	return dfs(root)
}

func dfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !dfs(root.Left) || root.Val <= last.Val {
		return false
	}
	last = root
	return dfs(root.Right)
}

func main() {

}
