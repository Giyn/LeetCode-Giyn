/*
-------------------------------------
# @Time    : 2021/11/3 22:52:16
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 344_反转字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
