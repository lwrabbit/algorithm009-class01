package main

import "fmt"

/*
二分查找前提：
	1、目标函数单调性
	2、存在上下界
	3、能够通过索引访问
*/
func binary_srearch_template(nums []int, num int) int {
	left,right := 0,len(nums)-1
	for left <= right{
		mid := left + (right-left)/2
		if nums[mid] == num{
			return mid
		}else if nums[mid] < num{
			left = mid+1
		}else{
			right = mid-1
		}
	}
	return -1
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(binary_srearch_template(nums,8))
}