package main

import "fmt"

func findCircleNum(M [][]int) int {
	// 参数检查
	if len(M) == 0 || len(M) != len(M[0]){
		return 0
	}
	visited := make([]bool,len(M))
	ret := 0
	for i := 0; i < len(M); i++{
		if !visited[i]{
			dfs(M,visited,i)
			ret++
		}
	}
	return ret
}

func dfs(M [][]int, visited []bool, i int) {
	for j := 0; j < len(M); j++{
		if M[i][j] == 1 && !visited[j]{
			visited[j] = true
			dfs(M,visited,j)
		}
	}
}

func findCircleNum1(M [][]int) int {
	// 参数检查
	if len(M) == 0 || len(M) != len(M[0]){
		return 0
	}
	n := len(M)
	uf := NewUnionFind(n)
	for i :=0;i < n; i++{
		for j := 0; j < n; j++{
			if M[i][j] == 1{
				uf.Union(i,j)
			}
		}
	}
	parentSet := make(map[int]struct{})
	for i := 0; i < n; i++{
		parentSet[uf.Find(i)] = struct{}{}
	}
	return len(parentSet)
}

func main() {
	M := [][]int{{1,1,0},{1,1,1},{0,1,1}}
	M1 := [][]int{{1,1,0},{1,1,0},{0,0,1}}
	fmt.Println(findCircleNum(M))
	fmt.Println(findCircleNum(M1))
	fmt.Println(findCircleNum1(M))
	fmt.Println(findCircleNum1(M1))
}
