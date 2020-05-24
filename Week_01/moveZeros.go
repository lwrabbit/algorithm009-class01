package main

import "fmt"

func MoveZeroes(nums []int) {
	var j int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			if i != j {
				nums[i] = 0
			}
			j++
		}
	}
}

func MoveZeroes2(nums []int) {
	i, j := 0, 0
	for j < len(nums) {
		nums[i], nums[j] = nums[j], nums[i]
		if nums[i] != 0 {
			i++
		}
		j++
		fmt.Println(nums)
	}
}

func MoveZeroes3(nums []int) {
	for index, num := range nums {
		if num == 0 {
			nums = append(nums[:index], nums[index+1:]...)
			nums = append(nums, 0)
		}
	}
}
