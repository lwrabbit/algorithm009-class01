package main

import "fmt"

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

var mSteps = make(map[int]int)

func climbStairs1(n int) int {
	if n < 3 {
		return n
	}
	if v,ok:= mSteps[n];ok{
		return v
	}
	steps := climbStairs(n-1) + climbStairs(n-2)
	mSteps[n] = steps
	return steps
}

func main() {
	fmt.Println(climbStairs(10))
	fmt.Println(climbStairs1(10))
}
