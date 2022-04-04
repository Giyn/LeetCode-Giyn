/*
-------------------------------------
# @Time    : 2021/11/4 21:45:50
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 151_翻转字符串里的单词.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  Bob    Loves  Alice!   "
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	var ans []string
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			left++
			right++
		} else {
			right++
			if (right < len(s) && s[right] == ' ') || (right == len(s)) {
				ans = append(ans, s[left:right])
				left = right
			}
		}
	}
	left, right = 0, len(ans)-1
	for ; left < right; left, right = left+1, right-1 {
		ans[left], ans[right] = ans[right], ans[left]
	}
	return strings.Join(ans, " ")
}
