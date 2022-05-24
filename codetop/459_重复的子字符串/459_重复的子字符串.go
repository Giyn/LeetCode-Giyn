/*
-------------------------------------
# @Time    : 2021/11/8 21:36:41
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 459_重复的子字符串.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
)

func main() {
	s := "abab"
	fmt.Println(repeatedSubstringPattern(s))
}

func repeatedSubstringPattern(s string) bool {
	//ss := s + s
	//return strings.Contains(ss[1:len(ss)-1], s)

	n := len(s)
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
