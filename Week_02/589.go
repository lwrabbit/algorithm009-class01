package main

func Preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	for i := 0; i < len(root.Children); i++ {
		res = append(res, Preorder(root.Children[i])...)
	}
	return append([]int{root.Val}, res...)
}
