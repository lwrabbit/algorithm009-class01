package main

import (
	"fmt"
	"math"
	"sort"
)

func maxSubArray(nums []int) int {
	// dp[i]: 表示0...i-1的包含[i-1]元素最大连续子数组的和
	// DB方程: dp[i] = max(dp[i-1]+nums[i], nums[i])
	dp := make([]int,len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++{
		//dp[i] = max(max(dp[i-1], dp[i-1]+nums[i]),nums[i])
		dp[i] = max(dp[i-1]+nums[i],nums[i])
	}
	sort.Ints(dp)
	return dp[len(dp)-1]
}

func maxSubArray1(nums []int) int {
	sum := math.MinInt64
	for i := 1; i < len(nums); i++{
		nums[i] = max(nums[i-1]+nums[i],nums[i])
		sum = max(sum,nums[i])
	}
	return sum
}

func max(x, y int) int {
	if x > y{
		return x
	}
	return y
}



func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray1(nums))
}