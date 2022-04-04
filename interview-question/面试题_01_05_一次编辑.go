/*
-------------------------------------
# @Time    : 2022/4/4 15:48:59
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 面试题_01_05_一次编辑.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	first := "pales"
	second := "pal"
	fmt.Println(oneEditAway(first, second))
}

func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	lf, ls := len(first), len(second)
	if lf > ls {
		first, second = second, first
		lf, ls = ls, lf
	}
	if ls-lf > 1 {
		return false
	}
	for i := 0; i < lf; i++ {
		if first[i] != second[i] {
			return first[i:] == second[i+1:] || first[i+1:] == second[i+1:] // 替换或者删除(包括添加)
		}
	}
	return true
}
