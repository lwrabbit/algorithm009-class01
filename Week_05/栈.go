package main

import (
	"container/list"
	"fmt"
)

func isValid(s string) bool {
	l := list.New()
	for _, char := range s {
		switch char {
		case []rune("(")[0]:
			l.PushFront(")")
		case []rune("[")[0]:
			l.PushFront("]")
		case []rune("{")[0]:
			l.PushFront("}")
		case char:
			if l.Len() == 0 || []rune(l.Remove(l.Front()).(string))[0] != char {
				return false
			}
		}
	}
	return l.Len() == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	window := list.New()
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			for window.Len() > 0 && nums[i] > window.Back().Value.(int) {
				window.Remove(window.Back())
			}
			window.PushBack(nums[i])
		} else {
			for window.Len() > 0 && nums[i] > window.Back().Value.(int) {
				window.Remove(window.Back())
			}
			window.PushBack(nums[i])
			max := window.Front().Value.(int)
			res = append(res, max)
			if max == nums[i-k+1] {
				window.Remove(window.Front())
			}
		}
	}
	return res
}

func plusOne(digits []int) []int {
	res := []int{}
	if len(digits) == 0 {
		return res
	}
	digit := digits[len(digits)-1]
	digit++
	if digit < 9 {
		res = append(digits[:len(digits)-1], digit)
	} else {
		res = append(plusOne(digits[:len(digits)-1]), digit)
	}
	return res
}

func main() {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	fmt.Println(maxSlidingWindow(nums, 3))
}
