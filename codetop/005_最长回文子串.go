/*
-------------------------------------
# @Time    : 2022/4/17 20:18:10
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 005_最长回文子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) (ans string) {
	n := len(s)
	for i := 0; i < len(s); i++ {
		l, r := i, i
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(ans) {
			ans = s[l+1 : r]
		}
		l, r = i, i+1
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		if r-l-1 > len(ans) {
			ans = s[l+1 : r]
		}
	}
	return
}
