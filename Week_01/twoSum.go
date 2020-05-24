package main

//暴力法 O(n^2)
func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//两遍hash O(n)
func TwoSum1(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = i
	}
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if valIdx, ok := numMap[complement]; ok {
			if valIdx != i {
				return []int{i, valIdx}
			}
		}
	}
	return []int{}
}

//一遍hash
func TwoSum2(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		if valIdx, ok := numMap[complement]; ok {
			return []int{i, valIdx}
		}
		numMap[nums[i]] = i
	}
	return []int{}
}
