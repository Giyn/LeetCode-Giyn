/*
-------------------------------------
# @Time    : 2022/3/30 0:44:23
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 014_最长公共前缀.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) (ans string) {
	if len(strs) == 0 {
		return ""
	}
	lcp := func(str1, str2 string) string {
		minLen := Min(len(str1), len(str2))
		idx := 0
		for idx < minLen && str1[idx] == str2[idx] {
			idx++
		}
		return str1[:idx]
	}
	ans = strs[0]
	for i := 1; i < len(strs); i++ {
		ans = lcp(ans, strs[i])
		if len(ans) == 0 {
			break
		}
	}
	return
}
