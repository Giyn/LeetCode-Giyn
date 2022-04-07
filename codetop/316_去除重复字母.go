/*
-------------------------------------
# @Time    : 2022/4/7 14:34:20
# @Author  : Giyn
# @Email   : giyn.jy@gmail.com
# @File    : 316_去除重复字母.go
# @Software: GoLand
-------------------------------------
*/

package main

import "fmt"

func main() {
	s := "cbacdcbc"
	fmt.Println(removeDuplicateLetters(s))
}

func removeDuplicateLetters(s string) string {
	remain := [26]int{}
	for _, ch := range s {
		remain[ch-'a']++
	}
	var stack []byte
	inStack := [26]bool{}
	for i := range s {
		if !inStack[s[i]-'a'] {
			for len(stack) > 0 && s[i] < stack[len(stack)-1] {
				top := stack[len(stack)-1] - 'a'
				// 如果后面没剩了,就不删除
				if remain[top] == 0 {
					break
				}
				stack = stack[:len(stack)-1]
				inStack[top] = false
			}
			stack = append(stack, s[i])
			inStack[s[i]-'a'] = true
		}
		remain[s[i]-'a']--
	}
	return string(stack)
}
