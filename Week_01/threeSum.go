package main

import "sort"

//暴力法 O(n^3),没实现去重
func ThreeSum(nums []int) [][]int {
	ret := make([][]int, 0, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ret
}

//双指针夹逼
func ThreeSum1(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	ret := make([][]int, 0, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left, right := i+1, len(nums)-1; left < right; {
			if nums[i]+nums[left]+nums[right] == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right {
					if nums[left] == nums[left+1] {
						left++
						continue
					}
					if nums[right] == nums[right-1] {
						right--
						continue
					}
					break
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}
