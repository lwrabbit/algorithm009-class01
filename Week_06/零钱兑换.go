package main

import (
	"fmt"
	"math"
)

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		min := math.MaxInt64
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i && dp[i-coins[j]] < min {
				min = dp[i-coins[j]] + 1
			}
		}
		dp[i] = min
	}
	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
