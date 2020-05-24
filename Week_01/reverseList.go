package main

type ListNode struct {
	val int
	next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.next == nil{
		return head
	}
	var pre *ListNode
	cur := head
	for{
		if cur == nil{
			return pre
		}
		cur.next,pre,cur = pre,cur,cur.next
	}
}
