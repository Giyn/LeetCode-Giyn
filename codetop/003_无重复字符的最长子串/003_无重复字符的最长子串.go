/*
-------------------------------------
# @Time    : 2022/4/16 23:57:21
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 003_无重复字符的最长子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) (ans int) {
	mp := make(map[byte]bool, len(s))
	right := 0
	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(mp, s[i-1])
		}
		for right < len(s) && !mp[s[right]] {
			mp[s[right]] = true
			right++
		}
		ans = max(ans, right-i)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
