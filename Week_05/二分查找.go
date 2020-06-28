package main

import "fmt"

//简单二分查找，找出一个数
func binary_search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target{
			return mid
		}else if nums[mid] < target {
			left = mid + 1
		}else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}

//找出第一个出现的目标索引
func binary_search_left_bound(nums []int, target int) int  {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target{
			right = mid - 1
		}else if nums[mid] < target{
			left = mid + 1
		}else if nums[mid] > target{
			right = mid - 1
		}
	}
	if left == len(nums) || nums[left] != target{
		return -1
	}
	return left
}

//找出最后一个出现的目标索引
func binary_search_right_bound(nums []int, target int) int  {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target{
			left = mid + 1
		}else if nums[mid] < target{
			left = mid + 1
		}else if nums[mid] > target{
			right = mid - 1
		}
	}
	if right < 0 || nums[left-1] != target{
		return -1
	}
	return right
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(binary_search(nums,7))
	fmt.Println(binary_search(nums,11))
	nums1 := []int{1,2,3,4,4,4,4,8,9,10}
	fmt.Println(binary_search_left_bound(nums1,11))
	fmt.Println(binary_search_right_bound(nums1,11))
}
