/*
-------------------------------------
# @Time    : 2022/5/13 17:33:58
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 面试题_01_05_一次编辑.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	first := "pale"
	second := "ple"
	fmt.Println(oneEditAway(first, second))
}

func oneEditAway(first string, second string) bool {
	if len(first) > len(second) {
		first, second = second, first
	}
	if len(second)-len(first) > 1 {
		return false
	}
	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			if len(first) == len(second) {
				return first[i+1:] == second[i+1:]
			}
			return first[i:] == second[i+1:]
		}
	}
	return true
}
