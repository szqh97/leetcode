package main

import "fmt"

/*
*    给定一个字符串，找出不含有重复字符的最长子串的长度。
*
*    示例 1:
*
*    输入: "abcabcbb"
*    输出: 3
*    解释: 无重复字符的最长子串是 "abc"，其长度为 3。
*
*    示例 2:
*
*    输入: "bbbbb"
*    输出: 1
*    解释: 无重复字符的最长子串是 "b"，其长度为 1。
*
*    示例 3:
*
*    输入: "pwwkew"
*    输出: 3
*    解释: 无重复字符的最长子串是 "wke"，其长度为 3。
*         请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串。
*      abcakbcbb -> bcak cakb

*
 */

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}
func lengthOfLongestSubstring2(s string) int {

	tb := [256]int{0}
	i := 0
	j := 0
	maxlen := 0
	for ; j < len(s); j++ {
		c := s[j]
		i = max(tb[c], i)
		maxlen = max(maxlen, j-i+1)
		tb[c] = j + 1
	}
	return maxlen
}
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	base := struct{}{}
	maxlen := 0
	for p := 0; p < len(s); p++ {
		curlen := 0
		mp := make(map[byte]struct{})
		for q := p; q < len(s); q++ {
			c := s[q]
			if _, ok := mp[c]; !ok {
				curlen = curlen + 1
				mp[c] = base
				if curlen > maxlen {
					maxlen = curlen
				}
			} else {
				break
			}
			if q == len(s)-1 {
				return maxlen
			}
		}
	}

	return maxlen
}
func main() {
	s := "pwwkew"
	s = "bbbb"
	s = "abcabcbb"
	s = "abcakbcbb"
	//	s = "aauk"
	fmt.Println(lengthOfLongestSubstring2(s))

}
