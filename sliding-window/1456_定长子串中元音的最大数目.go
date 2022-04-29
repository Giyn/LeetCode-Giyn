/*
-------------------------------------
# @Time    : 2022/4/28 18:25:57
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1456_定长子串中元音的最大数目.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils/math"
	"fmt"
)

func main() {
	s := "tnfazcwrryitgacaabwm"
	k := 4
	fmt.Println(maxVowels(s, k))
}

func maxVowels(s string, k int) (ans int) {
	mp := map[byte]bool{}
	for _, ch := range []byte{'a', 'e', 'i', 'o', 'u'} {
		mp[ch] = true
	}
	cnt := 0
	for i := 0; i < k; i++ {
		if mp[s[i]] {
			cnt++
		}
	}
	ans = cnt
	// 滑动窗口
	for left, right := 0, k; right < len(s); left, right = left+1, right+1 {
		if mp[s[right]] {
			cnt++
		}
		if mp[s[left]] {
			cnt--
		}
		ans = Max(ans, cnt)
	}
	return
}
