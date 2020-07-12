package main

func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		cnt++
		num &= num-1
	}
	return cnt
}
