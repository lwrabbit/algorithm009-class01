package main

//暴力求解，O(n^2)
func MaxArea(a []int) int {
	max := 0
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			area := (j - i) * minInt(a[i], a[j])
			max = maxInt(max, area)
		}
	}
	return max
}

//左右指针向中间收敛, O(n)
func MaxArea1(a []int) int {
	max := 0
	area := 0
	for i, j := 0, len(a)-1; i < j; {
		if a[i] < a[j] {
			area = (j - i) * a[i]
			i++
		} else {
			area = (j - i) * a[j]
			j--
		}
		max = maxInt(max, area)
	}
	return max
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x < y {
		return y
	}
	return x
}
