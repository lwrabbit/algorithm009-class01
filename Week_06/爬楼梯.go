package main

import "fmt"

//DP: f(n) = f(n-1) + f(n-2)
func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	f1, f2, f3 := 1, 2, 3
	for i := 2; i < n; i++ {
		f1, f2, f3 = f2, f3, f1+f2
	}
	return f3
}

//DP: f(n) = f(n-1) + f(n-2) + f(n-3)
//time: O(n) space: O(1)
func climbStairs1(n int) int {
	if n < 3 {
		return n
	}
	//初始化DP状态
	f1, f2, f3 := 1, 2, 4
	fn := 0
	//第4阶台阶递推到n级台阶
	for i := 3; i < n; i++ {
		fn := f1 + f2 + f3
		f1, f2, f3 = f2, f3, fn
	}
	return fn
}

//DP: f(n) = f(n-1) + f(n-2) + f(n-3)
//time: O(n) space: O(n)
func climbStairs2(n int) int {
	//1级台阶：1
	//2级台阶：1 （1+1步伐相同,需要排除）
	//3级台阶：3  (1+1+1步伐相同,需要排除)
	if n < 3 {
		return 1
	}
	if n == 3 {
		return n
	}
	//DP状态定义 dp[i][0，1，2] 0-最后一步跨1阶 1-最后一步跨2阶 2-最后一步跨3阶
	//DP方程 dp[i][0] = dp[i-1][1]+dp[i-1][2]
	//      dp[i][1] = dp[i-2][0]+dp[i-2][2]
	//      dp[i][2] = dp[i-3][0]+dp[i-3][1]
	dp := make([][]int,n)
	dp[0] = make([]int,3)
	dp[0][0],dp[0][1],dp[0][2] = 1,0,0
	dp[1] = make([]int,3)
	dp[1][0],dp[1][1],dp[1][2] = 0,1,0
	dp[2] = make([]int,3)
	dp[2][0],dp[2][1],dp[2][2] = 1,1,1
	for i := 3; i < n; i++{
		dp[i] = make([]int,3)
		dp[i][0] = dp[i-1][1]+dp[i-1][2]
		dp[i][1] = dp[i-2][0]+dp[i-2][2]
		dp[i][2] = dp[i-3][0]+dp[i-3][1]
	}
	return dp[n-1][0]+dp[n-1][1]+dp[n-1][2]
}

func main() {
	fmt.Println(climbStairs2(5))
}
