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
	s := "asdfasdfasdf"
	fmt.Println(repeatedSubstringPattern(s))
}

func getNext459(next []int, s string) {
	j := 0
	next[0] = j
	for i := 1; i < len(s); i++ {
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		if s[i] == s[j] {
			j++
		}
		next[i] = j
	}
}

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 0 {
		return false
	}
	next := make([]int, n)
	getNext459(next, s)
	if next[n-1] != 0 && n%(n-next[n-1]) == 0 {
		return true
	}
	return false
}

//func repeatedSubstringPattern(s string) bool {
//	ss := s + s
//	return strings.Contains(ss[1:len(ss)-1], s)
//}
