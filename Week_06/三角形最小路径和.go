package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	dp := triangle
	// dp[i][j] 表示到 [i,j] 结点的最小路径和
	// dp方程: dp[i][j] = min(dp[i+1][j],dp[i+1][j+1]) + dp[i][j]
	for i := len(dp)-2; i >= 0; i--{
		for j := 0 ; j < len(dp[i]); j++{
			dp[i][j] += min(dp[i+1][j],dp[i+1][j+1])
		}
	}
	return dp[0][0]
}


func min(x, y int) int  {
	if x > y {
		return y
	}
	return x
}

func main() {
	triangle := [][]int{{2},{3,4},{6,5,7},{4,1,8,3}}
	fmt.Println(minimumTotal(triangle))
}
