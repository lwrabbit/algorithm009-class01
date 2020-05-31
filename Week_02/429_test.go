/*
给定一个 N 叉树，返回其节点值的层序遍历。 (即从左到右，逐层遍历)。
例如，给定一个 3叉树 :
返回其层序遍历:
[
     [1],
     [3,2,4],
     [5,6]
]
说明:
树的深度不会超过 1000。
树的节点总数不会超过 5000。
*/
package main

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	grandson1 := &Node{Val: 5}
	grandson2 := &Node{Val: 6}
	children1 := &Node{Val: 3}
	children1.Children = append(children1.Children, grandson1)
	children1.Children = append(children1.Children, grandson2)
	children2 := &Node{Val: 2}
	children3 := &Node{Val: 4}
	root := &Node{Val: 1}
	root.Children = append(root.Children, children1)
	root.Children = append(root.Children, children2)
	root.Children = append(root.Children, children3)

	fmt.Println(LevelOrder(root))
}
