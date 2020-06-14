学习笔记

## 二分查找模版
```go
func bsearch(nums []int, target int) int{
	left, right := 0, len(nums) - 1
	for left <= right {
		mid := left + (right - left)>>1
		if nums[mid] == targrt {
			return mid
		}else if nums[mid] < target{
			left = mid -1
		}else{
			right = mid + 1
		}
	}
	return -1
}
```
### 查找旋转有序数组的最小值
- 1 有序数组分成了左右2个小的有序数组，而实际上要找的是右边有序数组的最小值
- 2 如果中间值大于右边的最大值，说明中间值还在左边的小数组里，需要left向右移动
- 3 否则就是中间值小于等于当前右边最大值，mid 已经在右边的小数组里了，但是至少说明了当前右边的right值不是最小值了或者不是唯一的最小值，需要慢慢向左移动一位。
