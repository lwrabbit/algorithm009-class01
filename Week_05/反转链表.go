package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var curPre *ListNode
	for cur != nil {
		cur.Next, curPre, cur = curPre, cur, cur.Next
	}
	return curPre
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstNode, secondNode := head, head.Next
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode
	return secondNode
}

func swapPairs1(head *ListNode) *ListNode {
	newHead := head
	for head != nil && head.Next != nil {
		head.Next, head.Next.Next = head.Next.Next, head
		head = head.Next.Next
	}
	if newHead.Next == nil {
		return newHead
	}
	return newHead.Next
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	preNode := dummy
	for head != nil && head.Next != nil {
		firstNode, secondNode := head, head.Next
		preNode.Next, firstNode.Next, secondNode.Next = secondNode, secondNode.Next, firstNode

		preNode = firstNode
		head = firstNode.Next
	}
	return dummy.Next
}

func swapPairs2(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	preNode := dummy
	for head != nil && head.Next != nil {
		firstNode, secondNode := head, head.Next
		preNode.Next, firstNode.Next, secondNode.Next = secondNode, secondNode.Next, firstNode

		preNode = firstNode
		head = firstNode.Next
	}
	return dummy.Next
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return false
	}
	fast,slow := head.Next.Next,head.Next
	for fast.Next != nil {
		if fast == slow{
			return true
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return false
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil{
		return false
	}
	fast,slow := head.Next,head
	for fast != slow{
		if fast == nil || fast.Next == nil{
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil{
		return nil
	}
	curGroupHead := head
	cur := head
	for i := 1; i <= k ;i++{
		if cur == nil{
			return  head
		}
		cur = cur.Next
	}
	newHead := reverse(curGroupHead,cur)
	curGroupHead.Next = reverseKGroup(cur,k)
	return newHead
}

func reverse(curGroupHead,nextGroupHead *ListNode)  *ListNode {
	cur := curGroupHead
	var pre *ListNode
	for cur != nextGroupHead{
		pre,cur,cur.Next = cur,cur.Next,pre
	}
	return pre
}