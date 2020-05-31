package main

import (
	"sort"
	"strings"
)

//暴力 O(nlogn)
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var ss, ts []string
	for _, srune := range []rune(s) {
		ss = append(ss, string(srune))
	}
	for _, trune := range []rune(t) {
		ts = append(ts, string(trune))
	}
	sort.Strings(ss)
	sort.Strings(ts)
	return strings.Join(ss, "") == strings.Join(ts, "")
}

//map统计字母次数 o(n)
func IsAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mCount[s[i]]++
		mCount[t[i]]--
	}
	for _, count := range mCount {
		if count != 0 {
			return false
		}
	}
	return true
}

//map
