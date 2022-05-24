/*
-------------------------------------
# @Time    : 2021/11/6 17:40:40
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_058_II_左旋转字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "abcdefg"
	n := 2
	fmt.Println(reverseLeftWords(s, n))
}

func reverseLeftWords(s string, n int) string {
	// b := []byte(s)
	// b = append(b[n:], b[:n]...)
	// return string(b)

	// 不开新空间
	b := []byte(s)
	reverse(b, 0, n-1)
	reverse(b, n, len(b)-1)
	reverse(b, 0, len(b)-1)
	return string(b)
}

func reverse(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
