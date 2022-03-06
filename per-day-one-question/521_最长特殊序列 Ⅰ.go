/*
-------------------------------------
# @Time    : 2022/3/5 0:05:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 521_最长特殊序列 Ⅰ.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	a := "aba"
	b := "cdc"
	fmt.Println(findLUSlength(a, b))
}

func findLUSlength(a string, b string) int {
	if len(a) > len(b) {
		return len(a)
	} else if len(b) > len(a) {
		return len(b)
	} else {
		if a == b {
			return -1
		} else {
			return len(a)
		}
	}
}
