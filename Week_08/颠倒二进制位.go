package main

func reverseBits(num uint32) uint32 {
	i, j := uint32(0), uint32(31)
	for i < j {
		ti, tj := 0, 0
		if num&(1<<i) != 0 {
			ti = 1 // 得到第i位是1还是0
		}
		if num&(1<<j) != 0 {
			tj = 1 // 得到第j位是1还是0
		}
		if ti == 0 { // 将第i位的放到第j位
			num &= ^(1 << j)
		} else {
			num |= 1 << j
		}
		if tj == 0 { // 将第j位的放到第i位
			num &= ^(1 << i)
		} else {
			num |= 1 << i
		}
		i++
		j--
	}
	return num
}
