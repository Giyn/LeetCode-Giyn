/*
-------------------------------------
# @Time    : 2022/1/31 13:38:33
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 005_最长回文子串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "cbbd"
	fmt.Println(longestPalindrome(s))
}

func longestPalindrome(s string) (ans string) {
	n := len(s)
	for center := 0; center < 2*n-1; center++ {
		left := center / 2
		right := left + center%2
		for left >= 0 && right < n && s[left] == s[right] {
			if right-left+1 > len(ans) {
				ans = s[left : right+1]
			}
			left--
			right++
		}
	}
	return
}
