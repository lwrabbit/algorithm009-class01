package main

import "fmt"

//数组index操作
func moveZeros(nums []int){
	j := 0
	for i:=0;i< len(nums);i++{
		if nums[i] != 0{
			nums[j],nums[i] = nums[i],nums[j]
			j++
		}
	}
}

func MaxArea(a []int) int {
	max := 0
	for i := 0;i < len(a)-1;i++{
		for j := i+1;j<len(a);j++{
			max := (j-i)
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeros(nums)
	fmt.Println(nums)
}
