package main

import "container/heap"

func MaxSlidingWindow(nums []int, k int) []int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	n := len(nums)
	res := make([]int, n-k+1)
	for i := 0; i < n; i++ {
		start := i - k
		if start >= 0 {
			pq.Remove(nums[start])
		}
		item := &Item{
			value:    nums[i],
			priority: nums[i],
		}

		heap.Push(&pq, item)
		if pq.Len() == k {
			res[start+1] = pq[0].value
		}
	}
	return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// todo interface compare
func (pq *PriorityQueue) Remove(value int) {
	for i, v := range *pq {
		if v != nil && v.value == value {
			heap.Remove(pq, i)
			break
		}
	}
}