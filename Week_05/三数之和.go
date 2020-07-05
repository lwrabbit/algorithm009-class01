package main

import (
	"fmt"
	"sort"
)

//排序夹逼
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				for nums[j+1] == nums[j] && j < k {
					j++
				}
				j++
			} else if sum > 0 {
				for nums[k-1] == nums[k] && j < k {
					k--
				}
				k--
			} else {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				for nums[j+1] == nums[j] && j < k {
					j++
				}
				j++
				for nums[k-1] == nums[k] && j < k {
					k--
				}
				k--
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			for m, n := j+1, len(nums)-1; m < n; {
				sum := nums[i] + nums[j] + nums[m] + nums[n]
				if sum < target {
					for m < n && nums[m+1] == nums[m] {
						m++
					}
					m++
				} else if sum > target {
					for m < n && nums[n-1] == nums[n] {
						n--
					}
					n--
				} else {
					res = append(res, []int{nums[i], nums[j], nums[m], nums[n]})
					for m < n && nums[m+1] == nums[m] {
						m++
					}
					m++
					for m < n && nums[n-1] == nums[n] {
						n--
					}
					n--
				}
			}
		}
	}
	return res
}

func main() {
	nums := []int{-5, -4, -3, -2, 1, 3, 3, 5}
	fmt.Println(fourSum(nums, -11))
}
