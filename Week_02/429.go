package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func LevelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	res := getLevelChildren(root.Children...)
	return append([][]int{{root.Val}}, res...)
}

func getLevelChildren(nodes ...*Node) [][]int {
	resLevel := [][]int{}
	nextLevelNodes := []*Node{}
	for i := 0; i < len(nodes); i++ {
		nextLevelNodes = append(nextLevelNodes, nodes[i])
		for j := 0; j < len(nodes[i].Children); j++ {
			res := []int{}
			res = append(res, nodes[i].Children[j].Val)
		}
	}
	resLevel = append(resLevel, res)
	resNextLevel := getLevelChildren(nextLevelNodes...)
	return append(resLevel, resNextLevel...)
}
