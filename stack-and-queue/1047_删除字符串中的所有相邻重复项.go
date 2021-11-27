/*
-------------------------------------
# @Time    : 2021/11/10 0:54:38
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 1047_删除字符串中的所有相邻重复项.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "abbaca"
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		if len(stack) != 0 {
			if stack[len(stack)-1] == s[i] {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, s[i])
	}
	return string(stack)
}
