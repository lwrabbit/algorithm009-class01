package main

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := &ListNode{val: 1}
	headNext := &ListNode{val: 2}
	headNextNext := &ListNode{val: 3}
	head.next, headNext.pre = headNext, head
	headNext.next, headNextNext.pre = headNextNext, headNext
	//printList(head)
	printList(ReverseList(head))
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.val)
		head = head.next
	}
}
