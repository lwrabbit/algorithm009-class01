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
//func minDepth(root *TreeNode) int {
//	curDepth := 0
//	return minDepthHelper(root,curDepth)
//}
//
//func minDepthHelper(root *TreeNode,curDepth int) int {
//	if root == nil{
//		return curDepth
//	}
//	if root.Left == nil && root.Right == nil{
//		return 1
//	}
//	curDepth++
//	return minInt(minDepthHelper(root.Left,curDepth),minDepthHelper(root.Right,curDepth))
//}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min_depth := math.MaxInt64
	if root.Left != nil {
		min_depth = minInt(minDepth(root.Left), min_depth)
	}
	if root.Right != nil {
		min_depth = minInt(minDepth(root.Right), min_depth)
	}
	return min_depth + 1
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {

}
