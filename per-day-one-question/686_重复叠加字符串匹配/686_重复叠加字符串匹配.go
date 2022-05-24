/*
-------------------------------------
# @Time    : 2021/12/22 0:05:08
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 686_重复叠加字符串匹配.go
# @Software: GoLand
-------------------------------------
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "abc"
	b := "cabcabca"
	fmt.Println(repeatedStringMatch(a, b))
}

func repeatedStringMatch(a string, b string) (ans int) {
	ans = 1
	tmp := a
	// 一个完整的b可能首部用到a的一部分，尾部用到a的一部分，像A'[AA....AA]A' , [ ] 内必然<=B的长度，故总长<=2*A+B
	maxLen := 2*len(a) + len(b)
	for len(a) <= maxLen {
		fmt.Println(a)
		if strings.Contains(a, b) {
			return
		} else {
			ans++
			a += tmp
		}
	}
	return -1
}
