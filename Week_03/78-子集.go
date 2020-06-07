package main

import "fmt"

//var ans = make([][]int,0)
//var l = make([]int,0)
//
//func subsets(nums []int) [][]int {
//	if len(nums) < 1 {
//		ans = append(ans,[]int{})
//		return ans
//	}
//	rsubsets(nums,0)
//	return ans
//}
//
//func rsubsets(nums []int,index int) {
//	if index == len(nums) {
//		ans = append(ans,[]int{})
//		return
//	}
//	rsubsets(nums,index+1)
//	l = append(l,nums[index])
//	rsubsets(nums,index+1)
//	l = l[:len(l)-1]
//}

func subsets(nums []int)[][]int  {
	res := [][]int{}
	for _,num := range nums{
		newSet := []int{}
		for _,subset := range res{
			newSubSet := append(subset,num)
			newSet = append(newSet,newSubSet...)
		}
		res = append(res,newSet)
	}
	return res
}

func main() {
	nums := []int{1,2,3}
	fmt.Println(subsets(nums))
}
