/*
242. 有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/
package main

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	s0, t0 := "anagram", "nagaram"
	s1, t1 := "rat", "car"
	fmt.Println(IsAnagram(s0, t0))
	fmt.Println(IsAnagram(s1, t1))
}

func TestIsAnagram1(t *testing.T) {
	s0, t0 := "anagram", "nagaram"
	s1, t1 := "rat", "car"
	fmt.Println(IsAnagram1(s0, t0))
	fmt.Println(IsAnagram1(s1, t1))
}
