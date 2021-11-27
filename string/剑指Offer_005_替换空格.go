/*
-------------------------------------
# @Time    : 2021/11/4 17:24:39
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 剑指Offer_005_替换空格.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}

func replaceSpace(s string) string {
	spaceCount := 0
	for _, c := range s {
		if c == ' ' {
			spaceCount++
		}
	}
	ss := make([]byte, len(s)+2*spaceCount)
	last := len(ss) - 1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			ss[last] = '0'
			ss[last-1] = '2'
			ss[last-2] = '%'
			last -= 2
		} else {
			ss[last] = s[i]
		}
		last--
	}
	return string(ss)
}
