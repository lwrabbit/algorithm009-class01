package main

type UnionFind struct {
	count  int
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionFind{
		n,
		parent,
	}
}

func (nf *UnionFind) Find(p int) int {
	for p != nf.parent[p] {
		nf.parent[p] = nf.parent[nf.parent[p]]
		p = nf.parent[p]
	}
	return p
}

func (nf *UnionFind) Union(p,q int){
	rootP := nf.Find(p)
	rootQ := nf.Find(q)
	if rootP == rootQ {
		return
	}
	nf.parent[rootP] = rootQ
	nf.count--
}
