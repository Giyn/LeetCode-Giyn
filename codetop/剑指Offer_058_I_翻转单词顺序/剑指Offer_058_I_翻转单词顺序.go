/*
-------------------------------------
# @Time    : 2022/4/8 23:03:36
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_058_I_翻转单词顺序.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "  hello world!  "
	fmt.Println(reverseWords(s))
}

func reverseWords(s string) string {
	var ans []string
	n := len(s)
	for left, right := n-1, n-1; left >= 0 && right >= 0; {
		for right >= 0 && s[right] == ' ' {
			right--
		}
		if right == -1 {
			break
		}
		left = right
		for left >= 0 && s[left] != ' ' {
			left--
		}
		ans = append(ans, s[left+1:right+1])
		right = left
	}
	return strings.Join(ans, " ")
}
