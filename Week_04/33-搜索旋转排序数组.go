package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[0] < nums[mid] && (nums[mid] < target || nums[0] > target) {
			left = mid + 1
		} else if nums[0] > nums[mid] && nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 1))
}
