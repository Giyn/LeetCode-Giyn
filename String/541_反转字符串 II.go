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

import "fmt"

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
			reverse541(ss[i : i+k])
		} else {
			reverse541(ss[i:])
		}
	}
	return string(ss)
}

func reverse541(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
