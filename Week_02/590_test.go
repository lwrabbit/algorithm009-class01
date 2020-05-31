/*
给定一个 N 叉树，返回其节点值的后序遍历。

例如，给定一个 3叉树 :
返回其后序遍历: [5,6,3,2,4,1].
说明: 递归法很简单，你可以使用迭代法完成此题吗?

*/
package main

import (
	"fmt"
	"testing"
)

func TestPostorder(t *testing.T) {
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

	fmt.Println(Postorder(root))
}
