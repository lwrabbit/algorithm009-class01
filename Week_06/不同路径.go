package main

import "fmt"

// time: O(m*n) space: O(m*n)
func myUniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// time: O(m*n) space: O(n)
func optWithDynamicRowMyUniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

// time: O(m*n) space: O(m)
func optWithDynamicColMyUniquePaths(m int, n int) int {
	dp := make([]int, m)
	for i := 0; i < m; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[m-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	for i := 0; i < len(obstacleGrid); i++ {
		for j := 0; j < len(obstacleGrid[0]); j++ {
			if i == 0 && j == 0 {
				if obstacleGrid[i][j] == 1 {
					return 0
				}
				obstacleGrid[i][j] = 1
			}
			if i == 0 && j != 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i][j-1]
				}
			}
			if i != 0 && j == 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j]
				}
			}
			if i != 0 && j != 0 {
				if obstacleGrid[i][j] == 1 {
					obstacleGrid[i][j] = 0
				} else {
					obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
				}
			}
		}
	}
	return obstacleGrid[len(obstacleGrid)-1][len(obstacleGrid[0])-1]
}

func main() {
	fmt.Println(myUniquePaths(7, 3))
	fmt.Println(optWithDynamicRowMyUniquePaths(7, 3))
	fmt.Println(optWithDynamicColMyUniquePaths(7, 3))
}
