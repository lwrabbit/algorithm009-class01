/*
给定一个二叉树，返回它的 前序 遍历。

 示例:

输入: [1,null,2,3]
   1
    \
     2
    /
   3

输出: [1,2,3]
进阶: 递归算法很简单，你可以通过迭代算法完成吗？
*/
package main

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	rightleft := &TreeNode{Val: 3}
	right := &TreeNode{Val: 2, Left: rightleft}
	root := &TreeNode{Val: 1, Right: right}
	fmt.Println(PreorderTraversal(root))
}
