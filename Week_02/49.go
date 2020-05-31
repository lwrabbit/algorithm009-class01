package main

import (
	"sort"
	"strings"
)

func GroupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	if len(strs) == 1 {
		return [][]string{strs}
	}
	mAnagrams := make(map[string][]int)
	for index, str := range strs {
		sortedStr := sortByteSlice([]byte(str))
		mAnagrams[sortedStr] = append(mAnagrams[sortedStr], index)
	}
	groups := [][]string{}
	for _, v := range mAnagrams {
		group := []string{}
		for _, index := range v {
			group = append(group, strs[index])
		}
		groups = append(groups, group)
	}
	return groups
}

func sortByteSlice(bs []byte) string {
	ss := []string{}
	for _, b := range bs {
		ss = append(ss, string(b))
	}
	sort.Strings(ss)
	return strings.Join(ss, "")
}
