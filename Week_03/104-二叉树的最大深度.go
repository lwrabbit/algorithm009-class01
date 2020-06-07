package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	depth := 0
	return maxDepthHelper(root,depth)
}

func maxDepthHelper(root *TreeNode,curDepth int)  int{
	if root == nil{
		return curDepth
	}
	curDepth++
	return maxInt(maxDepthHelper(root.Left,curDepth),maxDepthHelper(root.Right,curDepth))
}

func maxInt(x,y int)int  {
	if x<y{
		return y
	}
	return x
}

func main() {

}