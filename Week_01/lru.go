package main

type LRUCahce struct {
	cap    int
	header *Node
	tail   *Node
	m      map[int]*Node
}

type Node struct {
	key   int
	value int
	pre   *Node
	next  *Node
}

func Constructor(capacity int) LRUCahce {
	cache := LRUCahce{
		cap:    capacity,
		header: &Node{},
		tail:   &Node{},
		m:      make(map[int]*Node, capacity),
	}
	cache.header.next = cache.tail
	cache.tail.pre = cache.header
	return cache
}

func (lc *LRUCahce) Get(key int) int {
	//去缓存m中查找node
	if node, ok := lc.m[key]; !ok {
		//没有找到
		return -1
	} else {
		//存在，删除node并移动到head
		lc.remove(node)
		lc.putHead(node)
		return node.value
	}
}

func (lc *LRUCahce) Put(key, value int) {
	//首先去缓存m中查找此key是否存在
	if node, ok := lc.m[key]; ok {
		//存在，更新node.value,更新m
		node.value = value
		lc.m[key] = node
		//删除node并移动到head
		lc.remove(node)
		lc.putHead(node)
	} else {
		//不存在，检查容量
		if len(lc.m) >= lc.cap {
			//容量满，删除链尾,更新m
			deleteKey := lc.tail.pre.key
			lc.remove(lc.tail.pre)
			delete(lc.m, deleteKey)
		}
		//新建node,插入head，更新m
		newNode := &Node{key: key, value: value}
		lc.putHead(newNode)
		lc.m[key] = newNode
	}
}

func (lc *LRUCahce) remove(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}

func (lc *LRUCahce) putHead(node *Node) {
	curHeadNext := lc.header.next

	lc.header.next = node

	node.pre = lc.header
	node.next = curHeadNext

	curHeadNext.pre = node
}
