package main

import (
	"fmt"
	"math"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	minDP, maxDP := make([]int, len(nums)), make([]int, len(nums))
	copy(minDP, nums)
	copy(maxDP, nums)
	ans := maxDP[0]
	for i := 1; i < len(nums); i++ {
		maxDP[i] = max(max(maxDP[i-1]*maxDP[i], maxDP[i]), minDP[i-1]*maxDP[i])
		minDP[i] = min(min(minDP[i-1]*minDP[i], minDP[i]), maxDP[i-1]*minDP[i])
		ans = max(ans, maxDP[i])
	}
	return ans
}

func maxProduct1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	iMax, iMin := 1, 1
	ans := math.MinInt64
	for i := 0; i < len(nums); i++ {
		iMax, iMin = max(max(iMax*nums[i], nums[i]), iMin*nums[i]), min(min(iMin*nums[i], nums[i]), iMax*nums[i])
		fmt.Println(iMax)
		fmt.Println(iMin)
		ans = max(iMax, ans)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	nums := []int{2, 3, -2, 4}
	fmt.Println(maxProduct1(nums))
}
