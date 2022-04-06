/*
-------------------------------------
# @Time    : 2022/4/6 22:15:11
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1044_最长重复子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "banana"
	fmt.Println(longestDupSubstring(s))
}

func longestDupSubstring(s string) (ans string) {
	check := func(length int) (string, bool) {
		// 使用map记录子串是否出现过
		mp := map[string]bool{}
		for i := 0; i <= len(s)-length; i++ {
			subStr := s[i : i+length]
			if _, ok := mp[subStr]; ok {
				return subStr, true
			}
			mp[subStr] = true
		}
		return "", false
	}
	left, right := 0, len(s)-1
	for left <= right {
		mid := left + (right-left+1)>>1
		// 二分长度
		if res, ok := check(mid); ok {
			ans = res
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return
}
