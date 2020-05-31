package main

import "container/heap"

func GetLeastNumbers(arr []int, k int) []int {
	h := &MinIntHeap{}
	*h = append(*h,arr...)
	heap.Init(h)
	res := []int{}
	for i:=0;i<k;i++{
		res = append(res,heap.Pop(h).(int))
	}
	return res
}

// An IntHeap is a min-heap of ints.
type MinIntHeap []int

func (h MinIntHeap) Len() int           { return len(h) }
func (h MinIntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinIntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinIntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}