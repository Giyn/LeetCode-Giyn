/*
-------------------------------------
# @Time    : 2022/3/9 15:47:53
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 003_无重复字符的最长子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) (ans int) {
	n := len(s)
	mp := make(map[byte]bool, n)
	right := 0
	for i := 0; i < n; i++ {
		if i > 0 {
			// 起始字母前移
			delete(mp, s[i-1])
		}
		for right < n && !mp[s[right]] {
			mp[s[right]] = true
			right++
		}
		ans = math.Max(ans, right-i)
	}
	return ans
}
