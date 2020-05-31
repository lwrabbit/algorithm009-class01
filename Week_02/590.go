package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

type Node struct {
	Val      int
	Children []*Node
}

func Postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	for i := 0; i < len(root.Children); i++ {
		res = append(res, Postorder(root.Children[i])...)
	}
	return append(res, root.Val)
}
