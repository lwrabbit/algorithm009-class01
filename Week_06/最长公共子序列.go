package main

import "fmt"

func longestCommonSubsequence(text1 string, text2 string) int {
	len1, len2 := len(text1), len(text2)
	dp := make([][]int, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]int, len2)
		for j := 0; j < len2; j++ {
			if i == 0 && j == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
				}
			}
			if i == 0 && j != 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
			if i != 0 && j == 0 {
				if text1[i] == text2[j] {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i-1][j]
				}
			}
			if i != 0 && j != 0 {
				if text1[i] == text2[j] {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = max(dp[i-1][j], dp[i][j-1])
				}
			}
		}
	}
	return dp[len1-1][len2-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	text1, text2 := "abcde", "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
}
