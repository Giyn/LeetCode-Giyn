/*
-------------------------------------
# @Time    : 2021/11/3 23:02:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 541_反转字符串 II.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	. "LeetCodeGiyn/utils"
	"fmt"
)

func main() {
	s := "abcdefg"
	k := 2
	fmt.Println(reverseStr(s, k))
}

func reverseStr(s string, k int) string {
	ss := []byte(s)
	n := len(s)
	for i := 0; i < n; i += 2 * k {
		if i+k <= n {
			Reverse(ss[i:i+k], 0, len(ss[i:i+k])-1)
		} else {
			Reverse(ss[i:], 0, len(ss[i:])-1)
		}
	}
	return string(ss)
}
